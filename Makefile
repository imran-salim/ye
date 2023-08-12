github := github.com
username := i8abyte
repo := ye
package := $(github)/$(username)/$(repo)

all: get clean build

get:
	@echo "Installing dependencies"
	go get $(package)

clean:
	@echo "Removing binaries"
	go clean

build:
	@echo "Building Ye"
	go build
