IMPORT_PATH := github.com/123shang60/spnego-proxy/internal
PACKAGE := spnego-proxy
IMG ?= spnego-proxy:latest
PLATFORM ?= linux/amd64,linux/arm64

# params
TIME			:= $(shell date '+%Y-%m-%d %H:%M:%S')
BRANCH			:= $(shell git symbolic-ref --short -q HEAD)
SERIAL			:= $(shell git rev-parse --short HEAD)
VERSION			:= $(shell git rev-parse --abbrev-ref HEAD)
GOVERSION		:= $(shell go version)
COMMITID		:= $(shell git log  -1 --pretty=format:"%h")
COMMITDATE		:= $(shell git show -s --format=%ci)

GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BUILDPATH=build
BINPATH=$(BUILDPATH)

FLAGS = -p 10 -ldflags='-w -s -X "$(IMPORT_PATH)/common.Version=$(VERSION)" -X "$(IMPORT_PATH)/common.BuildTime=$(TIME)" -X "$(IMPORT_PATH)/common.Branch=$(BRANCH)" -X "$(IMPORT_PATH)/common.CommitId=$(COMMITID)" -X "$(IMPORT_PATH)/common.CommitDate=$(COMMITDATE)" -X "$(IMPORT_PATH)/common.GoVersion=$(GOVERSION)"'
TAG = "unknown"

ifeq ($(ENV), )
	ENV=debug
endif

all: build

clean:
	@rm -rf $(BUILDPATH)

.PHONY: build
build:
	@$(GOMOD) tidy
	@echo $(FLAGS)
	@CGO_ENABLED=0 $(GOBUILD) $(FLAGS) -o $(BINPATH)/spnego-proxy main.go 
	@echo "build spnego-proxy module done"

container: ## Build and Push image
	docker buildx build -f Dockerfile -t ${IMG} --platform=${PLATFORM} --push .