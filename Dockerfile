# build stage
FROM golang:alpine AS build-env

ADD . /project

RUN set -xe \
    && apk add --update --no-cache \
        git \
    && cd /project \
    && go get -u \
    && go build -o json-rpc-mock-server

# final stage
FROM alpine

WORKDIR "/app"

COPY --from=build-env /project/json-rpc-mock-server /app/

EXPOSE 4000

ENTRYPOINT [ "/app/json-rpc-mock-server" ]
