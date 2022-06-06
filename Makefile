build:
	go build .
test:
	go test -v ./...

install-usr-local:
	make build
	sudo cp gowood /usr/local/bin