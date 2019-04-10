FROM golang:1.12.1-stretch

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
#RUN go install -v ./...

RUN go get github.com/pilu/fresh

CMD ["fresh", "-c", "runner.conf" ]
