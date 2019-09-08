include Makefile.ledger
all: lint install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/csd
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/cscli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

lint:
	golangci-lint run
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
