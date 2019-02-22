GO_BIN ?= go
CURL_BIN ?= curl
SHELL_BIN ?= sh

deps: check-gopath
	$(GO_BIN) get -u github.com/Jeffail/gabs
	$(GO_BIN) get -u github.com/gofrs/uuid

	$(CURL_BIN) -sfL https://raw.githubusercontent.com/alecthomas/gometalinter/master/scripts/install.sh | $(SHELL_BIN) -s -- -b $GOPATH/bin
	$(CURL_BIN) -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | $(SHELL_BIN) -s -- -b $GOPATH/bin v1.12.2
	$(GO_BIN) get -u github.com/go-critic/go-critic/...
	$(GO_BIN) get -u github.com/Quasilyte/go-consistent

test:
	$(GO_BIN) test -v ./...
	cd validators
	$(GO_BIN) test -v ./...

lint:
	golangci-lint run
	gometalinter

	cd validators
	golangci-lint run
	gometalinter

check-gopath:
ifndef GOPATHs
	$(error GOPATH is undefined)
endif
