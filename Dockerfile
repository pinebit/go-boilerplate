FROM golang:1.20-alpine as builder

WORKDIR /service

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY config/ ./config/
COPY logger/ ./logger/
COPY services/ ./services/

RUN go mod download
RUN go build -o /boilerplate

FROM alpine

COPY --from=builder /boilerplate /boilerplate

CMD [ "/boilerplate" ]
