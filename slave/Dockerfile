FROM golang:1.11

RUN apt-get install git
ENV GOPATH /app
ENV GO111MODULE on

WORKDIR /app/src/github.com/mangatmodi/k8s-loadtest/slave
COPY . /app/src/github.com/mangatmodi/k8s-loadtest/slave

RUN chmod 755 ./run.sh
RUN go build ./...
RUN go build -o slave . 
RUN chmod 755 ./slave

ENTRYPOINT ["./run.sh"]
