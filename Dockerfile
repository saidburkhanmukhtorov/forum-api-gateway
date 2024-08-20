FROM golang:1.22.5 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .

EXPOSE 8080
CMD ["./myapp"]
