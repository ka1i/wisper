NAME    := wisper
SOURCE  := cmd/${NAME}/main.go

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always)
UPDATE  := $(shell date +"%Y.%m.%d %X")

BINARY  := bin/wisper_${VERSION}/${NAME}

all: build

.PHONY: build
build:        ## build this app.
	@echo "making ${BINARY}"
	@mkdir -p bin/wisper_${VERSION}
	@echo "${NAME} Version:${VERSION}\n${NAME} Last Update:${UPDATE}"
	@go build -ldflags "                                 		\
        -installsuffix 'static'                                 \
        -s -w                                                   \
        -X '$(shell go list -m)/pkg/version.VERSION=${VERSION}' \
        -X '$(shell go list -m)/pkg/version.UPDATE=${UPDATE}'   \
        "                                                       \
        -o ${BINARY} ${SOURCE}

.PHONY: next
next:           ## Build frontend.
	rm -r pkg/assets/web/*
	@cd web; yarn ; yarn static ; cp -r out/* ../pkg/assets/web/;cd ../

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