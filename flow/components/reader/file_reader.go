package components

// FileReader is a reader of files.
type FileReader struct{}

// NewFileReader creates new instance of a file reader.
func NewFileReader() *FileReader {
	return new(FileReader)
}
