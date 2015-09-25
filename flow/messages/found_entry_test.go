package messages

import (
	"net/http"
	"testing"

	"log"
	"net/url"

	"fmt"

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

func TestFoundEntryString(t *testing.T) {
	assert := assert.New(t)

	// given
	baseURLAddress := "https://localhost"
	baseURL, error := url.Parse(baseURLAddress)
	if nil != error {
		log.Fatalf("Error occured while parsing address \"%v\": %v", baseURLAddress, error)
	}

	relativeURL := "an_url"
	httpMethod := "a_method"
	entry := NewEntry(relativeURL, httpMethod)
	status := http.StatusBadGateway

	sut := NewFoundEntry(entry, status)

	// when
	result := sut.String(*baseURL)

	// then
	assert.Equal(httpMethod+" "+baseURLAddress+"/"+relativeURL+" "+fmt.Sprintf("%d", http.StatusBadGateway), result, "String representation is different than expected")
}
