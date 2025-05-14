# Makefile

NAME = btfilechecker
BINARY = build/$(NAME)
SRC = ./cmd/$(NAME)

.PHONY: build clean

build:
	mkdir -p $(dir $(BINARY))
	go build -o $(BINARY) $(SRC)

test:
	go test ./...

smoke:
	build/$(NAME) ./cmd/ ./data/$(NAME).lst

clean:
	rm -f $(BINARY)

