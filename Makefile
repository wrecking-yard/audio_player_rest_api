.PHONY: clean
clean:
	go clean -cache

.PHONY: get-swagger-ui
get-swagger-ui:
	mkdir -vp temp/
	mkdir -vp internal/server/swagger-ui
	curl https://registry.npmjs.org/swagger-ui-dist/-/swagger-ui-dist-5.27.0.tgz -o temp/swagger-ui-dist-5.27.0.tgz
	tar -xf temp/swagger-ui-dist-5.27.0.tgz --strip-components=1 -C internal/server/swagger-ui
	cp -v conf/api.json internal/server/swagger-ui/
	patch -R internal/server/swagger-ui/swagger-initializer.js patches/01-change-specs-url.patch
	rm -rfv temp/

.PHONY: clean-swagger-ui
clean-swagger-ui:
	rm -fr internal/server/swagger-ui

.PHONY: build
build:  clean
	go build -o apra

.PHONY: build-debug
build-debug: clean get-swagger-ui
	go build -gcflags=all="-N -l -w"

.PHONY: test
test:
	go test -v ./...
