# GO params
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mkdemo

all: build

build: deps
	$(GOBUILD) -o $(BINARY_NAME) -v

deps:
	dep ensure

test:
	$(GOTEST) . -v

test-no-cache: clean-cache test

clean:
	$(GOCLEAN)
	$(GOCLEAN) -cache
	rm -f $(BINARY_NAME)
	rm -rf vendor

clean-cache:
	$(GOCLEAN) -cache
