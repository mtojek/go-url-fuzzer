package configuration

import (
	"net/url"
	"os"
	"time"
)

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

func (this *Configuration) Headers() (result map[string]string, exists bool) {
	if this.headers != nil {
		result = *this.headers
		exists = true
	}
	return result, exists
}

func (this *Configuration) Methods() []string {
	return *this.methods
}

func (this *Configuration) OutputFile() (result string, exists bool) {
	if this.outputFile != nil {
		result = *this.outputFile
		exists = true
	}
	return result, exists
}

func (this *Configuration) ReportDirectory() (result string, exists bool) {
	if this.reportDirectory != nil {
		result = *this.reportDirectory
		exists = true
	}
	return result, exists
}

func (this *Configuration) UrlResponseTimeout() time.Duration {
	return *this.urlResponseTimeout
}

func (this *Configuration) WorkersNumber() uint64 {
	return *this.workersNumber
}

func (this *Configuration) WorkerWaitPeriod() time.Duration {
	return *this.workerWaitPeriod
}

func (this *Configuration) FuzzSetFile() os.File {
	return **this.fuzzSetFile
}

func (this *Configuration) BaseURL() url.URL {
	return **this.baseURL
}
