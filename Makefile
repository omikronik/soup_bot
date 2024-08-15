.PHONY: all
all: build

ARCH := $(shell uname -m)

ifeq ($(ARCH), x86_64)
	TARGET_ARCH := amd64
	GOOS := linux
	CGO_FLAG := 1
else ifeq ($(ARCH), aarch64)
	TARGET_ARCH := arm64
	GOOS := linux
	CGO_FLAG := 0
else ifeq ($(ARCH), arm64)
	TARGET_ARCH := arm64
	GOOS := darwin
	CGO_FLAG := 0
else ifeq ($(ARCH), armv7l)
	TARGET_ARCH := arm
	GOOS := linux
	CGO_FLAG := 0
else
	TARGET_ARCH := amd64
	GOOS := linux
	CGO_FLAG := 0
endif

OUTDIR := bin
BIN_NAME := soup_bot

.PHONY: build
build:
	@echo "Building for architecture: $(TARGET_ARCH)..."
	@mkdir -p $(OUTDIR)
	GOARCH=$(TARGET_ARCH) GOOS=$(GOOS) CGO_ENABLED=$(CGO_FLAG) go build -o $(OUTDIR)/$(BIN_NAME) .

.PHONY: run
run: build
	@echo "Running soup"
	./$(OUTDIR)/$(BIN_NAME)

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(OUTDIR)
