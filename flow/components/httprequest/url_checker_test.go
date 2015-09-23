package httprequest

import (
	"testing"

	"time"

	"log"
	"net/url"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
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
		HTTPErrorCode(404).
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
		HTTPErrorCode(404).
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

func assignChannel(instance *URLChecker, channel chan<- messages.Entry) {
	instance.FoundEntry = channel
}
