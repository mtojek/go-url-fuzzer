package httprequest

import (
	"testing"

	"time"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/stretchr/testify/assert"
)

func TestNewURLChecker(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := configuration.NewBuilder()
	configuration := builder.
		URLResponseTimeout(3 * time.Second).
		WorkerWaitPeriod(0).
		HTTPErrorCode(404).
		Build()

	// when
	sut := NewURLChecker(configuration)

	// then
	assert.NotNil(sut, "URL checker instance should be created")
	assert.NotNil(sut.client, "HTTP client should be set")
}
