package messages

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestFoundEntryProperties(t *testing.T) {
	assert := assert.New(t)

	// given
	absoluteURL := "an_url"
	httpMethod := "a_method"
	status := http.StatusBadGateway

	// when
	sut := NewFoundEntry(absoluteURL, httpMethod, status)

	// then
	assert.Equal(absoluteURL, sut.absoluteURL, "Absolute URL must be same as given")
	assert.Equal(httpMethod, sut.httpMethod, "HTTP method must be same as given")
	assert.Equal(status, sut.status, "HTTP status code must be same as given")
}

func TestFoundEntryString(t *testing.T) {
	assert := assert.New(t)

	// given
	absoluteURL := "an_url"
	httpMethod := "a_method"
	status := http.StatusBadGateway
	sut := NewFoundEntry(absoluteURL, httpMethod, status)

	// when
	result := sut.String()

	// then
	assert.Equal(httpMethod+" "+absoluteURL+" "+fmt.Sprintf("%d", http.StatusBadGateway), result, "String representation is different than expected")
}
