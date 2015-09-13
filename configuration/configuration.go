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

func (this *Configuration) Methods() (result []string, exists bool) {
	if this.methods != nil {
		result = *this.methods
		exists = true
	}
	return result, exists
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

func (this *Configuration) UrlResponseTimeout() (result time.Duration, exists bool) {
	if this.urlResponseTimeout != nil {
		result = *this.urlResponseTimeout
		exists = true
	}
	return result, exists
}

func (this *Configuration) WorkersNumber() (result uint64, exists bool) {
	if this.workersNumber != nil {
		result = *this.workersNumber
		exists = true
	}
	return result, exists
}

func (this *Configuration) WorkerWaitPeriod() (result time.Duration, exists bool) {
	if this.workerWaitPeriod != nil {
		result = *this.workerWaitPeriod
		exists = true
	}
	return result, exists
}

func (this *Configuration) FuzzSetFile() os.File {
	return **this.fuzzSetFile
}

func (this *Configuration) BaseURL() url.URL {
	return **this.baseURL
}
