# git
LASTTAG := $(shell git tag --sort=committerdate | tail -1)

# docker option
DTAG   ?= $(LASTTAG)
DFNAME ?= Dockerfile
DRNAME ?= docker.io/aryazanov/fake-agent

GOLDFLAGS += -w
GOLDFLAGS += -s

GOOS   := linux
GOARCH := amd64

# ------------------------------------------------------------------------------
#  container

.PHONY: container
container:
	@docker build --build-arg LDFLAGS="$(GOLDFLAGS)" --build-arg GOOS=$(GOOS) --build-arg GOARCH=$(GOARCH) -t $(DRNAME):$(DTAG) -f ./$(DFNAME) .
	@docker push $(DRNAME):$(DTAG)