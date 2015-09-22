package httprequest

import (
	"log"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// URLChecker consumes entries containing relative URL to check and HTTP method.
type URLChecker struct {
	flow.Component

	Entry <-chan messages.Entry
}

// NewURLChecker creates new instance of URL checker.
func NewURLChecker(configuration *configuration.Configuration) *URLChecker {
	return &URLChecker{}
}

// OnEntry perfoms the main URL check.
func (u *URLChecker) OnEntry(entry messages.Entry) {
	log.Println(entry)
}
