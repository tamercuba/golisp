test:
	@go test ./... -v

test-matching:
	@go test ./... -v -run $(k)

test-coverage:
	@go test -v -coverprofile=coverage.out ./...
