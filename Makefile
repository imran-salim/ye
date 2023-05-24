
all: get build

get:
	@echo "Installing packages and dependencies"
	go get ye-quote

build:
	@echo "Building ye-quote"
	go build