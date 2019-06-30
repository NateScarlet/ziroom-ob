.PHONY: all install test docs

all:

install: *.go */*.go
	go get -v ./cmd/ziroom-ob

test:
	go test --timeout 30s ./...

docs:
	$(MAKE) -C docs html
