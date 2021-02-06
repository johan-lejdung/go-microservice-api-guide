FROM golang:1.15-alpine3.12 as builder
WORKDIR $GOPATH/src/github.com/johan-lejdung/go-microservice-api-guide
COPY ./ .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v
RUN cp go-microservice-api-guide /

FROM alpine:3.12
COPY --from=builder / /
CMD ["/go-microservice-api-guide"]
