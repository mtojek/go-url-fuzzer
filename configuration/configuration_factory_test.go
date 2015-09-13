package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFlagsBoundConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := NewConfigurationFactory()

	// when
	result := sut.createFlagsBoundConfiguration()

	// then
	assert.NotNil(result)
	assert.NotNil(result.headers)
	assert.NotNil(result.methods)
	assert.NotNil(result.outputFile)
	assert.NotNil(result.reportDirectory)
	assert.NotNil(result.urlResponseTimeout)
	assert.NotNil(result.workersNumber)
	assert.NotNil(result.workerWaitPeriod)
	assert.NotNil(result.fuzzSetFile)
	assert.NotNil(result.baseURL)
}
