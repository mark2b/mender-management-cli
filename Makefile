MAKEFILE_PATH:=$(word $(words $(MAKEFILE_LIST)),$(MAKEFILE_LIST))
ROOT:=$(shell cd $(dir $(MAKEFILE_PATH));pwd)
SRC_DIR:=$(ROOT)
BIN_DIR:=$(ROOT)/bin

VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

LDFLAGS = "-X mender-management-cli/app.Version=$(VERSION) -X jema-artifact-gen/app.BuildTime=$(BUILD)"

all: build


build:
	cd $(SRC_DIR);GOOS=linux GOARCH=arm go build -ldflags=$(LDFLAGS) -o $(BIN_DIR)/mender-management-cli

clean:
	rm -f $(BIN_DIR)/*

