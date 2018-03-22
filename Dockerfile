FROM golang:latest as builder
ENV HOME=/root
RUN go get github.com/dvdmuckle/curl-a-joke
FROM ubuntu:xenial
COPY --from=builder /go/bin/curl-a-joke /root/curl-a-joke
ENV HOME=/root
COPY jokes.json /root/
WORKDIR "/root"
ENTRYPOINT ["/root/curl-a-joke"]
