package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFlagsBoundConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := NewFactory()

	// when
	result := sut.createFlagsBoundConfiguration()

	// then
	assert.NotNil(result, "Configuration should not be nil")
	assert.NotNil(result.headers, "Headers should not be nil")
	assert.NotNil(result.methods, "Methods should not be nil")
	assert.NotNil(result.outputFile, "Output file should not be nil")
	assert.NotNil(result.urlResponseTimeout, "Response timeout should not be nil")
	assert.NotNil(result.workersNumber, "Workers number should not be nil")
	assert.NotNil(result.httpErrorCode, "HTTP error code should not be nil")
	assert.NotNil(result.workerWaitPeriod, "Worker wait period should not be nil")
	assert.NotNil(result.fuzzSetFile, "Fuzz set file should not be nil")
	assert.NotNil(result.baseURL, "Base URL should not be nil")
}
