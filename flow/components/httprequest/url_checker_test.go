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

func TestNewURLChecker(t *testing.T) {
	assert := assert.New(t)

	// given
	address := "http://localhost:10605"
	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	builder := configuration.NewBuilder()
	configuration := builder.
		URLResponseTimeout(3 * time.Second).
		WorkerWaitPeriod(0).
		HTTPErrorCode(http.StatusNotFound).
		BaseURL(url).
		Build()

	// when
	sut := NewURLChecker(configuration)

	// then
	assert.NotNil(sut, "URL checker instance should be created")
	assert.NotNil(sut.client, "HTTP client should be set")
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
		HTTPErrorCode(http.StatusNotFound).
		BaseURL(url).
		Build()
	sut := NewURLChecker(configuration)
	foundEntries := make(chan messages.Entry, 4)
	assignChannel(sut, foundEntries)

	// when
	sut.OnEntry(firstEntry)
	sut.OnEntry(secondEntry)

	// then
	assert.Len(foundEntries, 0, "No entries should be considered as found")
}

func TestOnEntryURLsFound(t *testing.T) {
	assert := assert.New(t)

	// given
	scheme := "http"
	hostPort := "127.0.0.1:10607"
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
	foundEntries := make(chan messages.Entry, 4)
	assignChannel(sut, foundEntries)

	// when
	sut.OnEntry(firstEntry)
	sut.OnEntry(secondEntry)
	sut.OnEntry(thirdEntry)

	server.Stop()
	http.DefaultServeMux = http.NewServeMux()

	// then
	assert.Len(foundEntries, 2, "Two entries should be considered as found")
	assert.Equal(<-foundEntries, firstEntry, "First entry should be found")
	assert.Equal(<-foundEntries, secondEntry, "Second entry should be found")
}

func TestOnEntryAssignedHTTPErrorCode(t *testing.T) {
	assert := assert.New(t)

	// given
	scheme := "http"
	hostPort := "127.0.0.1:10607"
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
	foundEntries := make(chan messages.Entry, 4)
	assignChannel(sut, foundEntries)

	// when
	sut.OnEntry(firstEntry)
	sut.OnEntry(secondEntry)
	sut.OnEntry(thirdEntry)

	server.Stop()
	http.DefaultServeMux = http.NewServeMux()

	// then
	assert.Len(foundEntries, 1, "One entry should be considered as found")
	assert.Equal(<-foundEntries, thirdEntry, "Third entry should be found")
}

func assignChannel(instance *URLChecker, channel chan<- messages.Entry) {
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
