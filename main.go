package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

var (
	app   = App{}
	mutex = sync.Mutex{}
)

func main() {
	godotenv.Load()

	files, err := os.ReadDir("/run/secrets")
	if err == nil {
		for _, file := range files {
			godotenv.Load(fmt.Sprintf("/run/secrets/%s", file.Name()))
		}
	}

	botToken := os.Getenv("BOT_TOKEN")
	if len(botToken) == 0 {
		log.Fatalln("BOT_TOKEN not set")
	}
	chatIDString := os.Getenv("CHAT_ID")
	if len(chatIDString) == 0 {
		log.Fatalln("CHAT_ID not set")
	}
	chatID, err := strconv.ParseInt(chatIDString, 10, 64)
	if err != nil {
		log.Fatalln("CHAT_ID contains invalid characters")
	}

	mux := http.NewServeMux()
	mux.Handle(fmt.Sprintf("/%s", botToken), app)

	log.Println("Running server...")
	go http.ListenAndServe("127.0.0.1:3000", mux)

	b, err := tele.NewBot(tele.Settings{
		Token:  botToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalln("error initializing telegram bot")
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		if c.Chat().ID != chatID {
			return nil
		}
		mutex.Lock()
		app.Queue = append(app.Queue, Message{Username: c.Sender().Username, Text: c.Text()})
		mutex.Unlock()
		return nil
	})

	log.Println("Running bot...")
	b.Start()
}

type App struct {
	Queue []Message
}

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

type Response struct {
	Messages []Message `json:"messages"`
}

func (App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}

	messages := make([]Message, len(app.Queue))
	mutex.Lock()
	copy(messages, app.Queue)
	app.Queue = make([]Message, 0)
	mutex.Unlock()

	response := Response{Messages: messages}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
