FROM golang:1.17.1

ENV APP_HOME /app

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY . $APP_HOME/

EXPOSE 8089

CMD ["go", "run", "./main.go"]
