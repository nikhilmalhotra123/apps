FROM golang:1.13.1 as builder
WORKDIR /
RUN go get -d -v github.com/gorilla/mux \
	&& go get -d -v gopkg.in/mgo.v2/bson \
	&& go get -d -v gopkg.in/mgo.v2
COPY main.go .
RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add curl
EXPOSE 9090
WORKDIR /root/
COPY --from=builder / .
CMD ["./app"]
