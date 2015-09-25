package filewriter

import (
	"log"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// FileWriter receives found relative URLs and writes to the output file.
type FileWriter struct {
	flow.Component

	FoundEntry <-chan messages.FoundEntry
}

// NewFileWriter creates new instance of a file writer.
func NewFileWriter(configuration *configuration.Configuration) *FileWriter {
	return new(FileWriter)
}

// OnFoundEntry performs writing to file.
func (f *FileWriter) OnFoundEntry(foundEntry messages.FoundEntry) {
	log.Println(foundEntry.String())
}
