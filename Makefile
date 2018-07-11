.PHONY: deps build ins install-bin

all: deps build

deps:
	@echo "Downloading dependencies"
	@/bin/sh -c "./scripts/deps.sh"

build:
	@echo "Building ultraslackbot"
	@/bin/sh -c "./scripts/build.sh"

ins: deps build install-bin

install-bin:
	@echo "Installing ultraslackbot"
	@/bin/sh -c "./scripts/install-bin.sh"

docker-up:
	@echo "Wait..."
	@docker-compose up -d
