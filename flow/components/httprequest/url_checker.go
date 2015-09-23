package httprequest

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// URLChecker consumes entries containing relative URL to check and HTTP method.
type URLChecker struct {
	flow.Component

	Entry      <-chan messages.Entry
	FoundEntry chan<- messages.Entry

	client *http.Client

	headers       http.Header
	httpErrorCode int
	waitPeriod    time.Duration
}

// NewURLChecker creates new instance of URL checker.
func NewURLChecker(configuration *configuration.Configuration) *URLChecker {
	return &URLChecker{
		client:        createHTTPClient(configuration),
		headers:       createHTTPHeaders(configuration),
		httpErrorCode: int(configuration.HTTPErrorCode()),
		waitPeriod:    configuration.WorkerWaitPeriod(),
	}
}

func createHTTPClient(configuration *configuration.Configuration) *http.Client {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr, Timeout: configuration.URLResponseTimeout()}
	return client
}

func createHTTPHeaders(configuration *configuration.Configuration) http.Header {
	var preparedHeaders http.Header

	if headers, exists := configuration.Headers(); exists {
		for name, value := range headers {
			preparedHeaders[name] = []string{value}
		}
	}

	return preparedHeaders
}

// OnEntry perfoms the main URL check.
func (u *URLChecker) OnEntry(entry messages.Entry) {
	request, error := http.NewRequest(entry.HTTPMethod(), entry.RelativeURL(), nil)
	if nil != error {
		log.Fatalf("Could not create a new request (method: %v, URL: %v), error: %v", entry.HTTPMethod(), entry.RelativeURL(), error)
	}
	request.Header = u.headers

	response, error := u.client.Do(request)
	if nil == error {
		if response.StatusCode != u.httpErrorCode {
			u.FoundEntry <- entry
		}
	}

	u.waitIfNecessary()
}

func (u *URLChecker) waitIfNecessary() {
	if u.waitPeriod > 0 {
		time.Sleep(u.waitPeriod)
	}
}
