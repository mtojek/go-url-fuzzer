package httprequest

import (
	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// URLChecker consumes entries containing relative URL to check and HTTP method.
type URLChecker struct {
	flow.Component

	Entry      <-chan messages.Entry
	FoundEntry chan<- messages.Entry

	configuration *configuration.Configuration
}

// NewURLChecker creates new instance of URL checker.
func NewURLChecker(configuration *configuration.Configuration) *URLChecker {
	return &URLChecker{configuration: configuration}
}

// OnEntry perfoms the main URL check.
func (u *URLChecker) OnEntry(entry messages.Entry) {
	u.FoundEntry <- entry
}
