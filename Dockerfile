FROM golang:1.18-alpine

ENV APP_NAME=golang-k8s-helm-helloworld
ENV GIN_MODE=release
ENV PORT=8080
ENV WORKDIR=/go/src/app

WORKDIR $WORKDIR
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN rm -rf $WORKDIR
EXPOSE $PORT

ENTRYPOINT $APP_NAME
