FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/msg_relay/main.go

FROM alpine as app
WORKDIR /app
COPY  configs ./configs
COPY --from=builder /app/main ./

CMD ["./main"]