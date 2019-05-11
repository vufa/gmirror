IMPORT := github.com/countstarlight/gmirror
 
GO ?= go
SED_INPLACE := sed -i
EXTRA_GOFLAGS ?=

ifeq ($(OS), Windows_NT)
	EXECUTABLE := gmirror.exe
	#EXTRA_GOFLAGS = -tags netgo -ldflags '-H=windowsgui -extldflags "-static" -s'
else
	EXECUTABLE := gmirror
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		SED_INPLACE := sed -i ''
		#EXTRA_GOFLAGS = -ldflags '-s -extldflags "-sectcreate __TEXT __info_plist Info.plist"'
	else
		#EXTRA_GOFLAGS = -tags netgo -ldflags '-extldflags "-static" -s'
	endif
endif

GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*")
GOBINS := ${GOPATH}/bin
GOFMT ?= gofmt -s

GOFLAGS := -mod=vendor -v

PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f)

.PHONY: all
all: build

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -f $(EXECUTABLE)

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	# get all go files and run go fmt on them
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: build
build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	$(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -o $@
