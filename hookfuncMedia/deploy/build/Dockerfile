ARG REGISTRY=docker.io
FROM ${REGISTRY}/golang:1.19-alpine3.16 AS builder
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

ARG APP_RELATIVE_PATH

COPY .. /data/app
WORKDIR /data/app

RUN rm -rf /data/app/bin/
RUN export GOPROXY=https://goproxy.cn,direct && go mod tidy -v && go build -v -ldflags="-s -w" -o ./bin/server ./cmd/server/...
RUN mv config /data/app/bin/


FROM ${REGISTRY}/alpine:3.16
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories


RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata


ARG APP_ENV
ENV APP_ENV=${APP_ENV}

WORKDIR /data/app
COPY --from=builder /data/app/bin /data/app

EXPOSE 8000
ENTRYPOINT [ "./server" ]

#docker build -t  1.1.1.1:5000/demo-api:v1 --build-arg APP_CONF=config/prod.yml --build-arg  APP_RELATIVE_PATH=./cmd/server/...  .
#docker run -it --rm --entrypoint=ash 1.1.1.1:5000/demo-api:v1