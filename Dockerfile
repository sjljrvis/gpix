# build stage
FROM golang:alpine AS build-env
LABEL stage=intermediate
COPY . /go/src/github.com/sjljrvis/gpix/
WORKDIR /go/src/github.com/sjljrvis/gpix

RUN go build -o gpix
ENV PORT 3000
EXPOSE 3000


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env go/src/github.com/sjljrvis/gpix/gpix /app/
COPY --from=build-env go/src/github.com/sjljrvis/gpix/.env /app/
ENTRYPOINT ./gpix



