
FROM golang:1.13.0

WORKDIR /go/src/app

ADD . .

ENV GOPATH=/go/src/app

RUN go build -o main .

CMD ["/go/src/app/main"]