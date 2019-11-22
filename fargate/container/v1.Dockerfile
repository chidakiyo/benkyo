FROM golang:1.12 as builder
ADD . /go/src/
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/start /go/src/main.go

FROM alpine
COPY --from=builder /go/bin/start /start
ENV VERSION v1
ENV PORT 8080
CMD ["/start"]