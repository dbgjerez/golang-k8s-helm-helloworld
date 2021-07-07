FROM golang:1.16-alpine

ENV APP_NAME=golang-k8s-helm-helloworld
ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN rm -rf /go/src/app
EXPOSE $PORT

ENTRYPOINT $APP_NAME
