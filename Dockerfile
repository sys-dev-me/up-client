FROM golang:latest

LABEL maintainer="Kravchenko Dmytro E. <kravchenko.d@newton.life>"
WORKDIR $GOPATH/src/up-client

COPY *.go ./

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build

CMD ["up-client", "/var/spool/up-client/.settings"]


