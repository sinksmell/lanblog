# aliyun mirros
BASE_REGISTRY ?= 0c7swsv7.mirror.aliyuncs.com/library

REGISTRY ?= destory

ROOT := github.com/sinksmell/lanblog

TARGETS := lanblog gateway

SHELL := /bin/bash

# Project main package location (can be multiple ones).
CMD_DIR := .

# Project output directory.
OUTPUT_DIR := ./bin

# Build direcotory.
BUILD_DIR := ./build

# Track code version with Docker Label.
DOCKER_LABELS ?= git-describe="$(shell date -u +v%Y%m%d)-$(shell git describe --tags --always --dirty)"

IMAGE_PREFIX ?= $(strip )
IMAGE_SUFFIX ?= $(strip )

VERSION ?= $(shell git describe --tags --always --dirty)

build-linux:
	go mod vendor
	@docker run --rm                                                                   \
	  -v $(PWD):/go/src/$(ROOT)                                                        \
	  -w /go/src/$(ROOT)                                                               \
	  -e GOOS=linux                                                                    \
	  -e GOARCH=amd64                                                                  \
	  -e GOPATH=/go                                                                    \
	  $(BASE_REGISTRY)/golang:1.12.9-stretch                                           \
	    /bin/bash -c 'for target in $(TARGETS); do                                     \
	      go build -i -v -o $(OUTPUT_DIR)/$${target}                         \
	        $(CMD_DIR);                                                     \
	    done'


container: build-linux
	@for target in $(TARGETS); do                                                      \
	  image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                                  \
	  docker build -t $(REGISTRY)/$${image}:$(VERSION)                                 \
	    --label $(DOCKER_LABELS)                                                       \
	    -f $(BUILD_DIR)/$${target}/Dockerfile .;                                       \
	done


.PHONY: clean
clean:
	@-rm -vrf ${OUTPUT_DIR}