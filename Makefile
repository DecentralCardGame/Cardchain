###############################################################################
###                                  Build                                  ###
###############################################################################

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
