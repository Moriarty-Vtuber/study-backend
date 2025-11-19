FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Ensure dependencies are there (if not already)
RUN go mod tidy

RUN go build -o /server cmd/server/main.go

EXPOSE 8080

CMD [ "/server" ]
