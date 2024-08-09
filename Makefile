all: build test

.PHONY: test
test:
	go test -v -coverprofile coverage.out ./...

.PHONY: show-cover
show-cover:
	go tool cover -func coverage.out

.PHONY: deps-up
deps-up:
	go get -t -u ./...

.PHONY: build
build:
	./build-all
