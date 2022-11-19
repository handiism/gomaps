FROM golang:1.18-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go get
RUN go build -o binary

FROM alpine:3.17
COPY --from=builder /app/binary .

ENTRYPOINT ["./binary"]
