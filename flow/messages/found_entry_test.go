package messages

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoundEntryProperties(t *testing.T) {
	assert := assert.New(t)

	// given
	relativeURL := "an_url"
	httpMethod := "a_method"
	entry := NewEntry(relativeURL, httpMethod)
	status := http.StatusBadGateway

	// when
	sut := NewFoundEntry(entry, status)

	// then
	assert.Equal(entry, sut.entry, "Entry is different than expected")
	assert.Equal(status, sut.status, "HTTP status is different than expected")
}
