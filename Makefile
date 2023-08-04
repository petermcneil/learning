.PHONY: test
BUILD_FLAGS+=-tags osusergo,netgo -ldflags="-s -w -extldflags=-static"

test:
	@go fmt ./...
	@go vet ./...
	@CGO_ENABLED=0 go test ${BUILD_FLAGS} ./... -p 4