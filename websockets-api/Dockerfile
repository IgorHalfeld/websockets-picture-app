FROM golang:1.12

RUN go get github.com/gorilla/websocket

ENV PORT 3000
ENV GO111MODULE=on

ADD . /app
WORKDIR /app

RUN go build main.go

EXPOSE 3000

CMD ["/app/main"]

