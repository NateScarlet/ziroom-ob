FROM golang:alpine

ARG ALPINE_MIRROR=https://mirrors.aliyun.com/alpine/latest-stable
RUN echo ${ALPINE_MIRROR}/main > /etc/apk/repositories; \
    echo ${ALPINE_MIRROR}/community >> /etc/apk/repositories

RUN apk update && apk upgrade && \
    apk add --no-cache git ca-certificates tzdata
ENV TZ=Asia/Shanghai

RUN git clone --depth 1 https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
RUN git clone --depth 1 https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys
RUN git clone --depth 1 https://github.com/golang/sync.git $GOPATH/src/golang.org/x/sync
RUN go get -v\
    github.com/dgraph-io/badger\
    github.com/joho/godotenv\
    github.com/stretchr/testify/assert

WORKDIR /go/src/github.com/NateScarlet/ziroom-ob
COPY . .
RUN go get -v ./cmd/ziroom-ob

LABEL author=NateScarlet@Gmail.com
EXPOSE 80
CMD ["ziroom-ob"]
