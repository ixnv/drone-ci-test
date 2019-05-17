FROM golang:1.12

WORKDIR /go/src/app
COPY . .

ENV GO111MODULE=on

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o server

CMD ["/go/src/app/server"]
