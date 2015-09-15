package httpmethod

// Entry contains a relative URL to fuzz with corresponding HTTP method.
type Entry struct {
	relativeURL string
	httpMethod  string
}

func newEntry(relativeURL, httpMethod string) Entry {
	return Entry{relativeURL: relativeURL, httpMethod: httpMethod}
}

// RelativeURL returns a relative url to be fuzzed.
func (e *Entry) RelativeURL() string {
	return e.relativeURL
}

// HTTPMethod returns a HTTP method which will be used in fuzzing.
func (e *Entry) HTTPMethod() string {
	return e.httpMethod
}
