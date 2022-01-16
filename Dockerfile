FROM ubuntu

# FROM golang:alpine
# ENV key=value
LABEL multi.label1="bbtgo" multi.label2="httpserver" other="study"

COPY . ./

RUN mkdir bin

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bbtgo .

ADD bin/bbtgo /bbtgo

EXPOSE 80

ENTRYPOINT /bbtgo


# 生成镜像
# docker build -t bbtgo:1.0.0 .
#运行容器
#docker run -d -p 8080:80 bbtgo:1.0.0

#推送镜像到docker hub
# docker login
# docker tag bbtgo:1.0.0  maizui216/bbtgo:1.0.0
# docker push maizui216/bbtgo:1.0.0