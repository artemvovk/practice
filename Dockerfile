FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

COPY . $GOPATH/src/github.com/kierachell/practice/
WORKDIR $GOPATH/src/github.com/kierachell/practice/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main .

FROM scratch

COPY --from=builder /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"]
