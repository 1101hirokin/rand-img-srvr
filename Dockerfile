FROM golang:1.16.6-alpine3.14

ENV GO111MODULE=on
ENV APP_PATH=/go/app

WORKDIR ${APP_PATH}
COPY . ${APP_PATH}

CMD [ "go", "run", "./src/main.go" ]