package reader

import (
	"log"

	"github.com/mtojek/go-url-fuzzer/configuration"
)

// AbortableFileReader extends a standard file reader with aborting signals.
type AbortableFileReader struct {
	fileReader *fileReader
}

const fileReaderDataInChannelSize = 2 << 17

// NewAbortableFileReader creates an instance of abortable file reader.
func NewAbortableFileReader(configuration *configuration.Configuration) *AbortableFileReader {
	fileReader := newFileReader(configuration)
	return &AbortableFileReader{fileReader: fileReader}
}

// Pipe defines a pipe between a read lines and target output channel. It also supports aborting.
func (a *AbortableFileReader) Pipe(out chan<- string) {
	var in = make(chan string, fileReaderDataInChannelSize)
	var done = make(chan bool, 1)

	go a.fileReader.pipe(in, done)
	a.doPiping(in, out, done)
}

func (a *AbortableFileReader) doPiping(in <-chan string, out chan<- string, done <-chan bool) {
	abort := newAbort().signal()

	willBeClosed := false
	closing := false

	log.Println("Reading fuzz set file started.")

	for {
		select {
		case <-abort:
			log.Println("Reading fuzz set file aborted.")
			closing = true
		case line := <-in:
			out <- line
		case <-done:
			willBeClosed = true
		}

		if willBeClosed && len(in) == 0 {
			log.Println("Reading fuzz set file done.")
			closing = true
		}

		if closing {
			close(out)
			return
		}
	}
}
