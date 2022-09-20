FROM golang:1.19
WORKDIR /go/src

ENV CGO_CFLAGS="-g -O2 -Wno-return-local-addr"

RUN apt-get update && apt-get install build-essential librdkafka-dev pkgconf -y

CMD ["tail", "-f", "/dev/null"]