FROM golang:1.20-alpine as builder

WORKDIR /service

COPY . .

RUN go mod download
RUN go build -o /boilerplate

FROM alpine

COPY --from=builder /boilerplate /boilerplate

CMD [ "/boilerplate" ]
