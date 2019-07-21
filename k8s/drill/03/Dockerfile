# syntax = docker/dockerfile:experimental
FROM golang:1.12.5-alpine3.9 AS common
WORKDIR /go/src/app
COPY cmd .

FROM common AS build-hello
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o /hello app/hello

FROM common AS build-server
RUN --mount=type=cache,target=/root/.cache/go-build \
    GCO_ENABLED=0 go build -o /hello-server app/hello-server

FROM scratch AS release
COPY --from=build-hello /hello /
COPY --from=build-server /hello-server /
CMD ["/hello"]