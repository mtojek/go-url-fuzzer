package configuration

import (
	"net/url"
	"os"
	"time"
)

type ConfigurationBuilder struct {
	configuration *Configuration
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	configuration := newConfiguration()
	return &ConfigurationBuilder{configuration: configuration}
}

func (c *ConfigurationBuilder) Headers(headers map[string]string) *ConfigurationBuilder {
	c.configuration.headers = &headers
	return c
}

func (c *ConfigurationBuilder) Methods(methods []string) *ConfigurationBuilder {
	c.configuration.methods = &methods
	return c
}

func (c *ConfigurationBuilder) OutputFile(outputFile string) *ConfigurationBuilder {
	c.configuration.outputFile = &outputFile
	return c
}

func (c *ConfigurationBuilder) ReportDirectory(reportDirectory string) *ConfigurationBuilder {
	c.configuration.reportDirectory = &reportDirectory
	return c
}

func (c *ConfigurationBuilder) URLResponseTimeout(urlResponseTimeout time.Duration) *ConfigurationBuilder {
	c.configuration.urlResponseTimeout = &urlResponseTimeout
	return c
}

func (c *ConfigurationBuilder) HTTPErrorCode(httpErrorCode uint64) *ConfigurationBuilder {
	c.configuration.httpErrorCode = &httpErrorCode
	return c
}

func (c *ConfigurationBuilder) WorkersNumber(workersNumber uint64) *ConfigurationBuilder {
	c.configuration.workersNumber = &workersNumber
	return c
}

func (c *ConfigurationBuilder) WorkerWaitPeriod(workerWaitPeriod time.Duration) *ConfigurationBuilder {
	c.configuration.workerWaitPeriod = &workerWaitPeriod
	return c
}

func (c *ConfigurationBuilder) FuzzSetFile(fuzzSetFile *os.File) *ConfigurationBuilder {
	c.configuration.fuzzSetFile = &fuzzSetFile
	return c
}

func (c *ConfigurationBuilder) BaseURL(baseURL *url.URL) *ConfigurationBuilder {
	c.configuration.baseURL = &baseURL
	return c
}

func (c *ConfigurationBuilder) Build() *Configuration {
	return c.configuration
}
