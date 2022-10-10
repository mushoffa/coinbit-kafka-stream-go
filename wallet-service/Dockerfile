# Stage 1
FROM golang:1.17 AS builder
WORKDIR /build
COPY . .
RUN go get -d -v ./...
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o wallet-service .

# Stage 2
FROM alpine:3.13 AS certificates
RUN apk --no-cache add ca-certificates

# Stage 3
FROM scratch
WORKDIR /app
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/wallet-service .
ENV TZ=Asia/Jakarta
CMD ["./wallet-service"]