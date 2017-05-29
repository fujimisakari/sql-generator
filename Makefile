SCHEMA_PATH = ./schema.yaml

.PHONY: setup help create drop example

help:
	@echo "Usage:"
	@echo "  make setup     # install required modules"
	@echo "  make create    # output create table query"
	@echo "  make drop      # output drop table query"
	@echo "  make example   # output example insert query"

setup:
	go get gopkg.in/yaml.v2

create:
	go run output.go const.go model.go main.go $(SCHEMA_PATH) create

example:
	go run output.go const.go model.go main.go $(SCHEMA_PATH) example

drop:
	go run output.go const.go model.go main.go $(SCHEMA_PATH) drop
