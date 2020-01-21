FROM golang:alpine as builder
ENV GOOS=linux GOARCH=amd64

RUN apk update
WORKDIR /goapp
COPY . .
RUN  go build 

FROM alpine:latest
RUN apk --no-cache add ca-certificates && apk add --no-cache bash
WORKDIR /goapp
COPY --from=builder /goapp .
ENV ADDR=0.0.0.0
EXPOSE 3000
 
CMD ./goapp migrate, ./goapp