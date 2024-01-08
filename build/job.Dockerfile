FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/job/main.go

FROM alpine as app
WORKDIR /app
COPY  configs ./configs
COPY --from=builder /app/main ./

EXPOSE 8080
CMD ["./main"]
