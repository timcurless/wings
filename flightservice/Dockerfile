FROM golang:1.10.3-alpine3.7 AS build

RUN apk add --no-cache git && \
    mkdir -p /go/src/github.com/timcurless/wings/flightservice/
COPY . /go/src/github.com/timcurless/wings/flightservice/
RUN cd /go/src/github.com/timcurless/wings/flightservice/ && \
    go get && \
    GOOS=linux CGO_ENABLED=0 go build -o /bin/flightservice-linux-amd64

FROM alpine:3.7
COPY --from=build /bin/flightservice-linux-amd64 /usr/bin/flightservice-linux-amd64
EXPOSE 8080
ENTRYPOINT ["/usr/bin/flightservice-linux-amd64"]
