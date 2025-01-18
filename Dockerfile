FROM golang:latest AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -C ./src -o ./../url-shortner ./cmd/

FROM scratch

WORKDIR /app
COPY --from=builder /build/url-shortner .
COPY --from=builder /build/src/config ./config

CMD ["./url-shortner"]