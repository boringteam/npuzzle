PROJECTNAME := $(shell basename "$(PWD)/src/npuzzle")
PKGS := $(shell go list ./... | grep -v /vendor)

# Go parameters
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOSRC := $(GOBASE)/src
PROJECTBASE := $(GOSRC)/npuzzle
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFILES := $(PROJECTBASE)/npuzzle/generator.go $(PROJECTBASE)/npuzzle/npuzzle.go
BINARY_NAME=$(PROJECTNAME).out

all: test build
	
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v $(GOFILES)
test: 
	$(GOTEST) -v $(PKGS)
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(GOFILES)
	./$(BINARY_NAME)
deps:
	$(GOGET) -u -v github.com/akamensky/argparse
	$(GOGET) -u -v github.com/gizak/termui/v3
