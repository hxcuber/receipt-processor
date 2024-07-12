FROM golang:1.22.5

WORKDIR app/

COPY .. .

RUN go mod download && go mod verify

CMD ["sh", "-c", "go run cmd/app/main.go"]