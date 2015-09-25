package flow

import (
	"testing"

	"log"
	"os"

	"time"

	"net/url"

	"net/http"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/localserver"
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

func TestStartSimpleFuzzNoServerRunning(t *testing.T) {
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

func TestStartSimpleFuzzWithServerRunning(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "localhost:10612"
	scheme := "http"
	address := scheme + "://" + hostPort

	server := localserver.NewLocalServer(hostPort, scheme)
	firstHandler := newVisitHandler("/smietnik", "GET")
	secondHandler := newVisitHandler("/spotkania", "POST")
	thirdHandler := newVisitHandler("/aukcje", "PUT")

	http.HandleFunc(firstHandler.endpoint, firstHandler.handle)
	http.HandleFunc(secondHandler.endpoint, secondHandler.handle)
	http.HandleFunc(thirdHandler.endpoint, thirdHandler.handle)

	server.Start()

	var workersNumber uint64 = 4
	methods := []string{"GET", "POST", "PUT"}

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

	server.Stop()
	http.DefaultServeMux = http.NewServeMux()

	// then
	assert.Len(sut.input, 0, "Input channel should be empty now")
	assert.True(firstHandler.visitted, "Handler "+firstHandler.endpoint+" should be found")
	assert.True(secondHandler.visitted, "Handler "+secondHandler.endpoint+" should be found")
	assert.True(thirdHandler.visitted, "Handler "+thirdHandler.endpoint+" should be found")
}
