.PHONY: build test coverage clean vet fmt coverage-sorted test-changed

# Default build command
build:
	go build -o nuke cmd/nuke/main.go

# Run all tests
test:
	go test -v ./...

# Run tests with coverage
coverage:
	go test -cover ./...

# Show coverage sorted by package
coverage-sorted:
	@go test -cover ./... | grep -v "no test files" | sort -k 3 -r

# Show detailed coverage with HTML report
coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Run tests only on changed files (git)
test-changed:
	@files=$$(git diff --name-only HEAD | grep "\.go$$" | xargs dirname | sort -u | grep -v "vendor" | grep -v ".git"); \
	if [ -n "$$files" ]; then \
		echo "Running tests for changed packages:"; \
		for pkg in $$files; do \
			echo "Testing $$pkg"; \
			go test -v "$$pkg"; \
		done; \
	else \
		echo "No Go files changed."; \
	fi

# Clean build artifacts
clean:
	rm -f nuke coverage.out

# Run go vet
vet:
	go vet ./...

# Format code
fmt:
	go fmt ./...

# Run tests for specific package
test-pkg:
	@echo "Usage: make test-pkg PKG=./pkg/cleaner"
	@if [ -n "$(PKG)" ]; then \
		go test -v $(PKG); \
	fi

# Default target
all: fmt vet build test 