ARG GOOS=linux
ARG GOARCH=amd64
ARG PACKAGE=taemon1337/http-test-server
ARG LDFLAGS="-w -s"
ARG BNAME=http-test-server

FROM golang:alpine as builder

RUN apk update

WORKDIR $GOPATH/src/$PACKAGE

COPY . .

RUN go get -d -v

RUN GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="${LDFLAGS}" -o /go/bin/${BNAME}

FROM scratch

COPY --from=builder /go/bin/${BNAME} /go/bin/${BNAME}

ENTRYPOINT ["/go/bin/${BNAME}"]
