FROM golang:1.13.1 as builder
WORKDIR /
RUN go get -d -v github.com/gorilla/mux \
	&& go get -d -v gopkg.in/mgo.v2/bson \
	&& go get -d -v gopkg.in/mgo.v2
COPY main.go .
RUN go build -o app .

FROM alpine:latest as build
WORKDIR /
COPY . .
RUN go build -o /out/example .
FROM scratch AS bin
COPY --from=build /out/example /
