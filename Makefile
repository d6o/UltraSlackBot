.PHONY: clean deps plugins build ins install-bin

all: clean deps plugins build

plugins:
	@echo "Building plugins"
	@/bin/sh -c "./scripts/plugins.sh"

clean:
	@echo "Cleaning workspace"
	@/bin/sh -c "./scripts/clean.sh"

deps:
	@echo "Downloading dependencies"
	@/bin/sh -c "./scripts/deps.sh"

build: clean
	@echo "Building ultraslackbot"
	@/bin/sh -c "./scripts/build.sh"

ins: clean deps plugins build install-bin

install-bin:
	@echo "Installing ultraslackbot"
	@/bin/sh -c "./scripts/install-bin.sh"

docker-up:
	@echo "Wait..."
	@docker-compose up -d
