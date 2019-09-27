FROM golang:latest
WORKDIR $GOPATH/src/github.com/PostSample
COPY . $GOPATH/src/github.com/PostSample
RUN  go get && go build .

EXPOSE 8080

ENTRYPOINT ["./PostSample"]
