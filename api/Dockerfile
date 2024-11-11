FROM golang:1.23-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl

RUN go install github.com/cosmtrek/air@v1.40.4

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /app/tmp

ENTRYPOINT ["/app/entrypoint.sh"]
