package broadcaster

import (
	"testing"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/stretchr/testify/assert"
)

func TestNewResultBroadcaster(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := configuration.NewBuilder()
	configuration := builder.OutputFile("a_file").Build()

	// when
	sut := NewResultBroadcaster(configuration)

	// then
	assert.NotNil(sut, "New instance is not nil")
	assert.True(sut.isOutputFileDefined, "Output file has been defined")
}

func TestOnEntryOutputFileUndefined(t *testing.T) {
	assert := assert.New(t)

	// given
	foundEntry := messages.NewFoundEntry("1", "2", 3)

	builder := configuration.NewBuilder()
	configuration := builder.Build()
	sut := NewResultBroadcaster(configuration)

	printer := make(chan messages.FoundEntry, 1)
	fileWriter := make(chan messages.FoundEntry, 1)
	assignChannels(sut, printer, fileWriter)

	// when
	sut.OnFoundEntry(foundEntry)

	// then
	assert.False(sut.isOutputFileDefined, "Output file has been defined")
	assert.Len(printer, 1, "One entry directed to printer")
	assert.Len(fileWriter, 0, "No entries directed to file writer")
}

func TestOnEntryOutputFileDefined(t *testing.T) {
	assert := assert.New(t)

	// given
	foundEntry := messages.NewFoundEntry("1", "2", 3)

	builder := configuration.NewBuilder()
	configuration := builder.
		OutputFile("a_file").
		Build()
	sut := NewResultBroadcaster(configuration)

	printer := make(chan messages.FoundEntry, 1)
	fileWriter := make(chan messages.FoundEntry, 1)
	assignChannels(sut, printer, fileWriter)

	// when
	sut.OnFoundEntry(foundEntry)

	// then
	assert.True(sut.isOutputFileDefined, "Output file has been defined")
	assert.Len(printer, 1, "One entry directed to printer")
	assert.Len(fileWriter, 1, "One entry directed to file writer")
}

func assignChannels(instance *ResultBroadcaster, printer chan<- messages.FoundEntry, fileWriter chan<- messages.FoundEntry) {
	instance.Printer = printer
	instance.FileWriter = fileWriter
}
