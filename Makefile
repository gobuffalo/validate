GO_BIN ?= go
CURL_BIN ?= curl
SHELL_BIN ?= sh

deps: check-gopath
	$(GO_BIN) get -u github.com/Jeffail/gabs
	$(GO_BIN) get -u github.com/gofrs/uuid

	$(CURL_BIN) -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | $(SHELL_BIN) -s -- -b ${GOPATH}/bin v1.15.0
	$(GO_BIN) get -u github.com/Quasilyte/go-consistent

test:
	$(GO_BIN) test -v ./...
	cd validators
	$(GO_BIN) test -v ./...

lint:
	golangci-lint run
	go-consistent -v ./...

check-gopath:
ifndef GOPATH
	$(error GOPATH is undefined)
endif
