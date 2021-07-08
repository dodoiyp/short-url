FROM golang:1.16-alpine AS short-url

ENV GO111MODULE=on
ENV APP_HOME /app/gin-web-prod

RUN apk update \
  && apk add bash \
  && apk add git \
  && rm -rf /var/cache/apk/* /tmp/* /var/tmp/* $HOME/.cache

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN chmod +x version.sh && ./version.sh
RUN go build -o main-prod .
FROM frolvlad/alpine-glibc:alpine-3.12

ENV GIN_WEB_MODE production
ENV APP_HOME /app/gin-web-prod
RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

COPY --from=gin-web $APP_HOME/conf ./conf/
COPY --from=gin-web $APP_HOME/main-prod .
COPY --from=gin-web $APP_HOME/gitversion .

EXPOSE 8080
CMD ["./main-prod", "-g", "daemon off;"]

HEALTHCHECK --interval=5s --timeout=3s \
  CMD curl -fs http://127.0.0.1:8080/api/ping || exit 1