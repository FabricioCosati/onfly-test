FROM golang:latest

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]