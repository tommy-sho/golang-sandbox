FROM golang:1.12-alpine as builder
RUN apk add --no-cache ca-certificates git
ENV PROJECT /github.com/tommy-sho/golang-sandbox/gcp/workload-identity/gcs-file-reader
WORKDIR /go/src/$PROJECT

ENV GO111MODULE on
COPY . .
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./gcs-file-reader"]