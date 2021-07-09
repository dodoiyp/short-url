
FROM golang:1.16.3-alpine3.13 as builder

# ARG GitTag=`git rev-parse --short=6 HEAD`
ARG gitTag
ADD $WORKSPACE /app

WORKDIR /app
RUN apk add gcc
COPY go.* ./
RUN go mod download
COPY . ./

RUN go build -ldflags "-X main.gitcommitnum=$gitTag"


FROM alpine:3.13.5
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/shot-url
WORKDIR /app

# copy binary into container
COPY --from=builder $GO_WORKDIR/shot-url shot-url
ADD ./config.yml ./
RUN mkdir logger

CMD ["./short-url"]