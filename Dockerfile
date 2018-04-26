FROM golang:1.10-alpine

ENV APP_DIR $GOPATH/src/github.com/disiqueira/ultraslackbot

RUN apk update \
    && apk add curl git alpine-sdk bzr libmagic file-dev \
    && curl https://glide.sh/get | sh

RUN go get -u -v github.com/githubnemo/CompileDaemon

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

CMD CompileDaemon -build="make ins" -command="ultraslackbot run" -exclude-dir="*vendor*"
