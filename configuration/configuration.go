package configuration

import (
	"net/url"
	"os"
	"time"
)

// Configuration of the application
type Configuration struct {
	headers            *map[string]string
	methods            *[]string
	outputFile         *string
	urlResponseTimeout *time.Duration
	httpErrorCode      *uint64
	workersNumber      *uint64
	workerWaitPeriod   *time.Duration
	fuzzSetFile        **os.File
	baseURL            **url.URL
}

func newConfiguration() *Configuration {
	return new(Configuration)
}

// Headers method returns HTTP headers.
func (c *Configuration) Headers() (result map[string]string, exists bool) {
	if c.headers != nil {
		result = *c.headers
		exists = true
	}
	return result, exists
}

// Methods returns unique HTTP methods.
func (c *Configuration) Methods() []string {
	return *c.methods
}

// OutputFile method returns a path of the text output file.
func (c *Configuration) OutputFile() (result string, defined bool) {
	if c.outputFile != nil && *c.outputFile != "" {
		result = *c.outputFile
		defined = true
	}
	return result, defined
}

// URLResponseTimeout method returns fuzzed URL response timeout.
func (c *Configuration) URLResponseTimeout() time.Duration {
	return *c.urlResponseTimeout
}

// HTTPErrorCode method returns a defined HTTP status error code, e.g. 404
func (c *Configuration) HTTPErrorCode() uint64 {
	return *c.httpErrorCode
}

// WorkersNumber returns a number of active fuzzing workers.
func (c *Configuration) WorkersNumber() uint64 {
	return *c.workersNumber
}

// WorkerWaitPeriod returns a period of time between two fuzzed requests per worker.
func (c *Configuration) WorkerWaitPeriod() time.Duration {
	return *c.workerWaitPeriod
}

// FuzzSetFile returns a file with fuzzed relative URLs.
func (c *Configuration) FuzzSetFile() *os.File {
	return *c.fuzzSetFile
}

// BaseURL returns a base URL of the target website.
func (c *Configuration) BaseURL() url.URL {
	return **c.baseURL
}
