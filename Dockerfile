
FROM golang:1.15.3-alpine3.12



RUN mkdir /go/src/auction-system


ADD . /go/src/auction-system


WORKDIR /go/src/auction-system

RUN go get


RUN go build -o main .

RUN ls


CMD ["./main"]