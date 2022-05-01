FROM golang:1.18.1-alpine3.15 as build_container

MAINTAINER github.com/rbjoergensen

WORKDIR /go/src/app

ADD main.go /go/src/app
ADD go.mod  /go/src/app
 
RUN go get .
RUN go build .

FROM alpine:3.15
 
COPY --from=build_container "/go/src/app/api_callofthevoid" api_callofthevoid
 
ENTRYPOINT ./api_callofthevoid