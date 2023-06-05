all: get build

get:
	@echo "Installing packages and dependencies"
	go get github.com/imran-salim/ye

build:
	@echo "Building Ye"
	go build
