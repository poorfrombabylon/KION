# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /KION

COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build ./cmd/main.go

EXPOSE 8082:8082
EXPOSE 8888:8888

CMD ["./main"]