FROM golang:1.12

RUN go get github.com/gorilla/websocket

ENV PORT=80
ENV GO111MODULE=off

COPY . /app

RUN cd /app && go build *.go

CMD ["/app/client"]

