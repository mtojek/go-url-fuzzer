.PHONY: prepare test

build: go-get install test

go-get:
	go get golang.org/x/tools/cmd/vet
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get gopkg.in/alecthomas/kingpin.v2 

install:
	go get -t -v ./...

test:
	go test -v ./...
	golint
	go vet
	test -z "`gofmt -d .`"
	test -z "`goimports -l .`"

prepare:
	golint .
	go vet
	gofmt -s -w .
	goimports -w .
