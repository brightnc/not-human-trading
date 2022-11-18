GO_BINARY_NAME=not-human-trading
VERSION=$(shell git describe --tags || git rev-parse --short HEAD || echo "unknown version")
LDFLAGS+= -X "github.com/brightnc/not-human-trading/cmd.Version=$(VERSION)"
LDFLAGS+= -X "github.com/brightnc/not-human-trading/cmd.GoVersion=$(shell go version | sed -r 's/go version go(.*)\ .*/\1/')"

init:
	@echo "== 👩‍🌾 init =="
	brew install node
	brew install pre-commit
	brew install golangci-lint
	brew upgrade golangci-lint

	@echo "== pre-commit setup =="
	pre-commit install

	@echo "== install ginkgo =="
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...

	@echo "== install gomock =="
	go install github.com/golang/mock/mockgen@v1.6.0

# Always turn on go module when use `go` command.
GO := GO111MODULE=on go

# GO build preparation
.PHONY: ci
ci:
	$(GO) mod tidy && \
	$(GO) mod download && \
	$(GO) mod verify && \
	$(GO) mod vendor && \
	$(GO) fmt ./... \

# Build GO application
# -mod=vendor
# tells the go command to use the vendor directory. In this mode,
# the go command will not use the network or the module cache.
# -v
# print the names of packages as they are compiled.
# -a
# force rebuilding of packages that are already up-to-date.
# -o
# -ldsflags
# tells the version and go version.
.PHONY: build
build:
	$(GO) build -ldflags '$(LDFLAGS)' -a -v -o $(GO_BINARY_NAME) main.go

start:
	go run ./main.go serve-rest

mock:
	mockery --output=./mocks/pkgmock --outpkg=pkgmock --dir=./pkg/account --name=Accounter
	mockery --output=./mocks/pkgmock --outpkg=pkgmock --dir=./pkg/timer --name=Timer
	mockery --output=./mocks/repomock --outpkg=repomock --dir=./internal/core/port --name=KKPBankRepo
	mockery --output=./mocks/repomock --outpkg=repomock --dir=./internal/core/port --name=KKPBankClient
	mockery --output=./mocks/repomock --outpkg=repomock --dir=./internal/core/port --name=AMQPRepository

swag:
	swag init -g ./cmd/rest.go -o ./docs

.PHONY: test
test:
	go test ./...

# Clean up when build the application on local directory.
.PHONY: clean
clean:
	@rm -rf $(GO_BINARY_NAME) ./vendor

precommit.rehooks:
	pre-commit autoupdate
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg
