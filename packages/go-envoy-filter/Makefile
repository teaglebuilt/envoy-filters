ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
FILTER_DIR := $(ROOT_DIR)/filters/go-envoy-filter

build/go-envoy-filter:
	tinygo build -o $(FILTER_DIR)/go-envoy-filter.wasm -target=wasi $(FILTER_DIR)/main.go

tidy/go-envoy-filter:
	go mod tidy