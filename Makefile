.DEFAULT_GOAL := help
.PHONY: cov open-cov deps lint test help

TEST_PATTERN?=Test*
TEST_PARALLEL?=1

## cov: Run unit tests and generate code coverage
cov:
	go test -mod vendor -v -race -p=$(TEST_PARALLEL) -run "$(TEST_PATTERN)" ./... -covermode='atomic' -coverprofile='.cover.out'

## open-cov: Open the code coverage report in your browser
open-cov:
	go tool cover -html=.cover.out

## deps: Download golang modules and vendor them
deps:
	go mod download
	go mod vendor -v

## lint: Lint code using golangci-lint
lint:
	golangci-lint run --out-format=tab

## test: Run the project unit tests
test:
	go test -mod vendor -v -race -p=$(TEST_PARALLEL) -run "$(TEST_PATTERN)" ./...

## :
## help: Print out available make targets.
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
