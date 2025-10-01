github := github.com
username := imran-salim
repo := ye
package := $(github)/$(username)/$(repo)

.PHONY: all get test build clean run

all: get test build

get:
	@echo "Ready to work."
	@echo "Install package."
	go get $(package)
	@echo "Job's done."

test:
	@echo "Ready to work."
	@echo "Test Ye."
	go test -v ./test
	@echo "Job's done."

build:
	@echo "Ready to work."
	@echo "Build Ye."
	go build .
	@echo "Job's done."

clean:
	@echo "Ready to work."
	@echo "Clean Ye."
	go clean
	@echo "Job's done."

run: build
	@echo "Ready to work."
	@echo "Run Ye."
	./ye
	@echo "Job's done."

