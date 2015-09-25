package messages

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
)

// FoundEntry represents a found relative URL with the particular HTTP method and returned HTTP status code.
type FoundEntry struct {
	entry  Entry
	status int
}

// NewFoundEntry creates a new instance of entry.
func NewFoundEntry(entry Entry, status int) FoundEntry {
	return FoundEntry{entry: entry, status: status}
}

// Entry method returns found entry
func (f *FoundEntry) Entry() Entry {
	return f.entry
}

// Status method returns HTTP status code.
func (f *FoundEntry) Status() int {
	return f.status
}

func (f *FoundEntry) String(baseURL url.URL) string {
	var buffer bytes.Buffer

	absoluteURL, error := baseURL.Parse(f.entry.relativeURL)
	if nil != error {
		log.Fatalf("Error occured while preparing string representation of found entry, base URL: %v, relative URL: %v, error: %v", baseURL.String(), f.entry.relativeURL, error)
	}

	buffer.WriteString(f.entry.httpMethod)
	buffer.WriteByte(' ')
	buffer.WriteString(absoluteURL.String())
	buffer.WriteByte(' ')
	buffer.WriteString(fmt.Sprintf("%d", f.status))

	return buffer.String()
}
