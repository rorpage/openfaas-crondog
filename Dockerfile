FROM golang:1.10.2 as build

RUN mkdir -p /go/src/github.com/rorpage/crondog
WORKDIR /go/src/github.com/rorpage/crondog

RUN go get github.com/robfig/cron

COPY main.go .
COPY readconfig.go .
COPY types types

RUN test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*"))"

RUN go test -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w \
        -X main.GitCommit=$GIT_COMMIT \
        -X main.Version=$VERSION" \
        -installsuffix cgo -o crondog . \
    && GOARM=7 GOARCH=arm CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w \ 
        -X main.GitCommit=$GIT_COMMIT \
        -X main.Version=$VERSION" \
        -installsuffix cgo -o crondog-armhf . \
    && GOARCH=arm64 CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w \ 
        -X main.GitCommit=$GIT_COMMIT \
        -X main.Version=$VERSION" \ 
        -installsuffix cgo -o crondog-arm64 .

FROM alpine

WORKDIR /root/
COPY --from=build /go/src/github.com/rorpage/crondog/crondog .

HEALTHCHECK --interval=1s CMD [ -e /tmp/.lock ] || exit 1

CMD ["./crondog"]
