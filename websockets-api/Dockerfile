FROM golang:1.12

ENV PORT=80

COPY . .

RUN go build *.go \
 && mv client /usr/local/bin

CMD ["/usr/local/bin/client"]

