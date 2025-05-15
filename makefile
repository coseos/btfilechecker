# Makefile

NAME := btfilechecker
BINARY := build/$(NAME)
SRC := ./cmd/$(NAME)
VERSION := $(shell cat version)

.PHONY: build clean

build:
	mkdir -p $(dir $(BINARY))
	go build -o $(BINARY) $(SRC)

test:
	go test ./...

smoke:
	build/$(NAME) ./cmd/ ./data/$(NAME).lst

pack:
	zip -j build/$(NAME)_$(VERSION).zip build/$(NAME)

clean:
	rm -f $(BINARY)

