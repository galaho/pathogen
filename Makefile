VERSIONS_PACKAGE := github.com/galaho/pathogen/versions

COMMIT := $(shell git rev-parse --verify --short HEAD 2> /dev/null || echo "UNKNOWN")
COMMIT_FLAG := -X $(VERSIONS_PACKAGE).commit=$(COMMIT)

VERSION := $(shell cat VERSION || echo "UNKNOWN")
VERSION_FLAG := -X $(VERSIONS_PACKAGE).version=$(VERSION)

.PHONY: build
build: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 go build -o bin/pathogen -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: build.darwin
build.darwin: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/pathogen.darwin -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: build.linux
build.linux: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/pathogen.linux -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: build.windows
build.windows: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/pathogen.windows -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor

.PHONY: test
test: vendor
	@echo "--> Running tests..."
	@CGO_ENABLED=0 go test -v --coverprofile=./coverage/c.out --mod=vendor ./...

.PHONY: vendor
vendor:
	@echo "--> Vendoring dependencies..."
	@CGO_ENABLED=0 go mod vendor