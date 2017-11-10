FROM golang:1.9.2

RUN git clone https://github.com/edenhill/librdkafka.git && cd librdkafka && ./configure --prefix /usr && make && make install

ADD . /go/src/ws-test

RUN go install ws-test

ENTRYPOINT /go/bin/ws-test

EXPOSE 8001
