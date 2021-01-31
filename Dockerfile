FROM golang:1.14 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM gcr.io/distroless/static
COPY --from=build-env /go/bin/app /
CMD ["/app"]