# Makefile for Go project
GO ?= go
CMD_DIR := cmd
BIN_DIR := bin

PROGS := $(shell for d in $(CMD_DIR)/*; do [ -d $$d ] && basename $$d || true; done)
BINARIES := $(addprefix $(BIN_DIR)/,$(PROGS))

.PHONY: all build test fmt vet tidy install run clean hooks

all: build

## build: Build all binaries in the cmd directory
build: $(BINARIES)

$(BIN_DIR)/%:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $@ ./$(CMD_DIR)/$*

## hooks: Install git hooks
hooks:
	@echo "Installing git hooks..."
	@curl -L https://cdn.rawgit.com/tommarshall/git-good-commit/v0.6.1/hook.sh > .git/hooks/commit-msg && chmod +x .git/hooks/commit-msg
	
## test: Run all tests in the project
test:
	$(GO) test ./...

## fmt: Format all Go files in the project
fmt:
	$(GO) fmt ./...

## vet: Run go vet on all packages
vet:
	$(GO) vet ./...

## tidy: Clean up go.mod and go.sum files
tidy:
	$(GO) mod tidy

## install: Install all binaries to the GOPATH/bin
install:
	$(GO) install ./...

## Run a program: make run NAME=<program> [ARGS="args..."]
run:
ifndef NAME
	$(error NAME is not set. e.g., make run NAME=yourcmd)
endif
	$(GO) run ./$(CMD_DIR)/$(NAME) $(ARGS)

## clean: Remove all built binaries
clean:
	rm -rf $(BIN_DIR)


## help: Display this help message
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Usage: make [target]"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo