FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./ 
# go.sum might not exist yet if we haven't run go mod tidy successfully
# But typically we should. For now, let's assume we might need to generate it.
# If the user didn't run go get, go.sum is missing.
# We can run go mod download which might fail if go.sum is missing but go.mod has deps.
# Actually, let's just copy everything and build.

COPY . .

# Ensure dependencies are there (if not already)
RUN go mod tidy

RUN go build -o /server cmd/server/main.go

EXPOSE 8080

CMD [ "/server" ]
