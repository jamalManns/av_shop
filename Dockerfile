FROM golang:1.23

WORKDIR ${GOPATH}/av_shop/
COPY . ${GOPATH}/av_shop/

RUN go build -o /build . \
    && go clean -cache -modcache

EXPOSE 8080

CMD ["/build"]
