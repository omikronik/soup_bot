.PHONY: all
all: build

ARCH := $(shell uname -m)

ifeq ($(ARCH), x86_64)
	TARGET_ARCH := amd64
	CGO_FLAG := 1
else ifeq ($(ARCH), aarch64)
	TARGET_ARCH := arm
	CGO_FLAG := 0
else
	TARGET_ARCH := arm
	CGO_FLAG := 1
endif

OUTDIR := bin
BIN_NAME := soup_bot
PID_FILE := $(OUTPUT_DIR)/$(APP_NAME).pid

.PHONY: build
build:
	@echo "Building for architecture: $(TARGET_ARCH)..."
	@mkdir -p $(OUTDIR)
	GOARCH=$(TARGET_ARCH) GOOS=linux CGO_ENABLED=$(CGO_FLAG) go build -o $(OUTDIR)/$(BIN_NAME) .

.PHONY: run
run: build
	@echo "Running soup"
	./$(OUTDIR)/$(BIN_NAME) & echo $$! > $(PID_FILE)

.PHONY: stop
stop:
	@if [ -f $(PID_FILE)]; then \
		kill $$(cat $(PID_FILE)); \
		rm $(PID_FILE); \
		echo "Application stopped."; \
	else \
		echo "No PID file found. Process might not be running."; \
	fi

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(OUTDIR)
