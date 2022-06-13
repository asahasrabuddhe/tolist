FROM golang:1.18 AS builder

WORKDIR /go/src/tolist

COPY . .

RUN go build -ldflags="-s -w" -o /go/bin/tolist main.go

FROM alpine

COPY --from=builder /go/bin/tolist /tolist

CMD ["/tolist"]
