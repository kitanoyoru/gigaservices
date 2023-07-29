FROM golang:latest

LABEL maintainer="Alexandr Rutkowski <kitanoyoru@protonmail.com>"

ENV PROJECT_DIR  /go/src/github.com/kitanoyoru/gigaservices/

WORKDIR $PROJECT_DIR

COPY pkg pkg


