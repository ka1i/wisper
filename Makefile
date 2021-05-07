NAME    := wisper
SOURCE  := cmd/${NAME}/main.go
BINARY  := bin/${NAME}

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always)
UPDATE  := $(shell date +"%Y.%m.%d %X")

all: build

.PHONY: build
build:        ## build this app.
	@echo "making ${BINARY}"
	@mkdir -p bin
	@echo "${NAME} Version:${VERSION}\n${NAME} Last Update:${UPDATE}"
	@go build -ldflags "                                 		\
        -installsuffix 'static'                                 \
        -s -w                                                   \
        -X '$(shell go list -m)/pkg/version.VERSION=${VERSION}' \
        -X '$(shell go list -m)/pkg/version.UPDATE=${UPDATE}'   \
        "                                                       \
        -o ${BINARY} ${SOURCE}

.PHONY: help
help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: version
version:        ## Show the app version.
	@echo "Version :$(VERSION)"
	@echo "Last Update :$(UPDATE)"

.PHONY: clean
clean:          ## Clean build cache.
	@rm -rf bin
	@echo "clean [ ok ]"