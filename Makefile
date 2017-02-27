
package = github.com/snoby/mkv2Appletv


MAKE=make
BIN=mkv2Appletv
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Build=${BUILD}"

.DEFAULT_TARGET: ${BIN}


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

.PHONY: clean install release



