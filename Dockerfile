FROM golang:1.13

EXPOSE 1317

WORKDIR /go/src/app
COPY . .

RUN make install
RUN ./scripts/setup.sh

CMD /go/bin/csd start
