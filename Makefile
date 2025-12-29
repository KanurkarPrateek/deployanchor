.PHONY: build

build: build-backend build-cli

build-backend:
	@echo "Building backend..."
	@go build -o bin/backend ./backend

build-cli:
	@echo "Building cli..."
	@go build -o bin/anchor ./cli

install:
	@echo "Installing anchor cli..."
	@cp bin/anchor /usr/local/bin/anchor
