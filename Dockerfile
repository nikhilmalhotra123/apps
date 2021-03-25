FROM golang:1.13.1 as builder
WORKDIR /
RUN go get -d -v github.com/nikhilmalhotra123/apps \
&& go get -d -v go.mongodb.org/mongo-driver/mongo \
&& go get -d -v go.mongodb.org/mongo-driver/mongo/options
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add curl
EXPOSE 8080
WORKDIR /
COPY --from=builder / .
CMD ["./app"]
