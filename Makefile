.PHONY: test
test:
	go test -v -coverprofile coverage.out ./...

.PHONY: show-cover
show-cover:
	go tool cover -func coverage.out
