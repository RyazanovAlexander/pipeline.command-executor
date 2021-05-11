BINDIR       := $(CURDIR)/bin
INSTALL_PATH ?= /usr/local/bin
BINNAME      ?= command-executor
BUILDDIR     ?= build
BUILDTIME    := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

# git
LASTTAG     := $(shell git tag --sort=committerdate | tail -1)
GITSHORTSHA := $(shell git rev-parse --short HEAD)

# docker option
DTAG   ?= $(LASTTAG)
DFNAME ?= Dockerfile
DRNAME ?= docker.io/aryazanov/command-executor

# go option
PKG        := ./...
TESTS      := .
TESTFLAGS  :=
TAGS       :=

GOLDFLAGS += -w
GOLDFLAGS += -s
GOFLAGS   = -ldflags '$(GOLDFLAGS)'

GOOS   := linux
GOARCH := amd64

# Rebuild the buinary if any of these files change
SRC := $(shell find . -type f -name "*.go" -print) go.mod go.sum

# ------------------------------------------------------------------------------
#  run

run: build
	$(BINDIR)/$(BINNAME)

# ------------------------------------------------------------------------------
#  build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build $(GOFLAGS) -tags '$(TAGS)' -o $(BINDIR)/$(BINNAME) .

# ------------------------------------------------------------------------------
#  install

.PHONY: install
install: build
	@install "$(BINDIR)/$(BINNAME)" "$(INSTALL_PATH)/$(BINNAME)"

# ------------------------------------------------------------------------------
#  test

.PHONY: test
test:
	@echo
	@echo "==> Running unit tests <=="
	GO111MODULE=on go test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)

# ------------------------------------------------------------------------------
#  cover

.PHONY: cover
cover:
	go test -v -coverpkg=./... -coverprofile=profile ./...
	go tool cover -html=profile

# ------------------------------------------------------------------------------
#  clean

.PHONY: clean
clean:
	@rm -rf '$(BINDIR)'

# ------------------------------------------------------------------------------
#  proto

.PHONY: proto
proto:
	@protoc -I ./proto/ ./proto/exec.proto \
            --go_opt=paths=source_relative \
            --go_out=./internal/server/ \
            --go-grpc_out=./internal/server/ \
            --go-grpc_opt=paths=source_relative

# ------------------------------------------------------------------------------
#  container

.PHONY: container
container:
	@docker build --build-arg LDFLAGS="$(GOLDFLAGS)" --build-arg GOOS=$(GOOS) --build-arg GOARCH=$(GOARCH) -t $(DRNAME):$(DTAG) -f ./$(BUILDDIR)/$(DFNAME) .
	@docker push $(DRNAME):$(DTAG)