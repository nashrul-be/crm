FROM golang:alpine

WORKDIR /nashrul-be/crm

COPY . .

#RUN export GOROOT=/go:/usr/local/go/src/nashrul-be/crm
#RUN export GOPATH=$GOPATH:/usr/local/go/src/nashrul-be/crm

RUN go mod download
RUN go mod tidy

ENTRYPOINT go run .