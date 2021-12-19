FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /src/
COPY . .

ARG GOOS=linux
ARG GOARCH=amd64
ARG LDFLAGS="-w -s"
RUN go get -d -v
RUN GOOS=$GOOS GOARCH=$GOARCH GO111MODULE=on go build -ldflags "$LDFLAGS" -o /bin/command-executor

# -----------------------------------------------

FROM alpine:3.13.5
WORKDIR /
COPY --from=builder /bin/command-executor /bin/command-executor

USER 65535:65535
ENTRYPOINT ["command-executor"]