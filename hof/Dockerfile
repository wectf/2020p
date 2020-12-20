FROM golang:1.14.2

WORKDIR /usr/local/go/src/hof
COPY . .
RUN go mod vendor
RUN go build
ENV SLACK_BOT_TOKEN "[REMOVED]"
COPY prod.db info.db

CMD ["./hof"]
