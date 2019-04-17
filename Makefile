.PHONY: all run clean
all: build

# go list is the canonical utility to find go files
GOFILES := $(shell go list -f '{{ join .GoFiles "\n" }}' ./...)

build: .bin-stamp
	go build -o bin/hue-remote-v2
	chmod +x bin/hue-remote-v2

# directories do werid things in make, so we can use a stamp
.bin-stamp:
	mkdir -p bin
	touch .bin-stamp

# Use 'go run' so we don't have to worry about recompiling
run:
	go run .

clean:
	rm -rf bin
	rm .bin-stamp
