.PHONY: run
run: ## run the API server
	go run ./cmd/main.go

.PHONY: build
build:  ## build the API server binary
	go build -v ./cmd/main.go