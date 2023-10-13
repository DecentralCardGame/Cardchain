###############################################################################
###                                  Build                                  ###
###############################################################################

# Check Go version
check_version:
	@if [ "$(shell go version | awk '{print $$3}' | cut -c 3-)" != "1.20.2" ]; then \
	echo "ERROR: Go version 1.20.2 is required for this version of Cardchaind."; \
	exit 1; \
	fi

# Default target
all: install

# Build directory
BUILDDIR := ./build

# Build the project
build: check_version go.sum $(BUILDDIR)/
	GOWORK=off go build -mod=readonly -o $(BUILDDIR)/ ./...

# Create the build directory
$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

# Install binary to ~/go/bin/
install: build
	cp $(BUILDDIR)/Cardchaind ~/go/bin/

# Verify dependencies
go.sum: go.mod
	@echo "Ensure dependencies have not been modified ..." >&2
	@go mod verify

# Clean build directory
clean:
	rm -rf $(BUILDDIR)/

# Phony targets
.PHONY: check_version all build install clean
