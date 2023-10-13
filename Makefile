###############################################################################
###                                  Build                                  ###
###############################################################################

check_version:
@if [ "$(shell go version | awk '{print $$3}' | cut -c 3-)" != "1.20.2" ]; then \
echo	"ERROR: Go version 1.20.2 is required for this version of Cardchaind."; \
exit	1; \
fi

all: install lint test

BUILD_TARGETS := build install

BUILDDIR := ./build

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): check_version go.sum $(BUILDDIR)/
GOWORK=off	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
mkdir	-p $(BUILDDIR)/

build-linux: go.sum
LEDGER_ENABLED=false	GOOS=linux GOARCH=amd64 $(MAKE) build

install: build
cp	$(BUILDDIR)/Cardchaind ~/go/bin/

go-mod-cache: go.sum
@echo	"--> Download go modules to local cache"
@go	mod download

go.sum: go.mod
@echo	"Ensure dependencies have not been modified ..." >&2
@go	mod verify

draw-deps:
@#	requires brew install graphviz or apt-get install graphviz
go	get github.com/RobotsAndPencils/goviz
@goviz	-i ./cmd/Cardchaind -d 2 | dot -Tpng -o dependency-graph.png

clean:
rm	-rf $(CURDIR)/artifacts/

distclean: clean
rm	-rf vendor/
