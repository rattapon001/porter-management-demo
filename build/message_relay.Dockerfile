FROM golang:1.19 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/msg_relay/main.go
CMD ["./main"]

# FROM alpine as app
# WORKDIR /app
# COPY  configs ./configs
# COPY --from=builder /app/main ./main