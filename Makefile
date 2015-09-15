.PHONY: test

build: go-get install test

go-get:
	go get golang.org/x/tools/cmd/vet
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/opennota/check/cmd/defercheck
	go get github.com/opennota/check/cmd/structcheck
	go get github.com/opennota/check/cmd/varcheck
	go get github.com/mtojek/goflow
	go get gopkg.in/alecthomas/kingpin.v2
	go get github.com/stretchr/testify

install:
	go get -t -v ./...

test:
	go test -v ./...
	go test -race  -i ./...
	go list ./... | sed -e 's;github.com/mtojek/go-url-fuzzer;.;' | xargs defercheck
	go list ./... | sed -e 's;github.com/mtojek/go-url-fuzzer;.;' | xargs varcheck
	go list ./... | sed -e 's;github.com/mtojek/go-url-fuzzer;.;' | xargs structcheck
	golint ./...
	go tool vet -v=true .
	test -z "`gofmt -d .`"
	test -z "`goimports -l .`"

cc: #cleancode
	gofmt -s -w .
	goimports -w .

dev: install
	go-url-fuzzer -h "r: 1" -h "br:2" -m "POST" -m "GET" -m "PUT" Makefile http://httbase-url
