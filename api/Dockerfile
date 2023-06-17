FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o api cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/api /build/api

CMD ["./api"]