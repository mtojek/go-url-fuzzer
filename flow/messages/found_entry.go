package messages

import (
	"bytes"
	"fmt"
)

// FoundEntry represents a found relative URL with the particular HTTP method and returned HTTP status code.
type FoundEntry struct {
	absoluteURL string
	httpMethod  string
	status      int
}

// NewFoundEntry creates a new instance of found entry.
func NewFoundEntry(absoluteURL string, httpMethod string, status int) FoundEntry {
	return FoundEntry{
		absoluteURL: absoluteURL,
		httpMethod:  httpMethod,
		status:      status,
	}
}

// AbsoluteURL method returns a found absolute URL.
func (f *FoundEntry) AbsoluteURL() string {
	return f.absoluteURL
}

// HTTPMethod returns a HTTP method.
func (f *FoundEntry) HTTPMethod() string {
	return f.httpMethod
}

// Status method returns a HTTP response code of visited found URL.
func (f *FoundEntry) Status() int {
	return f.status
}

// String method returns a string representation of instance.
func (f *FoundEntry) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(f.httpMethod)
	buffer.WriteByte(' ')
	buffer.WriteString(f.absoluteURL)
	buffer.WriteByte(' ')
	buffer.WriteString(fmt.Sprintf("%d", f.status))
	return buffer.String()
}
