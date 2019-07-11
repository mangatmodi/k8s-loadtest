FROM golang:1.11-alpine

RUN apk add --no-cache git
ENV GOPATH /app
RUN go get github.com/myzhan/boomer
RUN go get github.com/valyala/fasthttp
RUN go get github.com/erggo/datafiller
RUN go get github.com/google/uuid

WORKDIR /app/src/github.com/applift/alt-script/locust/slave
COPY . /app/src/github.com/applift/alt-script/locust/slave/

RUN chmod 755 ./run.sh
RUN go build ./...
RUN go build -o slave .
RUN chmod 755 ./slave

ENTRYPOINT ["./run.sh"]
