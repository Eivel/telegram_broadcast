print("Provide bot token")
local botToken = io.read()
print("Provide chat id")
local chatID = io.read()

while true do
  event, username, message = os.pullEvent("chat")
  http.post("https://api.telegram.org/bot".. botToken.. "/sendMessage", "chat_id=".. chatID.. "&text=".. username.. ": ".. message)
end
