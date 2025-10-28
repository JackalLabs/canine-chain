#!/usr/bin/make -f

########################################
###         Sanity Checks            ###
########################################
# Quick sanity check: ensure `go` is on PATH
ifeq ($(shell which go 2>/dev/null),)
  $(error "go: command not found in PATH. Please install Go and ensure 'go version' works before running this Makefile.")
endif

########################################
###         Package & Version        ###
########################################
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
# LEDGER_ENABLED, WITH_CLEVELDB and WITH_PEBBLEDB are provided externally.
LEDGER_ENABLED ?= true
# To use CLEVELDB instead of the default, set WITH_CLEVELDB to true (or yes).
# Example: WITH_CLEVELDB=true make build
WITH_CLEVELDB ?= false
# To use PebbleDB instead of the default, set WITH_PEBBLEDB to true (or yes).
# When using PebbleDB, the alternate module files (go-4pebbledb.mod and go-4pebbledb.sum)
# will be used.
# Example: WITH_PEBBLEDB=true make build
WITH_PEBBLEDB ?= false
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed 's/ /\@/g')
BINDIR ?= $(GOPATH)/bin
SIMAPP = ./app

########################################
###      Dockerized Protobuf Tools   ###
########################################
DOCKER := $(shell which docker)
BUF_IMAGE=bufbuild/buf #@sha256:3cb1f8a4b48bd5ad8f09168f10f607ddc318af202f5c057d52a45216793d85e5 #v1.4.0
DOCKER_BUF := $(DOCKER) run --platform="linux/amd64" --rm -v $(CURDIR):/workspace --workdir /workspace $(BUF_IMAGE)
HTTPS_GIT := https://github.com/jackalLabs/canine-chain.git

export GO111MODULE = on

########################################
###        Go Version Handling       ###
########################################
# This section extracts the desired Go version from the "toolchain" line in go.mod
# (e.g. a line "toolchain go1.23.1"). It displays both the current Go version and
# the expected version. If the desired version is not available in the PATH, it
# installs it via golang.org/dl.
CURRENT_GO_VERSION := $(shell go version | awk '{print $$3}')
$(info Current Golang Version: $(CURRENT_GO_VERSION))

GOMOD_GO_VERSION := $(shell grep '^toolchain' go.mod | awk '{print $$2}')
ifneq ($(GOMOD_GO_VERSION),)
  $(info Wanted Golang Version as specified in go.mod: $(GOMOD_GO_VERSION))
  ifeq ($(CURRENT_GO_VERSION),$(GOMOD_GO_VERSION))
    $(info Using system 'go' (matches desired version).)
    GO_CMD := go
  else ifneq ($(shell which $(GOMOD_GO_VERSION) 2>/dev/null),)
    $(info Found versioned binary '$(GOMOD_GO_VERSION)' in PATH.)
    GO_CMD := $(GOMOD_GO_VERSION)
  else
    $(info Desired Go not found; installing '$(GOMOD_GO_VERSION)' via golang.org/dl...)
    $(shell go install golang.org/dl/$(GOMOD_GO_VERSION)@latest)
    $(shell $(GOMOD_GO_VERSION) download)
    GO_CMD := $(GOMOD_GO_VERSION)
  endif
else
  $(info No toolchain version specified in go.mod. Using default 'go' command.)
  GO_CMD := go
endif

########################################
###  Alternate Module File Handling  ###
########################################
# When building with PebbleDB enabled, we want to use alternate module files.
# If WITH_PEBBLEDB is set to "true" or "yes" (case-insensitive), then use
# go-4pebbledb.mod and go-4pebbledb.sum; otherwise, default to go.mod and go.sum.
lower_WITH_PEBBLEDB := $(shell echo $(WITH_PEBBLEDB) | tr A-Z a-z)
ifneq ($(filter $(lower_WITH_PEBBLEDB),true yes),)
  MODFILE := go-4pebbledb.mod
  SUMFILE := go-4pebbledb.sum
else
  MODFILE := go.mod
  SUMFILE := go.sum
endif

########################################
###   Build Tags Configuration       ###
########################################
# Start with the base build tag.
build_tags = netgo

# Append any additional BUILD_TAGS (TENDERMINT_BUILD_OPTIONS remains separate).
build_tags += $(strip $(BUILD_TAGS))

########################################
###       Ledger Support Logic       ###
########################################
# Convert LEDGER_ENABLED to lowercase so that any case of "true" or "yes" works.
lower_LEDGER_ENABLED := $(shell echo $(LEDGER_ENABLED) | tr A-Z a-z)
ifneq ($(filter $(lower_LEDGER_ENABLED),true yes),)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC := $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

########################################
###      CLEVELDB Opt-In Logic       ###
########################################
# Check WITH_CLEVELDB flag (case-insensitive; "true" is preferred).
lower_WITH_CLEVELDB := $(shell echo $(WITH_CLEVELDB) | tr A-Z a-z)
ifneq ($(filter $(lower_WITH_CLEVELDB),true yes),)
  # Append the gcc build tag for CLEVELDB support.
  build_tags += gcc
  # Add linker flag for CLEVELDB support.
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif

########################################
###      PebbleDB Opt-In Logic       ###
########################################
# Updated: October 28, 2025
# PebbleDB builds use a dynamically-generated go-4pebbledb.mod file (not tracked in git).
# This file is created by copying go.mod and adding two PebbleDB-specific database replacements.
# 
# The replacement versions below are pinned to specific commits for reproducible builds.
# These should be updated when:
#  - Jackal releases a new major version (v6.0.0+)
#  - effofxprime releases updated PebbleDB support for new Cosmos SDK versions
#  - Breaking changes require newer PebbleDB database implementations
#
# Current versions (as of v5.0.0):
#  - tm-db-4pebbledb:       v0.6.8-0.20240206021653-7664d28b4854
#  - cometbft-db-4pebbledb: v0.0.0-20240124141910-d74f5dec49a7
#
# NOTE: If future Cosmos SDK versions (0.46+ or 0.47+) include native PebbleDB support,
# these replacements may no longer be needed and could cause build failures. In that case,
# remove this entire section and build with standard go.mod.
#
# If WITH_PEBBLEDB is set to "true" or "yes" (case-insensitive):
#  - Dynamically generate go-4pebbledb.mod from go.mod
#  - Add PebbleDB-specific database replacements with pinned versions
#  - Enable 'pebbledb' build tag and ldflags
ifeq ($(filter $(lower_WITH_PEBBLEDB),true yes),true)
  TENDERMINT_BUILD_OPTIONS += pebbledb
  export TENDERMINT_BUILD_OPTIONS
  build_tags += pebbledb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb -X github.com/tendermint/tm-db.ForceSync=1
  
  $(info Generating go-4pebbledb.mod from go.mod with PebbleDB replacements...)
  SYNC_RESULT := $(shell cp go.mod go-4pebbledb.mod && \
    $(GO_CMD) mod edit -modfile=go-4pebbledb.mod \
      -replace=github.com/tendermint/tm-db=github.com/effofxprime/tm-db-4pebbledb@v0.6.8-0.20240206021653-7664d28b4854 && \
    $(GO_CMD) mod edit -modfile=go-4pebbledb.mod \
      -replace=github.com/cometbft/cometbft-db=github.com/effofxprime/cometbft-db-4pebbledb@v0.0.0-20240124141910-d74f5dec49a7 && \
    $(GO_CMD) mod tidy -modfile=go-4pebbledb.mod 2>&1)
  
  # Check if sync failed (errors from go mod edit or go mod tidy)
  ifneq ($(findstring error,$(SYNC_RESULT)),)
    $(error PebbleDB mod file generation failed: $(SYNC_RESULT))
  endif
  
  $(info ✓ go-4pebbledb.mod generated successfully)
endif

########################################
### Convert Build Tags for Go Build  ###
########################################
# Convert build_tags into a comma-separated list.
build_tags := $(strip $(build_tags))
whitespace :=
empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))

########################################
###   Linker Flags Configuration     ###
########################################
ldflags += -X github.com/cosmos/cosmos-sdk/version.Name=canine \
	-X github.com/cosmos/cosmos-sdk/version.AppName=canined \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X github.com/jackalLabs/canine-chain/app.Bech32Prefix=jkl \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))
BUILD_FLAGS := -tags "$(build_tags_comma_sep)" -ldflags '$(ldflags)' -trimpath

########################################
###         Build Targets            ###
########################################
all: install lint test

# Use the appropriate module file (MODFILE/SUMFILE) for dependency checks and builds.
build: $(SUMFILE)
ifeq ($(OS),Windows_NT)
	exit 1
else
	$(GO_CMD) build -mod=readonly -modfile=$(MODFILE) $(BUILD_FLAGS) -o build/canined ./cmd/canined
endif

build_cli: build

build-contract-tests-hooks:
ifeq ($(OS),Windows_NT)
	$(GO_CMD) build -mod=readonly -modfile=$(MODFILE) $(BUILD_FLAGS) -o build/contract_tests.exe ./cmd/contract_tests
else
	$(GO_CMD) build -mod=readonly -modfile=$(MODFILE) $(BUILD_FLAGS) -o build/contract_tests ./cmd/contract_tests
endif

install: $(SUMFILE)
	$(GO_CMD) install -mod=readonly -modfile=$(MODFILE) $(BUILD_FLAGS) ./cmd/canined

########################################
###       Tools & Dependencies       ###
########################################
go-mod-cache: $(SUMFILE)
	@echo "--> Download go modules to local cache"
	@$(GO_CMD) mod download -modfile=$(MODFILE)

$(SUMFILE): $(MODFILE)
	@echo "--> Ensure dependencies have not been modified"
	@$(GO_CMD) mod verify -modfile=$(MODFILE)

draw-deps:
	@# Requires graphviz (brew install graphviz or apt-get install graphviz)
	$(GO_CMD) get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/canined -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf snapcraft-local.yaml build/

distclean: clean
	rm -rf vendor/

########################################
###            Testing               ###
########################################
local: install
	./scripts/test-node.sh $(address)

test: test-unit
test-all: test-race test-cover
test-sim: test-sim-import-export test-sim-full-app

test-unit:
	@VERSION=$(VERSION) $(GO_CMD) test -short -mod=readonly -modfile=$(MODFILE) -tags='ledger test_ledger_mock' ./...

test-race:
	@VERSION=$(VERSION) $(GO_CMD) test -mod=readonly -modfile=$(MODFILE) -race -tags='ledger test_ledger_mock' ./...

test-cover:
	@$(GO_CMD) test -mod=readonly -modfile=$(MODFILE) -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...

benchmark:
	@$(GO_CMD) test -mod=readonly -modfile=$(MODFILE) -bench=. ./...

test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 5 TestAppImportExport

test-sim-full-app: runsim
	@echo "Running short multi-seed application simulation. This may take awhile!"
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 10 TestFullAppSimulation

test-sim-bench:
	@VERSION=$(VERSION) $(GO_CMD) test -mod=readonly -modfile=$(MODFILE) -benchmem -run ^BenchmarkFullAppSimulation -bench ^BenchmarkFullAppSimulation -cpuprofile cpu.out github.com/jackalLabs/canine-chain/app

runsim:
	$(GO_CMD) install github.com/cosmos/tools/cmd/runsim@latest

########################################
###             Linting              ###
########################################
format-tools:
	$(GO_CMD) install mvdan.cc/gofumpt@latest
	gofumpt -l -w .

lint: format-tools
	golangci-lint run --fix

format: format-tools
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gofumpt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs goimports -w -local github.com/jackalLabs/canine-chain

########################################
###         Protobuf Tools           ###
########################################
# thanks juno ;)
protoVer = v0.7
protoImageName = tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen = jackal-proto-gen-$(protoVer)
containerProtoGenAny = jackal-proto-gen-any-$(protoVer)
containerProtoGenSwagger = jackal-proto-gen-swagger-$(protoVer)
containerProtoFmt = jackal-proto-fmt-$(protoVer)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^$(containerProtoGen)$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) sh ./scripts/protocgen.sh; fi

# This generates the SDK's custom wrapper for google.protobuf.Any.
proto-gen-any:
	@echo "Generating Protobuf Any"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^$(containerProtoGenAny)$$"; then docker start -a $(containerProtoGenAny); else docker run --name $(containerProtoGenAny) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) sh ./scripts/protocgen-any.sh; fi

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^$(containerProtoGenSwagger)$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) sh ./scripts/protoc-swagger-gen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^$(containerProtoFmt)$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=main

# Note: The following targets are declared in .PHONY but have no corresponding rules:
#       install-debug, test-build, proto-update-deps
# They are placeholders for future enhancements.

.PHONY: proto-all proto-gen proto-gen-any proto-swagger-gen proto-format proto-lint proto-check-breaking proto-update-deps docs
.PHONY: all install install-debug go-mod-cache draw-deps clean build format test test-all test-build test-cover test-unit test-race test-sim-import-export local
