FROM golang:alpine3.17

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD ["go", "run", "server.go"]
