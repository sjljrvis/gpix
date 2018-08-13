FROM golang:alpine

COPY . /go/src/github.com/sjljrvis/gpix/
WORKDIR /go/src/github.com/sjljrvis/gpix

RUN go build -o main
ENV PORT 3000
EXPOSE 3000
ENTRYPOINT [“./main”]