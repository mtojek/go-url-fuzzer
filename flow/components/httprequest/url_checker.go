package httprequest

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"net/url"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// URLChecker consumes entries containing relative URL to check and HTTP method.
type URLChecker struct {
	flow.Component

	Entry      <-chan messages.Entry
	FoundEntry chan<- messages.FoundEntry

	client        *http.Client
	baseURL       url.URL
	headers       http.Header
	httpErrorCode int
	waitPeriod    time.Duration
}

// NewURLChecker creates new instance of URL checker.
func NewURLChecker(configuration *configuration.Configuration) *URLChecker {
	return &URLChecker{
		client:        createHTTPClient(configuration),
		baseURL:       configuration.BaseURL(),
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
	preparedHeaders := http.Header{}

	if headers, exists := configuration.Headers(); exists {
		for name, value := range headers {
			preparedHeaders[name] = []string{value}
		}
	}

	return preparedHeaders
}

// OnEntry perfoms the main URL check.
func (u *URLChecker) OnEntry(entry messages.Entry) {
	absoluteURL, error := u.baseURL.Parse(entry.RelativeURL())
	if nil != error {
		log.Fatalf("Could not build absolute URL, base URL: %v, relative URL: %v, error: %v", u.baseURL, entry.RelativeURL(), error)
	}

	request := u.createRequest(entry.HTTPMethod(), absoluteURL.String())
	response, error := u.client.Do(request)
	if nil == error {
		defer func() {
			if nil != response.Body {
				response.Body.Close()
			}
		}()

		if response.StatusCode != u.httpErrorCode {
			foundEntry := messages.NewFoundEntry(absoluteURL.String(), entry.HTTPMethod(), response.StatusCode)
			u.FoundEntry <- foundEntry
		}
	}

	u.waitIfNecessary()
}

func (u *URLChecker) createRequest(method string, absoluteURL string) *http.Request {
	if request, error := http.NewRequest(method, absoluteURL, nil); nil != error {
		log.Fatalf("Could not create request, URL: %v, error: %v", absoluteURL, error)
	} else {
		request.Header = u.headers
		return request
	}
	return nil
}

func (u *URLChecker) waitIfNecessary() {
	if u.waitPeriod > 0 {
		time.Sleep(u.waitPeriod)
	}
}
