.PHONY: test

build: go-get install test

go-get:
	go get golang.org/x/tools/cmd/vet
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/remyoudompheng/go-misc/deadcode
	go get gopkg.in/alecthomas/kingpin.v2
	go get github.com/stretchr/testify

install:
	go get -t -v ./...

test:
	go test -v ./...
	go test -race  -i ./...
	go list ./... | sed -e 's;github.com/mtojek/go-url-fuzzer;.;' | xargs deadcode
	golint ./...
	go vet
	test -z "`gofmt -d .`"
	test -z "`goimports -l .`"

prepare:
	gofmt -s -w .
	goimports -w .

dev: install
	go-url-fuzzer -h "r: 1" -h "br:2" -m "POST" -m "GET" -m "PUT" Makefile http://httbase-url
