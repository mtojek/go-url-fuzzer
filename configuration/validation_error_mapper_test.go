package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapUnknownErrorTag(t *testing.T) {
	assert := assert.New(t)

	// given
	const nonExistingTag = 666999666
	sut := newValidationErrorMapper()

	// when
	result := sut.mapErrorTag(nonExistingTag)

	// then
	assert.NotNil(result, "There should be error returned")
	assert.Equal(result.Error(), unknownErrorMessage)
}
