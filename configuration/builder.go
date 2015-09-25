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
func (b *Builder) Headers(headers map[string]string) *Builder {
	b.configuration.headers = &headers
	return b
}

// Methods method allows to define HTTP methods.
func (b *Builder) Methods(methods []string) *Builder {
	b.configuration.methods = &methods
	return b
}

// OutputFile method allows to define a result output file.
func (b *Builder) OutputFile(outputFile string) *Builder {
	b.configuration.outputFile = &outputFile
	return b
}

// URLResponseTimeout method allows to define a fuzzed URL response timeout.
func (b *Builder) URLResponseTimeout(urlResponseTimeout time.Duration) *Builder {
	b.configuration.urlResponseTimeout = &urlResponseTimeout
	return b
}

// HTTPErrorCode method allows to define a HTTP error code.
func (b *Builder) HTTPErrorCode(httpErrorCode uint64) *Builder {
	b.configuration.httpErrorCode = &httpErrorCode
	return b
}

// WorkersNumber method allows to define number of fuzzying workers.
func (b *Builder) WorkersNumber(workersNumber uint64) *Builder {
	b.configuration.workersNumber = &workersNumber
	return b
}

// WorkerWaitPeriod method allows to define waiting period between each URL checks per worker.
func (b *Builder) WorkerWaitPeriod(workerWaitPeriod time.Duration) *Builder {
	b.configuration.workerWaitPeriod = &workerWaitPeriod
	return b
}

// FuzzSetFile method allows to define a file with fuzz entries.
func (b *Builder) FuzzSetFile(fuzzSetFile *os.File) *Builder {
	b.configuration.fuzzSetFile = &fuzzSetFile
	return b
}

// BaseURL method allows to define a base URL.
func (b *Builder) BaseURL(baseURL *url.URL) *Builder {
	b.configuration.baseURL = &baseURL
	return b
}

// Build method returns a built earlier instance.
func (b *Builder) Build() *Configuration {
	return b.configuration
}
