ifndef GOOS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	GOOS := darwin
else ifeq ($(UNAME_S),Linux)
	GOOS := linux
else
$(error "$$GOOS is not defined. If you are using Windows, try to re-make using 'GOOS=windows make ...' ")
endif
endif

PACKAGES    := $(shell go list ./... | grep -v '/lib/')

BUILD_FLAGS := -ldflags "-X github.com/bytom/bystack/version.GitCommit=`git rev-parse HEAD`"

BYSTACKD_BINARY32 := bystackd-$(GOOS)_386
BYSTACKD_BINARY64 := bystackd-$(GOOS)_amd64

BYSTACKCLI_BINARY32 := bystackcli-$(GOOS)_386
BYSTACKCLI_BINARY64 := bystackcli-$(GOOS)_amd64

VERSION := $(shell awk -F= '/Version =/ {print $$2}' version/version.go | tr -d "\" ")

BYSTACKD_RELEASE32 := bystackd-$(VERSION)-$(GOOS)_386
BYSTACKD_RELEASE64 := bystackd-$(VERSION)-$(GOOS)_amd64

BYSTACKCLI_RELEASE32 := bystackcli-$(VERSION)-$(GOOS)_386
BYSTACKCLI_RELEASE64 := bystackcli-$(VERSION)-$(GOOS)_amd64

BYSTACK_RELEASE32 := bystack-$(VERSION)-$(GOOS)_386
BYSTACK_RELEASE64 := bystack-$(VERSION)-$(GOOS)_amd64

all: test target release-all install

bystackd:
	@echo "Building bystackd to cmd/bystackd/bystackd"
	@go build $(BUILD_FLAGS) -o cmd/bystackd/bystackd cmd/bystackd/main.go

bystackcli:
	@echo "Building bystackcli to cmd/bystackcli/bystackcli"
	@go build $(BUILD_FLAGS) -o cmd/bystackcli/bystackcli cmd/bystackcli/main.go

install:
	@echo "Installing bystackd and bystackcli to $(GOPATH)/bin"
	@go install ./cmd/bystackd
	@go install ./cmd/bystackcli

target:
	mkdir -p $@

binary: target/$(BYSTACKD_BINARY32) target/$(BYSTACKD_BINARY64) target/$(BYSTACKCLI_BINARY32) target/$(BYSTACKCLI_BINARY64)

ifeq ($(GOOS),windows)
release: binary
	cd target && cp -f $(BYSTACKD_BINARY32) $(BYSTACKD_BINARY32).exe
	cd target && cp -f $(BYSTACKCLI_BINARY32) $(BYSTACKCLI_BINARY32).exe
	cd target && md5sum  $(BYSTACKD_BINARY32).exe $(BYSTACKCLI_BINARY32).exe >$(BYSTACK_RELEASE32).md5
	cd target && zip $(BYSTACK_RELEASE32).zip  $(BYSTACKD_BINARY32).exe $(BYSTACKCLI_BINARY32).exe $(BYSTACK_RELEASE32).md5
	cd target && rm -f  $(BYSTACKD_BINARY32) $(BYSTACKCLI_BINARY32)  $(BYSTACKD_BINARY32).exe $(BYSTACKCLI_BINARY32).exe $(BYSTACK_RELEASE32).md5
	cd target && cp -f $(BYSTACKD_BINARY64) $(BYSTACKD_BINARY64).exe
	cd target && cp -f $(BYSTACKCLI_BINARY64) $(BYSTACKCLI_BINARY64).exe
	cd target && md5sum  $(BYSTACKD_BINARY64).exe $(BYSTACKCLI_BINARY64).exe >$(BYSTACK_RELEASE64).md5
	cd target && zip $(BYSTACK_RELEASE64).zip  $(BYSTACKD_BINARY64).exe $(BYSTACKCLI_BINARY64).exe $(BYSTACK_RELEASE64).md5
	cd target && rm -f  $(BYSTACKD_BINARY64) $(BYSTACKCLI_BINARY64)  $(BYSTACKD_BINARY64).exe $(BYSTACKCLI_BINARY64).exe $(BYSTACK_RELEASE64).md5
else
release: binary
	cd target && md5sum  $(BYSTACKD_BINARY32) $(BYSTACKCLI_BINARY32) >$(BYSTACK_RELEASE32).md5
	cd target && tar -czf $(BYSTACK_RELEASE32).tgz  $(BYSTACKD_BINARY32) $(BYSTACKCLI_BINARY32) $(BYSTACK_RELEASE32).md5
	cd target && rm -f  $(BYSTACKD_BINARY32) $(BYSTACKCLI_BINARY32) $(BYSTACK_RELEASE32).md5
	cd target && md5sum  $(BYSTACKD_BINARY64) $(BYSTACKCLI_BINARY64) >$(BYSTACK_RELEASE64).md5
	cd target && tar -czf $(BYSTACK_RELEASE64).tgz  $(BYSTACKD_BINARY64) $(BYSTACKCLI_BINARY64) $(BYSTACK_RELEASE64).md5
	cd target && rm -f  $(BYSTACKD_BINARY64) $(BYSTACKCLI_BINARY64) $(BYSTACK_RELEASE64).md5
endif

release-all: clean
	GOOS=darwin  make release
	GOOS=linux   make release
	GOOS=windows make release

clean:
	@echo "Cleaning binaries built..."
	@rm -rf cmd/bystackd/bystackd
	@rm -rf cmd/bystackcli/bystackcli
	@rm -rf target
	@rm -rf $(GOPATH)/bin/bystackd
	@rm -rf $(GOPATH)/bin/bystackcli
	@echo "Cleaning temp test data..."
	@rm -rf test/pseudo_hsm*
	@rm -rf blockchain/pseudohsm/testdata/pseudo/
	@echo "Cleaning sm2 pem files..."
	@rm -rf crypto/sm2/*.pem
	@echo "Done."

target/$(BYSTACKD_BINARY32):
	CGO_ENABLED=0 GOARCH=386 go build $(BUILD_FLAGS) -o $@ cmd/bystackd/main.go

target/$(BYSTACKD_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/bystackd/main.go

target/$(BYSTACKCLI_BINARY32):
	CGO_ENABLED=0 GOARCH=386 go build $(BUILD_FLAGS) -o $@ cmd/bystackcli/main.go

target/$(BYSTACKCLI_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/bystackcli/main.go

test:
	@echo "====> Running go test"
	@go test $(PACKAGES)

benchmark:
	@go test -bench $(PACKAGES)

functional-tests:
	@go test -timeout=5m -tags="functional" ./test 

ci: test

.PHONY: all target release-all clean test benchmark
