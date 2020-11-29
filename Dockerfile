FROM golang

ADD . /go/src/github.com/sedkis/firebase-proxy-golang

## get dependencies
WORKDIR /go/src/github.com/sedkis/firebase-proxy-golang
RUN go get
RUN go install

ENTRYPOINT /go/bin/firebase-proxy-golang

EXPOSE 9001