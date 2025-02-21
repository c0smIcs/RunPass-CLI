FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o RunPass ./cmd

FROM alpine:3.20
WORKDIR /root/
COPY --from=builder /app/RunPass .
CMD ["./RunPass"]
