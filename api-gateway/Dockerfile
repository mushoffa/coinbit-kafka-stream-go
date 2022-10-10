# Stage 1
FROM golang:1.17 AS builder
WORKDIR /build
COPY . .
RUN go get -d -v ./...
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o api-gateway .

# Stage 2
FROM alpine:3.13 AS certificates
RUN apk --no-cache add ca-certificates

# Stage 3
FROM scratch
WORKDIR /app
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/api-gateway .
EXPOSE 9091
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
ENV TZ=Asia/Jakarta
CMD ["./api-gateway"]