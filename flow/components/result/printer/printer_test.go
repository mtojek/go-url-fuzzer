package printer

import (
	"testing"

	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/stretchr/testify/assert"
)

func TestNewPrinter(t *testing.T) {
	assert := assert.New(t)

	// when
	sut := NewPrinter()

	// then
	assert.NotNil(sut, "New instance is not nil")
}

func TestOnEntry(t *testing.T) {
	// given
	foundEntry := messages.NewFoundEntry("1", "2", 3)
	sut := NewPrinter()

	// when
	sut.OnFoundEntry(foundEntry)
}
