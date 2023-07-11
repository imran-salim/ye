all: get build

get:
	@echo "Installing packages and dependencies"
	go get github.com/narmilas/ye

build:
	@echo "Building Ye"
	go build
