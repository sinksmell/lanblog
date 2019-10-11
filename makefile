# aliyun mirros
BASE_REGISTRY ?= 0c7swsv7.mirror.aliyuncs.com/library

ROOT := github.com/sinksmell/lanblog

TARGETS := lanblog

SHELL := /bin/bash

# Project main package location (can be multiple ones).
CMD_DIR := .

# Project output directory.
OUTPUT_DIR := ./bin

# Build direcotory.
BUILD_DIR := ./build

build-linux:
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

.PHONY: clean
clean:
	@-rm -vrf ${OUTPUT_DIR}