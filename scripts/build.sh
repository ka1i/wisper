#!/bin/bash --posix

# Args: source, target
# eg: build.sh ./main.go ./bin/program

TAG=$(echo $([ -n "$(git status -s)"  ] && echo "$(git rev-parse --short HEAD)-dev" || echo "$(git rev-parse --short HEAD)"))
UPT=$(date +"%Y/%m/%d %T %z")
ENV=$(uname -snr)

echo "making $2"
mkdir -p bin
echo "v0.0.1" > .version
echo "Version:${TAG}"

go mod tidy

go build -ldflags "                                     \
    -installsuffix 'static'                             \
    -s -w                                               \
    -X '$(go list -m)/pkg/version.tagStr=${TAG}'\
    -X '$(go list -m)/pkg/version.uptStr=${UPT}'\
    -X '$(go list -m)/pkg/version.envStr=${ENV}'\
    " \
    -o $2 $1
