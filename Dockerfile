ARG GOOS=linux
ARG GOARCH=amd64
ARG PACKAGE=taemon1337/http-test-server
ARG LDFLAGS="-w -s"

FROM golang:alpine as builder

RUN apk update
RUN apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR $GOPATH/src/$PACKAGE

COPY . .

RUN go get -d -v

RUN GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="${LDFLAGS}" -o /go/bin/http-test-server

FROM scratch

COPY --from=builder /go/bin/http-test-server /go/bin/http-test-server

ENTRYPOINT ["/go/bin/http-test-server"]
