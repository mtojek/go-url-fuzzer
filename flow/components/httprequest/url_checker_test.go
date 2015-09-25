package httprequest

import (
	"testing"

	"time"

	"log"
	"net/url"

	"net/http"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/mtojek/localserver"
	"github.com/stretchr/testify/assert"
)

func TestNewURLCheckerClient(t *testing.T) {
	assert := assert.New(t)

	// given
	address := "http://localhost:10605"
	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	expectedURLResponseTimeout := 99 * time.Second
	builder := configuration.NewBuilder()
	configuration := builder.
		URLResponseTimeout(expectedURLResponseTimeout).
		WorkerWaitPeriod(0).
		HTTPErrorCode(http.StatusNotFound).
		BaseURL(url).
		Build()

	// when
	sut := NewURLChecker(configuration)

	// then
	assert.NotNil(sut, "URL checker instance should be created")
	assert.NotNil(sut.client, "HTTP client should be set")
	assert.Equal(expectedURLResponseTimeout, sut.client.Timeout, "URL response timeout should be same as given")
}

func TestNewURLCheckerProperties(t *testing.T) {
	assert := assert.New(t)

	// given
	address := "http://localhost:10605"
	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	expectedWorkerWaitPeriod := 0 * time.Second
	expectedHTTPErrorCode := http.StatusNotFound

	builder := configuration.NewBuilder()
	configuration := builder.
		URLResponseTimeout(3 * time.Second).
		WorkerWaitPeriod(expectedWorkerWaitPeriod).
		HTTPErrorCode(uint64(expectedHTTPErrorCode)).
		BaseURL(url).
		Build()

	// when
	sut := NewURLChecker(configuration)

	// then
	assert.NotNil(sut, "URL checker instance should be created")
	assert.Equal(*url, sut.baseURL, "Base URL should be same as given")
	assert.Equal(expectedHTTPErrorCode, sut.httpErrorCode, "HTTP error code should be same as given")
	assert.Equal(expectedHTTPErrorCode, sut.httpErrorCode, "HTTP error code should be same as given")
	assert.Equal(expectedWorkerWaitPeriod, sut.waitPeriod, "Worker wait period should be same as given")
}

func TestOnEntryNoURLsFound(t *testing.T) {
	assert := assert.New(t)

	// given
	address := "http://localhost:10606"
	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	firstEntry := messages.NewEntry("/aaa", "GET")
	secondEntry := messages.NewEntry("/bbb", "POST")

	builder := configuration.NewBuilder()
	configuration := builder.
		URLResponseTimeout(3 * time.Second).
		WorkerWaitPeriod(0).
		HTTPErrorCode(uint64(http.StatusNotFound)).
		BaseURL(url).
		Build()
	sut := NewURLChecker(configuration)
	foundEntries := make(chan messages.FoundEntry, 4)
	assignChannel(sut, foundEntries)

	// when
	sut.OnEntry(firstEntry)
	sut.OnEntry(secondEntry)

	// then
	assert.Len(foundEntries, 0, "No entries should be considered as found")
}

func TestOnEntryURLsHTTPS(t *testing.T) {
	assert := assert.New(t)

	// given
	scheme := "https"
	hostPort := "127.0.0.1:10607"
	server := localserver.NewLocalServer(hostPort, scheme)

	firstRegisteredPattern := "/ddd"
	secondRegisteredPattern := "/eee"
	http.HandleFunc(firstRegisteredPattern, noOperationHandler)
	http.HandleFunc(secondRegisteredPattern, noOperationHandler)
	server.StartTLS("../../../resources/certs/server_ca.pem", "../../../resources/certs/server_ca.key")

	address := scheme + "://" + hostPort

	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	firstEntry := messages.NewEntry(firstRegisteredPattern, "GET")
	secondEntry := messages.NewEntry(secondRegisteredPattern, "POST")
	thirdEntry := messages.NewEntry("/ccc", "POST")

	builder := configuration.NewBuilder()
	configuration := builder.
		URLResponseTimeout(3 * time.Second).
		WorkerWaitPeriod(0).
		HTTPErrorCode(http.StatusNotFound).
		BaseURL(url).
		Build()
	sut := NewURLChecker(configuration)
	foundEntries := make(chan messages.FoundEntry, 4)
	assignChannel(sut, foundEntries)

	// when
	sut.OnEntry(firstEntry)
	sut.OnEntry(secondEntry)
	sut.OnEntry(thirdEntry)

	server.Stop()
	http.DefaultServeMux = http.NewServeMux()

	// then
	assert.Len(foundEntries, 2, "Two entries should be considered as found")
	assert.Equal(messages.NewFoundEntry(address+firstEntry.RelativeURL(), firstEntry.HTTPMethod(), http.StatusOK), <-foundEntries, "First entry should be found")
	assert.Equal(messages.NewFoundEntry(address+secondEntry.RelativeURL(), secondEntry.HTTPMethod(), http.StatusOK), <-foundEntries, "Second entry should be found")
}

func TestOnEntryURLsFound(t *testing.T) {
	assert := assert.New(t)

	// given
	scheme := "http"
	hostPort := "127.0.0.1:10608"
	server := localserver.NewLocalServer(hostPort, scheme)

	firstRegisteredPattern := "/aaa"
	secondRegisteredPattern := "/bbb"
	http.HandleFunc(firstRegisteredPattern, noOperationHandler)
	http.HandleFunc(secondRegisteredPattern, noOperationHandler)
	server.Start()

	address := scheme + "://" + hostPort

	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	firstEntry := messages.NewEntry(firstRegisteredPattern, "GET")
	secondEntry := messages.NewEntry(secondRegisteredPattern, "POST")
	thirdEntry := messages.NewEntry("/ccc", "POST")

	builder := configuration.NewBuilder()
	configuration := builder.
		URLResponseTimeout(3 * time.Second).
		WorkerWaitPeriod(0).
		HTTPErrorCode(http.StatusNotFound).
		BaseURL(url).
		Build()
	sut := NewURLChecker(configuration)
	foundEntries := make(chan messages.FoundEntry, 4)
	assignChannel(sut, foundEntries)

	// when
	sut.OnEntry(firstEntry)
	sut.OnEntry(secondEntry)
	sut.OnEntry(thirdEntry)

	server.Stop()
	http.DefaultServeMux = http.NewServeMux()

	// then
	assert.Len(foundEntries, 2, "Two entries should be considered as found")
	assert.Equal(messages.NewFoundEntry(address+firstEntry.RelativeURL(), firstEntry.HTTPMethod(), http.StatusOK), <-foundEntries, "First entry should be found")
	assert.Equal(messages.NewFoundEntry(address+secondEntry.RelativeURL(), secondEntry.HTTPMethod(), http.StatusOK), <-foundEntries, "Second entry should be found")
}

func TestOnEntryAssignedHTTPErrorCode(t *testing.T) {
	assert := assert.New(t)

	// given
	scheme := "http"
	hostPort := "127.0.0.1:10609"
	server := localserver.NewLocalServer(hostPort, scheme)

	firstRegisteredPattern := "/aaa"
	secondRegisteredPattern := "/bbb"
	http.HandleFunc(firstRegisteredPattern, forbiddenHandler)
	http.HandleFunc(secondRegisteredPattern, forbiddenHandler)
	server.Start()

	address := scheme + "://" + hostPort

	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	firstEntry := messages.NewEntry(firstRegisteredPattern, "GET")
	secondEntry := messages.NewEntry(secondRegisteredPattern, "POST")
	thirdEntry := messages.NewEntry("/ccc", "POST")

	builder := configuration.NewBuilder()
	configuration := builder.
		URLResponseTimeout(3 * time.Second).
		WorkerWaitPeriod(0).
		HTTPErrorCode(http.StatusForbidden).
		BaseURL(url).
		Build()
	sut := NewURLChecker(configuration)
	foundEntries := make(chan messages.FoundEntry, 4)
	assignChannel(sut, foundEntries)

	// when
	sut.OnEntry(firstEntry)
	sut.OnEntry(secondEntry)
	sut.OnEntry(thirdEntry)

	server.Stop()
	http.DefaultServeMux = http.NewServeMux()

	// then
	assert.Len(foundEntries, 1, "One entry should be considered as found")
	assert.Equal(messages.NewFoundEntry(address+thirdEntry.RelativeURL(), thirdEntry.HTTPMethod(), http.StatusNotFound), <-foundEntries, "Third entry should be found")
}

func TestOnEntryHTTPHeaders(t *testing.T) {
	assert := assert.New(t)

	// given
	address := "http://localhost:10610"
	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	firstHeaderName := "first"
	firstHeaderValue := "8903434"
	secondHeaderName := "second"
	secondHeaderValue := "34234324"

	headers := map[string]string{firstHeaderName: firstHeaderValue, secondHeaderName: secondHeaderValue}
	expectedURLResponseTimeout := 99 * time.Second
	builder := configuration.NewBuilder()
	configuration := builder.
		Headers(headers).
		URLResponseTimeout(expectedURLResponseTimeout).
		WorkerWaitPeriod(0).
		HTTPErrorCode(http.StatusNotFound).
		BaseURL(url).
		Build()

	// when
	sut := NewURLChecker(configuration)

	// then
	assert.NotNil(sut, "URL checker instance should be created")
	assert.Len(sut.headers, 2, "The number of HTTP headers is different than expected")
	assert.Equal([]string{firstHeaderValue}, sut.headers[firstHeaderName], "HTTP header is different")
	assert.Equal([]string{secondHeaderValue}, sut.headers[secondHeaderName], "HTTP header is different")
}

func TestCreateRequest(t *testing.T) {
	assert := assert.New(t)

	// given
	address := "http://localhost:10611"
	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	firstHeaderName := "first"
	firstHeaderValue := "8903434"
	secondHeaderName := "second"
	secondHeaderValue := "34234324"

	httpMethod := "DELETE"

	headers := map[string]string{firstHeaderName: firstHeaderValue, secondHeaderName: secondHeaderValue}
	expectedURLResponseTimeout := 99 * time.Second
	builder := configuration.NewBuilder()
	configuration := builder.
		Headers(headers).
		URLResponseTimeout(expectedURLResponseTimeout).
		WorkerWaitPeriod(0).
		HTTPErrorCode(http.StatusNotFound).
		BaseURL(url).
		Build()

	sut := NewURLChecker(configuration)

	// when
	result := sut.createRequest(httpMethod, address)

	// then
	assert.NotNil(result, "HTTP result should not be nil")
	assert.Equal(httpMethod, result.Method, "HTTP method should be same as given")
	assert.Equal([]string{firstHeaderValue}, result.Header[firstHeaderName], "HTTP header is different")
	assert.Equal([]string{secondHeaderValue}, result.Header[secondHeaderName], "HTTP header is different")
}

func assignChannel(instance *URLChecker, channel chan<- messages.FoundEntry) {
	instance.FoundEntry = channel
}

func noOperationHandler(response http.ResponseWriter, request *http.Request) {
	log.Printf("NoOperationHandler, incoming request: %v", request.RequestURI)
	response.WriteHeader(http.StatusOK)
}

func forbiddenHandler(response http.ResponseWriter, request *http.Request) {
	log.Printf("Forbidden handler, incoming request: %v", request.RequestURI)
	response.WriteHeader(http.StatusForbidden)
}
