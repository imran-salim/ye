github := github.com
username := imran-salim
repo := ye
package := $(github)/$(username)/$(repo)

all: get build

get:
	@echo "Ready to work."
	@echo "Install package."
	go get $(package)
	@echo "Job's done."

build:
	@echo "Ready to work."
	@echo "Build Ye."
	go build
	@echo "Job's done."

