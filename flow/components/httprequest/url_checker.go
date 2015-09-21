package httprequest

import (
	"log"

	"github.com/mtojek/go-url-fuzzer/flow/components/httpmethod"
	"github.com/trustmaster/goflow"
)

// URLChecker consumes entries containing relative URL to check and HTTP method.
type URLChecker struct {
	flow.Component

	Entry <-chan httpmethod.Entry
}

// NewURLChecker creates new instance of URL checker.
func NewURLChecker(configuration urlCheckerConfiguration) *URLChecker {
	return &URLChecker{}
}

// OnEntry perfoms the main URL check.
func (u *URLChecker) OnEntry(entry httpmethod.Entry) {
	log.Println(entry)
}

// TODO inne port
// TODO zdefiniuj bledna strone z parametru, default 404
