export CGO_ENABLED = 0

export GO111MODULE := on

project_dir = $(shell pwd)
builds_dir = $(shell echo $${builds_dir:-$(project_dir)/builds})

UNAME_S := $(shell uname -s)
UNAME_P := $(shell uname -p)

GOOS = linux
ifeq ($(UNAME_S),Darwin)
  GOOS = darwin
endif

GOARCH = amd64
ifneq ($(UNAME_P),x86_64)
  GOARCH = 386
endif

.PHONY: _build
_build:
	@echo "=> building dockma via go build"
	@echo ""
	@builds_dir=${builds_dir} GOOS=${GOOS} GOARCH=${GOARCH} go build -o "$(builds_dir)/dockma_$(GOOS)_$(GOARCH)" cmd/dockma.go
	@echo "built $(builds_dir)/doctl_$(GOOS)_$(GOARCH)"

.PHONY: native
native: _build
	@echo "==> build local version"
	@echo ""
	@mv $(builds_dir)/doctl_$(GOOS)_$(GOARCH) $(builds_dir)/doctl
	@echo "installed as $(builds_dir)/doctl"

.PHONY: clean
clean:
	@echo "==> remove builds"
	@echo ""
	@rm -rf builds

.PHONY: example
example:
	@echo "==> test"
	@echo ""
	@echo "built $(builds_dir)/doctl_$(GOOS)_$(GOARCH)"
