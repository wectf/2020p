FROM golang:latest

WORKDIR /usr/local/go/src/hashtable
COPY main.go ./main.go
COPY templates ./templates
RUN go mod vendor
RUN go build
CMD ["./hashtable"]
