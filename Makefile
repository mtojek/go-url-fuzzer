.PHONY: test

build: go-get install test

go-get:
	go get github.com/trustmaster/goflow
	go get gopkg.in/alecthomas/kingpin.v2
	go get github.com/stretchr/testify
	go get github.com/stretchr/testify/assert
	go get github.com/mtojek/localserver
	go get github.com/alecthomas/gometalinter && gometalinter --install

install:
	go get -t -v ./...

test:
	go test -v ./...
	go test -race  -i ./...
	gometalinter --disable-all --enable=vet --enable=golint --enable=goimports --enable=gofmt .

cc: #cleancode
	gofmt -s -w .
	goimports -w .

dev: install
	go-url-fuzzer -h "a: 1" -h "b: 2" -m "POST" -m "GET" -m "PUT" resources/input-data/fuzz_02.txt http://www.wp.pl/
