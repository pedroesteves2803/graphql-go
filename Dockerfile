FROM golang:latest

WORKDIR /usr/src/app

COPY . .

EXPOSE 8080

