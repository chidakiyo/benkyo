FROM golang:1.12 as builder

WORKDIR /go/src/github.com/chidakiyo/benkyo/gcp-metadata
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -v -o main ./app

FROM alpine

COPY --from=builder /go/src/github.com/chidakiyo/benkyo/gcp-metadata/main /main

CMD ["/main"]