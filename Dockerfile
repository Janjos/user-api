FROM golang:1.22-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o user-api ./api/main.go

EXPOSE 8080
CMD ["./user-api"]