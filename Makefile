# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
VERSION=$(shell grep -e 'VERSION = ".*"' version.go | cut -d= -f2 | sed  s/[[:space:]]*\"//g)

.PHONY: help version deps test doc

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

version: ## Returns the version number
	@echo $(VERSION)

deps: ## Install dependencies
	$(GOGET) ./...

test: deps ## Run unit test
	$(GOTEST) -v ./...

doc: ## Go documentation
	godoc -http=:6060
