FROM golang:1.27.3-trixie

WORKDIR /go/sipgo

COPY . . 

RUN go mod download