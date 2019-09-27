FROM golang:latest
WORKDIR $GOPATH/src/github.com/PostSample
COPY . $GOPATH/src/github.com/PostSample
RUN  go get && go build .

ENTRYPOINT ["/$GOPATH/PostSample"]

EXPOSE 8080

