test:
	@go test ./... -v

test-matching:
	@go test ./... -v -run $(k)
