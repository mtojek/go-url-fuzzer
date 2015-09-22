package configuration

import (
	"net/url"
	"os"
	"time"
)

// Builder is responsible for building a new instance of Configuration.
type Builder struct {
	configuration *Configuration
}

// NewBuilder creates a new builder.
func NewBuilder() *Builder {
	configuration := newConfiguration()
	return &Builder{configuration: configuration}
}

// Headers method allows to define HTTP headers.
func (c *Builder) Headers(headers map[string]string) *Builder {
	c.configuration.headers = &headers
	return c
}

// Methods method allows to define HTTP methods.
func (c *Builder) Methods(methods []string) *Builder {
	c.configuration.methods = &methods
	return c
}

// OutputFile method allows to define a result output file.
func (c *Builder) OutputFile(outputFile string) *Builder {
	c.configuration.outputFile = &outputFile
	return c
}

// ReportDirectory method allows to define a report directory.
func (c *Builder) ReportDirectory(reportDirectory string) *Builder {
	c.configuration.reportDirectory = &reportDirectory
	return c
}

// URLResponseTimeout method allows to define a fuzzed URL response timeout.
func (c *Builder) URLResponseTimeout(urlResponseTimeout time.Duration) *Builder {
	c.configuration.urlResponseTimeout = &urlResponseTimeout
	return c
}

// HTTPErrorCode method allows to define a HTTP error code.
func (c *Builder) HTTPErrorCode(httpErrorCode uint64) *Builder {
	c.configuration.httpErrorCode = &httpErrorCode
	return c
}

// WorkersNumber method allows to define number of fuzzying workers.
func (c *Builder) WorkersNumber(workersNumber uint64) *Builder {
	c.configuration.workersNumber = &workersNumber
	return c
}

// WorkerWaitPeriod method allows to define waiting period between each URL checks per worker.
func (c *Builder) WorkerWaitPeriod(workerWaitPeriod time.Duration) *Builder {
	c.configuration.workerWaitPeriod = &workerWaitPeriod
	return c
}

// FuzzSetFile method allows to define a file with fuzz entries.
func (c *Builder) FuzzSetFile(fuzzSetFile *os.File) *Builder {
	c.configuration.fuzzSetFile = &fuzzSetFile
	return c
}

// BaseURL method allows to define a base URL.
func (c *Builder) BaseURL(baseURL *url.URL) *Builder {
	c.configuration.baseURL = &baseURL
	return c
}

// Build method returns a built earlier instance.
func (c *Builder) Build() *Configuration {
	return c.configuration
}
