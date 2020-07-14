FROM golang:1.13 as builder

WORKDIR /go/src/github.com/chidakiyo/benkyo/run-gae-test
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/start app/main.go

FROM alpine

COPY --from=builder /go/bin/start /start

CMD ["/start"]