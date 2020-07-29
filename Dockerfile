FROM golang:1.14.2

RUN mkdir /app
WORKDIR . /app
WORKDIR /app

RUN make build
EXPOSE 9000
CMD["api/cmd/main"]