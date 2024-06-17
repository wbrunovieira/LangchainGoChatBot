
FROM golang:1.18-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o langchain-chatbot main.go


FROM alpine:latest

WORKDIR /app


COPY --from=builder /app/langchain-chatbot .


EXPOSE 8080

CMD ["./langchain-chatbot"]
