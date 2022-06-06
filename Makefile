.PHONY: all build clean
build:
	go build -o build/ .
test:
	go test -v ./...
install-usr-local:
	make build
	sudo cp build/gowood /usr/local/bin

build-multi-arch:
	./go-build-multiarch.sh github.com/haliliceylan/gowood