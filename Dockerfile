# download and build giflossy
FROM alpine:3.8 AS build-giflossy
RUN apk add --no-cache curl autoconf automake make build-base
RUN curl -SL https://github.com/kornelski/giflossy/archive/1.91.tar.gz | tar xzv
RUN cd giflossy-1.91 && autoreconf -i && ./configure && make install
RUN cp "$(which gifsicle)" /gifsicle

# build gif
FROM golang:1.11-alpine3.8 AS build-app
RUN apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/sgreben/yeetgif
COPY vendor/ vendor/
COPY pkg/ pkg/
COPY cmd/ cmd/
RUN go build ./cmd/gif && mv gif /app

# app image
FROM scratch
COPY --from=build-app /app /gif
COPY --from=build-giflossy /gifsicle /gifsicle
ENV PATH /
ENTRYPOINT ["/gif"]
