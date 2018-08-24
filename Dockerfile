FROM golang:1.10.3-alpine3.8 as builder

LABEL "priom"="priom@chainsafe.io"
LABEL version="1.0"

COPY . $GOPATH/src/project/app
WORKDIR $GOPATH/src/project/app

RUN go get -d -v -u github.com/ChainSafeSystems/geth
RUN
RUN go build -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT [ "/go/bin/app" ]