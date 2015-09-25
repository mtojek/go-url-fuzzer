package filewriter

import (
	"log"

	"fmt"
	"os"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// FileWriter receives found relative URLs and writes to the output file.
type FileWriter struct {
	flow.Component

	FoundEntry <-chan messages.FoundEntry
	outputFile *os.File
}

// NewFileWriter creates new instance of a file writer.
func NewFileWriter(configuration *configuration.Configuration) *FileWriter {
	var outputFile *os.File
	var error error

	path, defined := configuration.OutputFile()
	if defined {
		log.Printf("Opening file \"%v\" for writing.", path)

		outputFile, error = os.Create(path)
		if nil != error {
			log.Fatalf("Error occured while opening file \"%v\", error: %v", path, error)
		}
	}

	return &FileWriter{outputFile: outputFile}
}

// OnFoundEntry performs writing to file.
func (f *FileWriter) OnFoundEntry(foundEntry messages.FoundEntry) {
	f.outputFile.WriteString(fmt.Sprintln(foundEntry.String()))
}

// Finish method is responsible for closing the output file.
func (f *FileWriter) Finish() {
	if nil != f.outputFile {
		log.Printf("Closing results file.")
		defer f.outputFile.Close()
	}
}
