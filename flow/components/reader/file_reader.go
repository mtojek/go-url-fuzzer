package components

import (
	"os"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/components/abort"
)

// FileReader is a reader of files.
type FileReader struct {
	inputFile os.File
}

// NewFileReader creates new instance of a file reader.
func NewFileReader(configuration *configuration.Configuration) *FileReader {
	inputFile := configuration.FuzzSetFile()
	return &FileReader{inputFile: inputFile}
}

// Pipe read file contents to channel. Piping can be aborted with Ctrl-C.
func (f *FileReader) Pipe(chan string) {
	notifier := abort.NewNotifier()
	notifier.Notify()
}
