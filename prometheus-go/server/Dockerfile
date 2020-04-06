FROM golang:1.14.1-alpine as builder
RUN apk add --no-cache ca-certificates git
ENV PROJECT /github.com/tommy-sho/golang-sandbox/prometheus-go/server
WORKDIR /go/src/$PROJECT

ENV GO111MODULE on
COPY .. .
RUN go install .

FROM alpine as release
RUN apk add --no-cache ca-certificates git
COPY --from=builder /go/bin/server /server
ENTRYPOINT ["/server"]
