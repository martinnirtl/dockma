export CGO_ENABLED = 0

export GO111MODULE := on

.PHONY: build
build:
	@echo "building dockma..."
	@go build -mod vendor -v -o "builds/dockma" cmd/dockma.go
	@echo "built dockma 🐳"

.PHONY: clean
clean:
	@echo "cleaning up..."
	@rm -vrf builds dist
	@echo "done"
