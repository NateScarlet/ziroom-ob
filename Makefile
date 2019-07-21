
all: docs/_build/html

docs/_build/html/.git:
	git fetch -fn origin gh-pages:gh-pages
	rm -rf docs/_build/html/*
	git worktree add -f docs/_build/html gh-pages

.PHONY: all install test deploy-docs

install: */*/*.go
	go get -v ./cmd/ziroom-ob

test:
	go test --timeout 30s ./...

docs/_build/html: docs/_build/html/.git docs/conf.py docs/*.rst
	$(MAKE) -C docs html

deploy-docs: docs/_build/html
	cd docs/_build/html ; git add --all && git commit -m 'docs: build' -m '[skip ci]' && git push
