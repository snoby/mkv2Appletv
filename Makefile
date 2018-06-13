
package = github.com/snoby/mkv2Appletv


MAKE=make
BIN=mkv2Appletv
BUILD=`git rev-parse HEAD`
VERSION='0.0.6'

LDFLAGS=-ldflags "-X main.Build=${BUILD}"

# this makefile uses gothub so your env needs to have the personal access token to upload the release files.

.DEFAULT_TARGET: ${BIN}

tag:
	git tag ${VERSION} 
	git push --tags

github_release: clean install release tag 
	gothub release --tag ${VERSION} 
	gothub upload --file release/mkv2Appletv-darwin-amd64 --name mkv2Appletv-darwin-amd64
	gothub upload --file release/mkv2Appletv-linux-amd64 --name mkv2Appletv-linux-amd64

build:
	go build

install: build
	go install

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/mkv2Appletv-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o release/mkv2Appletv-darwin-amd64


check:
	golint
	go vet -v


clean:
	if [ -f ${BIN} ] ; then rm ${BIN}; fi

.PHONY: clean install release .github_release



