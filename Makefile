
all:
	scripts/build.sh

dist:
	scripts/dist.sh

clean:
	rm bin/my-projectname || true
	rm -rf .gopath || true

test:
	go test ./...

.PHONY: all	dist clean test
