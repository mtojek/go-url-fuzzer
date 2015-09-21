package httpmethod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntryProducer(t *testing.T) {
	assert := assert.New(t)

	// given
	methods := []string{"method_a", "method_b"}
	configuration := newEntryProducerMockedConfiguration(methods)

	// when
	sut := NewEntryProducer(configuration)

	// then
	assert.Equal(sut.methods, methods, "HTTP Methods should be the same as defined by used")
}

func TestOnRelativeURLThreeMethods(t *testing.T) {
	assert := assert.New(t)

	// given
	anURL := "an_url"
	entries := make(chan Entry, 3)

	methods := []string{"method_a", "method_b", "method_c"}
	configuration := newEntryProducerMockedConfiguration(methods)
	sut := NewEntryProducer(configuration)
	sut.Entry = entries

	// when
	sut.OnRelativeURL(anURL)

	// then
	assert.Equal(<-entries, newEntry(anURL, "method_a"), "Entries are not equal")
	assert.Equal(<-entries, newEntry(anURL, "method_b"), "Entries are not equal")
	assert.Equal(<-entries, newEntry(anURL, "method_c"), "Entries are not equal")
	assert.Len(entries, 0, "Entries channel should be empty now")
}

func TestOnRelativeURLOneMethod(t *testing.T) {
	assert := assert.New(t)

	// given
	anURL := "an_url"
	entries := make(chan Entry, 3)

	methods := []string{"method_a"}
	configuration := newEntryProducerMockedConfiguration(methods)
	sut := NewEntryProducer(configuration)
	sut.Entry = entries

	// when
	sut.OnRelativeURL(anURL)

	// then
	assert.Equal(<-entries, newEntry(anURL, "method_a"), "Entries are not equal")
	assert.Len(entries, 0, "Entries channel should be empty now")
}
