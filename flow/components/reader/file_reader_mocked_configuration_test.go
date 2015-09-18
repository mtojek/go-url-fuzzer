package reader

import "os"

type fileReaderMockedConfiguration struct {
	file *os.File
}

func newFileReaderMockedConfiguration(file *os.File) *fileReaderMockedConfiguration {
	return &fileReaderMockedConfiguration{file: file}
}

func (f *fileReaderMockedConfiguration) FuzzSetFile() *os.File {
	return f.file
}
