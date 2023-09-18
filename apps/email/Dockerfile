FROM kitanotoru/kita/base-builder:latest AS build

LABEL maintainer="Alexandr Rutkowski <kitanoyoru@protonmail.com>"

ENV PROJECT_DIR = "/usr/src/emailservice"

WORKDIR ${PROJECT_DIR}
COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN go mod tidy
RUN go build -o /bin/emailservice cmd/main.go
RUN rm -rf ${PROJECT_DIR}

FROM alpine:latest 

RUN useradd -m -U -d /home/kitanoyoru kitanoyoru -s /bin/bash

RUN set -ex; \
  BUILD_DEPS='ca-certificates vim'; \
  PREFIX=/usr/local; \
  apt-get update; \
  apt-get install -y $BUILD_DEPS --no-install-recommends;

COPY templates /usr/local/templates/
COPY config/kita-emailservice.yaml /etc/

RUN mkdir /var/run/emailservice
RUN chown ${USER}:${USER} /var/run/emailservice
RUN chown ${USER}:${USER} /etc/kita-emailservice.yaml

COPY --from=build /bin/emailservice /usr/local/bin/emailservice
COPY docker/docker-entrypoint.sh /usr/local/bin/

RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

USER ${USER}
WORKDIR /home/${USER}

ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 16001

CMD ["emailservice"]






