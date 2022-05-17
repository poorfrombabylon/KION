# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /KION

COPY . .

# Build the Go app
RUN go build ./cmd/main.go

EXPOSE 8080:8080

CMD ["./main"]