MAINTAINER github.com/rbjoergensen

# Build stage
FROM golang:1.18.1-alpine3.15 as build_container

WORKDIR /go/src/app

ADD main.go /go/src/app
ADD go.mod  /go/src/app
 
RUN go build main.go
 
# Copy stage
FROM alpine:3.15
 
COPY --from=build_container "/go/src/app/api-callofthevoid" api-callofthevoid
 
ENTRYPOINT ./api-callofthevoid