package reader

import (
	"log"

	"github.com/mtojek/go-url-fuzzer/configuration"
)

// AbortableFileReader extends a standard file reader with aborting signals.
type AbortableFileReader struct {
	fileReader *fileReader
}

const fileReaderDataOutChannelSize = 2 >> 18

// NewAbortableFileReader creates an instance of abortable file reader.
func NewAbortableFileReader(configuration *configuration.Configuration) *AbortableFileReader {
	fileReader := newFileReader(configuration)
	return &AbortableFileReader{fileReader: fileReader}
}

// Pipe defines a pipe between a read lines and target output channel. It also supports aborting.
func (a *AbortableFileReader) Pipe(out chan string) {
	abort := newAbort().signal()

	var fileReaderDataOut = make(chan string, fileReaderDataOutChannelSize)
	var fileReaderDone = make(chan bool, 1)
	go a.fileReader.pipe(fileReaderDataOut, fileReaderDone)

	for {

		select {
		case <-abort:
			log.Println("STOP!")
			close(fileReaderDataOut)

			return
		case line := <-fileReaderDataOut:
			log.Println(line)
			//out <- line // TODO odkomentowac
		case <-fileReaderDone:
			log.Println("DONE!")
			close(fileReaderDataOut)

			return
		}
	}
}
