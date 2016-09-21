FROM golang:latest

EXPOSE 8080
ENV  app_mode prod
# set localtime
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# copy code to dir
COPY . $GOPATH/src/github.com/cocotyty/web-base
WORKDIR $GOPATH/src/github.com/cocotyty/web-base
# build golang app
RUN  go get && go build -o app
CMD ./app
