package messages

// Entry contains a relative URL with corresponding HTTP method.
type Entry struct {
	relativeURL string
	httpMethod  string
}

// NewEntry creates a new instance of entry.
func NewEntry(relativeURL, httpMethod string) Entry {
	return Entry{relativeURL: relativeURL, httpMethod: httpMethod}
}

// RelativeURL returns a relative url.
func (e *Entry) RelativeURL() string {
	return e.relativeURL
}

// HTTPMethod returns a HTTP method used in fuzzing.
func (e *Entry) HTTPMethod() string {
	return e.httpMethod
}
