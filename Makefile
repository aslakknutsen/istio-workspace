PROJECT_NAME:=istio-workspace
PACKAGE_NAME:=github.com/aslakknutsen/istio-workspace

CUR_DIR = $(shell pwd)
BINARY_DIR:=${PWD}/dist

.PHONY: all
all: tools format lint compile ## (default)

.PHONY: help
help:
	 @echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sort | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"

.PHONY: deps
deps: ## Fetches all dependencies using dep
	dep ensure -v

.PHONY: format
format: ## Removes unneeded imports and formats source code
	@goimports -l -w pkg/ cmd/ version/

.PHONY: tools
tools: ## Installs required go tools
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	@go get -u golang.org/x/tools/cmd/goimports

.PHONY: lint
lint: deps ## Concurrently runs a whole bunch of static analysis tools
	@golangci-lint run

.PHONY: codegen
codegen:
    # In case of multiple directories set as GOPATH use the last one defined
	GOPATH=$(shell echo ${GOPATH} | rev | cut -d':' -f1 | rev) \
	vendor/k8s.io/code-generator/generate-groups.sh \
    deepcopy \
    github.com/aslakknutsen/istio-workspace/pkg/generated \
    github.com/aslakknutsen/istio-workspace/pkg/apis \
    istio:v1alpha1 \
    --go-header-file "./go.header.txt"

.PHONY: compile
compile: deps codegen $(BINARY_DIR)/$(PROJECT_NAME)

.PHONY: clean
clean:
	rm -rf $(BINARY_DIR)

# Build configuration
BUILD_TIME=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GITUNTRACKEDCHANGES:=$(shell git status --porcelain --untracked-files=no)
COMMIT:=$(shell git rev-parse --short HEAD)
ifneq ($(GITUNTRACKEDCHANGES),)
  COMMIT := $(COMMIT)-dirty
endif
LDFLAGS="-w -X main.Commit=${COMMIT} -X main.BuildTime=${BUILD_TIME}"

SRCS=$(shell find ./pkg -name "*.go") $(shell find ./cmd -name "*.go")

$(BINARY_DIR):
	[ -d $@ ] || mkdir -p $@

$(BINARY_DIR)/$(PROJECT_NAME): $(BINARY_DIR) $(SRCS)
	GOOS=linux CGO_ENABLED=0 go build -a -tags netgo -ldflags ${LDFLAGS} -o $@ ./cmd/istio-workspace/

# Docker build

DOCKER?=$(if $(or $(in_docker_group),$(is_root)),docker,sudo docker)
DOCKER_IMAGE?=$(PROJECT_NAME)
DOCKER_REPO?=docker.io/aslakknutsen

docker-build: ## Builds the docker image
	@echo "Building docker image $(DOCKER_IMAGE_CORE)"
	$(DOCKER) build \
		-t $(DOCKER_REPO)/$(DOCKER_IMAGE):$(COMMIT) \
		-f $(CUR_DIR)/Dockerfile $(CUR_DIR)