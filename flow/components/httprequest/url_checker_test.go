package httprequest

import (
	"testing"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/stretchr/testify/assert"
)

func TestNewURLChecker(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := configuration.NewBuilder()
	configuration := builder.Build()

	// when
	sut := NewURLChecker(configuration)

	// then
	assert.NotNil(sut, "URL checker instance should be created")
	assert.Equal(configuration, sut.configuration, "There should be set the same configuration")
}
