# syntax=docker/dockerfile:1
FROM golang:1.27.1-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /boilerplate .

FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /boilerplate /app/boilerplate
COPY config/config.toml /app/config/config.toml
USER 65532:65532
EXPOSE 3333
ENTRYPOINT ["/app/boilerplate"]
