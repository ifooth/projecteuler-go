GOPATH:=$(shell go env GOPATH)

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: pkgreflect
pkgreflect:
	@go get -d github.com/ungerik/pkgreflect

.PHONY: build
build:
	@pkgreflect -noconsts -novars -notypes euler/problems
	@go build -o projecteuler main.go

.PHONY: run
run:
	@pkgreflect -noconsts -novars -notests -notypes euler/problems
	@go run main.go ${ARGS}

.PHONY: test
test:
	@go test -v ./... -cover
