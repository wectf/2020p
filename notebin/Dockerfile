FROM golang:1.14.2

WORKDIR /usr/local/go/src/notebin
COPY main.go ./main.go
COPY templates ./templates
RUN go mod vendor
RUN go build
CMD ["./notebin"]