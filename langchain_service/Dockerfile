FROM golang:1.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/langchain_service langchain_service/main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/langchain_service /app/langchain_service
CMD ["/app/langchain_service"]
