test:
	@go test ./... -v

test-matching:
	@go test ./... -run $(k) -v
