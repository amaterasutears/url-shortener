FROM golang:1.22.4-alpine

WORKDIR /app

RUN apk update && apk add make

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]