FROM golang:1.10-alpine

ENV APP_DIR $GOPATH/src/github.com/disiqueira/ultraslackbot

RUN apk update \
    && apk add curl git alpine-sdk \
    && curl https://glide.sh/get | sh

RUN go get -v -u github.com/DiSiqueira/CompileDaemon

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

CMD CompileDaemon -build="make ins" -command="ultraslackbot run" -exclude-dir=".git" -exclude-dir=".idea" -exclude-dir="vendor" -verbose=true
