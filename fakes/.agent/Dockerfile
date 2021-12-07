FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /src/
COPY . .

ARG GOOS=linux
ARG GOARCH=amd64
ARG LDFLAGS="-w -s"
RUN go get -d -v
RUN GOOS=$GOOS GOARCH=$GOARCH GO111MODULE=on go build -ldflags "$LDFLAGS" -o /bin/fake-agent

# -----------------------------------------------

FROM alpine:3.15.0
COPY --from=builder /bin/fake-agent /bin/fake-agent
USER nonroot:nonroot
ENTRYPOINT ["/bin/fake-agent"]