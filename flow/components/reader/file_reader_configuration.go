package reader

import "os"

type fileReaderConfiguration interface {
	FuzzSetFile() *os.File
}
