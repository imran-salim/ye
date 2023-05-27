all: get build

get:
	@echo "Installing packages and dependencies"
	go get ye

build:
	@echo "Building Ye"
	go build
