package configuration

import (
	"net/url"
	"time"
	"os"
)

type Configuration struct {
	headers            *map[string]string
	methods            *[]string
	outputFile         *string
	reportDirectory    *string
	urlResponseTimeout *time.Duration
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

func (this *Configuration) WorkerWaitPeriod() (result time.Duration, exists bool) {
	if this.workerWaitPeriod != nil {
		result = *this.workerWaitPeriod
		exists = true
	}
	return result, exists
}

func (this *Configuration) FuzzSetFile() (result os.File, exists bool) {
	if this.fuzzSetFile != nil && (*this.fuzzSetFile) != nil {
		result = **this.fuzzSetFile
		exists = true
	}
	return result, exists
}

func (this *Configuration) BaseURL() (result url.URL, exists bool) {
	if this.baseURL != nil && (*this.baseURL) != nil {
		result = **this.baseURL
		exists = true
	}
	return result, exists
}