.PHONY: test

build: go-get install test

go-get:
	go get golang.org/x/tools/cmd/vet
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/trustmaster/goflow
	go get gopkg.in/alecthomas/kingpin.v2
	go get github.com/stretchr/testify
	go get github.com/mtojek/localserver

install:
	go get -t -v ./...

test:
	go test -v ./...
	go test -race  -i ./...
	golint ./...
	go tool vet -v=true .
	test -z "`gofmt -d .`"
	test -z "`goimports -l .`"

cc: #cleancode
	gofmt -s -w .
	goimports -w .

dev: install
	go-url-fuzzer -h "a: 1" -h "b: 2" -m "POST" -m "GET" -m "PUT" resources/input-data/fuzz_02.txt http://www.wp.pl/
