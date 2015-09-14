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
	reportDirectory    *string
	urlResponseTimeout *time.Duration
	workersNumber      *uint64
	workerWaitPeriod   *time.Duration
	fuzzSetFile        **os.File
	baseURL            **url.URL
}

func newConfiguration() *Configuration {
	return new(Configuration)
}

// Headers method returns HTTP headers.
func (configuration *Configuration) Headers() (result map[string]string, exists bool) {
	if configuration.headers != nil {
		result = *configuration.headers
		exists = true
	}
	return result, exists
}

// Methods returns unique HTTP methods.
func (configuration *Configuration) Methods() []string {
	return *configuration.methods
}

// OutputFile method returns a path of the text output file.
func (configuration *Configuration) OutputFile() (result string, exists bool) {
	if configuration.outputFile != nil {
		result = *configuration.outputFile
		exists = true
	}
	return result, exists
}

// ReportDirectory method returns a target directory of HTML report files.
func (configuration *Configuration) ReportDirectory() (result string, exists bool) {
	if configuration.reportDirectory != nil {
		result = *configuration.reportDirectory
		exists = true
	}
	return result, exists
}

// URLResponseTimeout method returns fuzzed URL response timeout.
func (configuration *Configuration) URLResponseTimeout() time.Duration {
	return *configuration.urlResponseTimeout
}

// WorkersNumber returns a number of active fuzzing workers.
func (configuration *Configuration) WorkersNumber() uint64 {
	return *configuration.workersNumber
}

// WorkerWaitPeriod returns a period of time between two fuzzed requests per worker.
func (configuration *Configuration) WorkerWaitPeriod() time.Duration {
	return *configuration.workerWaitPeriod
}

// FuzzSetFile returns a file with fuzzed relative URLs.
func (configuration *Configuration) FuzzSetFile() os.File {
	return **configuration.fuzzSetFile
}

// BaseURL returns a base URL of the target website.
func (configuration *Configuration) BaseURL() url.URL {
	return **configuration.baseURL
}
