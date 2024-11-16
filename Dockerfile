FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o /app/bin/gh2altstore

ENTRYPOINT ["/app/bin/gh2altstore"]
