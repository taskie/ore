.PHONY: build install test fmt coverage dep-init dep-ensure dep-graph pre-commit install-pre-commit

VERSION := $(shell git describe --tags)
COMMIT := $(shell git rev-parse HEAD)
DATE := $(shell env TZ= date --iso-8601=seconds)
LDFLAGS := "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

build:
	go build -v -ldflags $(LDFLAGS)

install:
	go install -v -ldflags $(LDFLAGS)

test:
	go test ./...

fmt:
	find . -name '*.go' | grep -v ./vendor/ | xargs gofmt -w

coverage:
	mkdir -p test/coverage
	go test -coverprofile=test/coverage/cover.out ./...
	go tool cover -html=test/coverage/cover.out -o test/coverage/cover.html

dep-init:
	-rm -rf vendor/
	-rm -f Gopkg.toml Gopkg.lock
	dep init

dep-ensure:
	dep ensure

dep-graph: images/dependency.png

images/dependency.png: Gopkg.lock
	mkdir -p images
	dep status -dot | grep -v '^The status of ' | dot -Tpng -o images/dependency.png

pre-commit:
	$(MAKE) dep-ensure
	$(MAKE) fmt
	$(MAKE) build
	$(MAKE) coverage

install-pre-commit:
	echo 'make pre-commit' >.git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
