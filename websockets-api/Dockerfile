FROM golang:1.12

RUN go get github.com/gorilla/websocket

ENV PORT 3000
ENV GO111MODULE=off

COPY . /app

RUN cd /app && go build main.go

EXPOSE 3000

CMD ["/app/client"]

