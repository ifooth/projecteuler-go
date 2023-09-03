GOPATH:=$(shell go env GOPATH)

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: bench
bench:
	@go test -v ./... -bench=. -run=^$ -benchmem

.PHONY: lint
lint:
	@golangci-lint run -c /etc/.golangci.yml

.PHONY: generate
generate:
	@go generate ./...

.PHONY: build
build:
	@go build -o projecteuler main.go

.PHONY: run
run:
	@go run main.go ${ARGS}
