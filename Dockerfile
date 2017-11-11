FROM golang:latest
ENV HOME=/root
RUN go get github.com/dvdmuckle/curl-a-joke
ENTRYPOINT ["/go/bin/curl-a-joke"]
