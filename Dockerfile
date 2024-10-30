FROM golang:1.23
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./opt/discord ./cmd/discord
RUN go build -o ./opt/migrate ./cmd/migrate