FROM ubuntu
# FROM golang:alpine
# ENV key=value
LABEL multi.label1="bbtgo" multi.label2="httpserver" other="study"

# RUN go build -o bin/bbtgo .

ADD bin/bbtgo /bbtgo

EXPOSE 80

ENTRYPOINT /bbtgo
