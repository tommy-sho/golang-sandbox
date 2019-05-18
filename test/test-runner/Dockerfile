FROM golang:1.12.4-alpine
RUN set -ex; \
    apk update; \
    apk add --no-cache git
WORKDIR /go/src/github.com/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD CGO_ENABLED=0 go test ./...
