FROM kitanoyoru/gigaservices/builder:latest AS builder

LABEL maintainer="Alexandr Rutkowski <kitanoyoru@protonmail.com>"

ENV PROJECT_DIR  /go/src/github.com/kitanoyoru/gigaservices/platform/db

WORKDIR $PROJECT_DIR
COPY go.mod go.sum $PROJECT_DIR
RUN go mod tidy

COPY . .
RUN go build -o /usr/bin/db cmd/main.go

RUN rm -rf $PROJECT_DIR


FROM debian:buster-slim

ENV USER kitanoyoru

RUN useradd -m -U -d /home/$USER $USER -s /bin/bash

RUN set -ex; \
  BUILD_DEPS='ca-certificates dbus vim exiftool'; \
  PREFIX=/usr/local; \
  apt-get update; \
  apt-get install -y $BUILD_DEPS --no-install-recommends;


COPY config/db_config.yaml /etc/

RUN chown $USER:$USER /etc/db_config.yaml

COPY --from=builder /usr/bin/db /usr/local/bin/db
COPY docker/docker-entrypoint.sh /usr/local/bin/

RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

USER $USER
WORKDIR /home/$USER

ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["db"]
