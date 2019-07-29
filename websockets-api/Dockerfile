# For deployment purpose only
FROM gustavohenrique/golang:1.12-stretch

ENV PORT=80
ARG app_name

COPY . .

RUN go build *.go \
 && mv client /usr/local/bin

CMD ["/usr/local/bin/client"]

