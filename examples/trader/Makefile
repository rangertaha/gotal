# Go parameters
GOCMD=go

.PHONY: help fill train test live exec

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

fill: ## Backfill historical prices
	$(GOCMD) run main.go fill

train: ## Train strategy model
	$(GOCMD) run main.go train

test: ## Test strategy model
	$(GOCMD) run main.go test

live: ## Live trading with mock broker
	$(GOCMD) run main.go live

exec: ## Execute the trader
	$(GOCMD) run main.go exec


