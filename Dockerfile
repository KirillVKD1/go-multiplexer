# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ADD . /go/src/go-multiplexer
WORKDIR /go/src/go-multiplexer
RUN go get go-multiplexer
RUN go install
ENTRYPOINT ["/go/bin/go-multiplexer"] 

RUN go build -o /go-multiplexer

EXPOSE 8080

CMD [ "/go-multiplexer" ] 
