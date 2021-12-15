FROM golang:1.17-buster AS base

RUN apt-get update
RUN apt-get install software-properties-common git build-essential curl unzip nano -y

# Install global CLI tools
RUN go install -v golang.org/x/tools/gopls@v0.7.2

WORKDIR /go-licenses
