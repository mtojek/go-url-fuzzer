package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadersAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	headerName := "a_header_name"
	headerValue := "a_header_value"
	expectedHeaders := map[string]string{headerName: headerValue}

	sut := newConfiguration()
	sut.headers = &expectedHeaders

	// when
	actualHeaders, exists := sut.Headers()

	// then
	assert.Equal(actualHeaders, expectedHeaders, "Assgined different headers.")
	assert.True(exists)
}

func TestHeadersNotAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := newConfiguration()

	// when
	actualHeaders, exists := sut.Headers()

	// then
	assert.Empty(actualHeaders, "Headers should be empty.")
	assert.False(exists)
}

func TestMethodsAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	method := "a_method"
	expectedMethods := []string{method}

	sut := newConfiguration()
	sut.methods = &expectedMethods

	// when
	actualMethods, exists := sut.Methods()

	// then
	assert.Equal(actualMethods, expectedMethods, "Assgined different methods.")
	assert.True(exists)
}

func TestMethodsNotAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := newConfiguration()

	// when
	actualMethods, exists := sut.Headers()

	// then
	assert.Empty(actualMethods, "Methods should be empty.")
	assert.False(exists)
}
