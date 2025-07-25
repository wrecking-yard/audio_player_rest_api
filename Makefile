.PHONY: clean
clean:
	go clean -cache

.PHONY: build
build: clean
	go build -o apra

.PHONY: build-debug
build-debug: clean
	go build -gcflags=all="-N -l -w"

.PHONY: test
test:
	go test -v ./...
