VERSIONS_PACKAGE := github.com/galaho/pathogen/versions
ROOT := $(shell git rev-parse --show-toplevel)

COMMIT := $(shell git rev-parse --verify --short HEAD 2> /dev/null || echo "UNKNOWN")
COMMIT_FLAG := -X $(VERSIONS_PACKAGE).commit=$(COMMIT)

VERSION := $(shell cat VERSION || echo "UNKNOWN")
VERSION_FLAG := -X $(VERSIONS_PACKAGE).version=$(VERSION)

.PHONY: build
build: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 go build -o bin/pathogen -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor ./cmd/pathogen.go

.PHONY: build.darwin
build.darwin: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/pathogen.darwin -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor ./cmd/pathogen.go

.PHONY: build.linux
build.linux: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/pathogen.linux -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor ./cmd/pathogen.go

.PHONY: build.windows
build.windows: vendor
	@echo "--> Building binary..."
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/pathogen.windows -ldflags "$(VERSION_FLAG) $(COMMIT_FLAG)" --mod=vendor ./cmd/pathogen.go

.PHONY: test
test: vendor
	@echo "--> Running tests..."
	@CGO_ENABLED=0 go test -v --coverprofile=./coverage/c.out --mod=vendor ./...

.PHONY: vendor
vendor:
	@echo "--> Vendoring dependencies..."
	@CGO_ENABLED=0 go mod vendor

.PHONY: ci
ci:
	@echo "--> Testing the CI..."
	@docker run --rm -v $(ROOT):/go/src/github.com/galaho/pathogen -w /go/src/github.com/galaho/pathogen goreleaser/goreleaser release --snapshot
