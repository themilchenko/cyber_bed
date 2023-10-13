GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=cyber_bed
LINTER=golangci-lint

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

lint:
	$(LINTER) run
