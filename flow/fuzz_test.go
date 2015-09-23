package flow

import (
	"testing"

	"log"
	"os"

	"time"

	"net/url"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/stretchr/testify/assert"
)

func TestNewFuzzMinimalConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	var workersNumber uint64 = 2
	methods := []string{"GET", "POST"}
	address := "http://localhost:10603"
	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	builder := configuration.NewBuilder()
	configuration := builder.
		WorkersNumber(workersNumber).
		WorkerWaitPeriod(0).
		Methods(methods).
		URLResponseTimeout(3 * time.Second).
		HTTPErrorCode(404).
		BaseURL(url).
		Build()

	// when
	sut := NewFuzz(configuration)

	// then
	assert.NotNil(sut, "Instance should be created")
	assert.NotNil(sut.input, "Input channel should be defined")
	assert.NotNil(sut.graph, "Flow graph should be created")
	assert.NotNil(sut.configuration, "Configuration should be set")
}

func TestStartSimpleFuzz(t *testing.T) {
	assert := assert.New(t)

	// given
	var workersNumber uint64 = 2
	methods := []string{"GET", "POST", "PUT"}
	address := "http://localhost:10604"
	url, error := url.Parse(address)
	if nil != error {
		log.Fatalf("Error occured while parsing an URL: %v, error: %v", address, error)
	}

	inputFile, error := os.OpenFile("../resources/input-data/fuzz_03.txt", os.O_RDONLY, 0666)
	if nil != error {
		log.Fatal("TestStartFuzz: ", error)
	}

	builder := configuration.NewBuilder()
	configuration := builder.
		WorkersNumber(workersNumber).
		WorkerWaitPeriod(0).
		Methods(methods).
		URLResponseTimeout(3 * time.Second).
		FuzzSetFile(inputFile).
		HTTPErrorCode(404).
		BaseURL(url).
		Build()
	sut := NewFuzz(configuration)

	// when
	sut.Start()

	// then
	assert.Len(sut.input, 0, "Input channel should be empty now")
}
