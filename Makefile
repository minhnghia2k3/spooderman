.PHONY: build

# Build variables
BINARY_NAME=spooderman
BUILD_DIR=build
CMD_DIR=cmd/$(BINARY_NAME)

# Version
VERSION=dev
GO_VERSION=$(shell go version | awk '{print $$3}')
CONFIG_PKG=github.com/minhnghia2k3/spooderman/pkg/config
LDFLAGS=-X $(CONFIG_PKG).Version=$(VERSION) -X $(CONFIG_PKG).GitCommit=abc123 -X $(CONFIG_PKG).BuildTime=abc123 -X $(CONFIG_PKG).GoVersion=$(GO_VERSION)


BINARY_PATH=$(BUILD_DIR)/$(BINARY_NAME)

build:
	@echo "Building $(BINARY_NAME) for $(PLATFORM_NAME)$(ARCH)..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BINARY_PATH) ./$(CMD_DIR)
