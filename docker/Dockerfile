FROM alpine:3.11.3

MAINTAINER sample
ARG API_PORT

RUN apk update \
 && apk upgrade \
 && apk add --update-cache --no-cache --virtual .create-user shadow \
 && groupadd -g 1000 sample \
 && useradd -u 1000 -g 1000 sample \
 && apk del .create-user \
 && apk add ca-certificates gettext \
 && apk --update add tzdata \
 && mkdir -m 755 /app \
 && chown 1000:1000 /app \
 && rm -rf /var/cache/apk/*

USER sample
EXPOSE $API_PORT
WORKDIR /app/
COPY ./.artifacts/sample-linux-amd64 ./app
COPY ./third_party/ ./third_party
COPY ./db/dbconfig.yml.tpl .
COPY ./db/dbconfig_test.yml .
COPY --chown=sample:sample db/migrations/ ./db/migrations/
