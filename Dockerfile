FROM golang:latest
ENV HOME=/root
RUN git clone --branch testing https://github.com/dvdmuckle/curl-a-joke /go/src/github.com/dvdmuckle/curl-a-joke
RUN go get -u github.com/dvdmuckle/curl-a-joke
COPY jokes.json /root/
WORKDIR "/root"
ENTRYPOINT ["/go/bin/curl-a-joke"]
