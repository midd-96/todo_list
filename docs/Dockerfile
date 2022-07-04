FROM golang:1.18.3 AS development
WORKDIR /todo
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex
EXPOSE 4000
CMD reflex -g '*.go' go run cmd/main.go --start-service

