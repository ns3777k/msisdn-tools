.PHONY: test install

install:
	go get github.com/golang/lint/golint
	go get github.com/gordonklaus/ineffassign
	go get github.com/client9/misspell/cmd/misspell

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
