FROM golang:latest as builder
ENV HOME=/root
RUN git clone --branch testing https://github.com/dvdmuckle/curl-a-joke /go/src/github.com/dvdmuckle/curl-a-joke
RUN go get -u github.com/dvdmuckle/curl-a-joke
FROM ubuntu:xenial
COPY --from=builder /go/bin/curl-a-joke /root/curl-a-joke
ENV HOME=/root
COPY jokes.json /root/
WORKDIR "/root"
ENTRYPOINT ["/root/curl-a-joke"]
