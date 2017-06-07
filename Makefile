.PHONY: setup build

help:
	@echo "Usage:"
	@echo "  make setup   # install required modules"
	@echo "  make build   # sql-generator build"

setup:
	go get gopkg.in/yaml.v2

build:
	go build
