package httpmethod

import (
	"testing"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/stretchr/testify/assert"
)

func TestNewEntryProducer(t *testing.T) {
	assert := assert.New(t)

	// given
	methods := []string{"method_a", "method_b"}
	builder := configuration.NewBuilder()
	configuration := builder.Methods(methods).Build()

	// when
	sut := NewEntryProducer(configuration)

	// then
	assert.Equal(sut.methods, methods, "HTTP Methods should be the same as defined by used")
}

func TestOnRelativeURLThreeMethods(t *testing.T) {
	assert := assert.New(t)

	// given
	anURL := "an_url"
	entries := make(chan messages.Entry, 3)

	methods := []string{"method_a", "method_b", "method_c"}
	builder := configuration.NewBuilder()
	configuration := builder.Methods(methods).Build()
	sut := NewEntryProducer(configuration)
	sut.Entry = entries

	// when
	sut.OnRelativeURL(anURL)

	// then
	assert.Equal(<-entries, messages.NewEntry(anURL, "method_a"), "Entries are not equal")
	assert.Equal(<-entries, messages.NewEntry(anURL, "method_b"), "Entries are not equal")
	assert.Equal(<-entries, messages.NewEntry(anURL, "method_c"), "Entries are not equal")
	assert.Len(entries, 0, "Entries channel should be empty now")
}

func TestOnRelativeURLOneMethod(t *testing.T) {
	assert := assert.New(t)

	// given
	anURL := "an_url"
	entries := make(chan messages.Entry, 3)

	methods := []string{"method_a"}
	builder := configuration.NewBuilder()
	configuration := builder.Methods(methods).Build()
	sut := NewEntryProducer(configuration)
	sut.Entry = entries

	// when
	sut.OnRelativeURL(anURL)

	// then
	assert.Equal(<-entries, messages.NewEntry(anURL, "method_a"), "Entries are not equal")
	assert.Len(entries, 0, "Entries channel should be empty now")
}
