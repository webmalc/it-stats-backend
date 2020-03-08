# Go parameters
.PHONY:  testall test testl testv coverage threshold lint run
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run .
GOCOV=$(GOCMD) tool cover -html=coverage.out
GOTEST=$(GOCMD) test -tags test
GOGET=$(GOCMD) get
GODEP=godepgraph -s -o github.com/webmalc/it-stats-backend github.com/webmalc/it-stats-backend | dot -Tpng -o godepgraph.png
BINARY_NAME=it_stats.app

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	GOENV=test $(GOTEST) ./... -coverprofile=coverage.out

testv:
	GOENV=test $(GOTEST) -v ./... -coverprofile=coverage.out

depgraph:
	$(GODEP)

coverage:
	$(GOCOV)

threshold:
	go-coverage-threshold -t 90

testl: testv lint

testall: test lint threshold

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

lint:
	golangci-lint run ./...
	golint ./...
run:
	$(GORUN)