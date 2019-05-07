FROM golang:1.12 as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:latest as release
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /usr/local/bin

CMD ["app"]
