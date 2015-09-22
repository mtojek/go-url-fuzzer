package flow

import (
	"testing"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/stretchr/testify/assert"
)

func TestNewFuzzMinimalConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	var workersNumber uint64 = 2
	methods := []string{"GET", "POST"}

	builder := configuration.NewBuilder()
	configuration := builder.
		WorkersNumber(workersNumber).
		Methods(methods).
		Build()

	// when
	sut := NewFuzz(configuration)

	// then
	assert.NotNil(sut, "Instance should be created")
	assert.NotNil(sut.graph, "Flow graph should be created")
	assert.NotNil(sut.configuration, "Configuration should be set")
}
