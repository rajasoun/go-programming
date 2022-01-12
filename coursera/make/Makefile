APP := $(shell basename $(CURDIR))
 OS := $(shell echo $(shell uname -s) | tr '[:upper:]' '[:lower:]')

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help:	## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: 	## Build 
	go build -v -o bin/${APP}-$(OS)  src/*.go

run: 	## Run 
	bin/${APP}-$(OS)

clean:	## Clean 
	go clean
	rm -fr bin/${APP}-$(OS)


