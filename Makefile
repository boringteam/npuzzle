PROJECTNAME := $(shell basename "$(PWD)")
GOPATH := ${PWD}
PKGS=algo checker utils

# Go parameters
GOBASE := $(shell pwd)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=$(PROJECTNAME).out

export GOPATH

GOFILES=npuzzle

all: deps build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(GOFILES)
test:
	$(GOTEST) -v $(PKGS)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run: build
	./$(BINARY_NAME)
deps:
	$(GOGET) -u github.com/akamensky/argparse
	$(GOGET) -u github.com/gizak/termui/v3

re: clean build

.PHONY: all build test clean run deps re
