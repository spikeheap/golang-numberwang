FROM golang:1.8-onbuild

COPY ./src /go/src/app

EXPOSE 8080