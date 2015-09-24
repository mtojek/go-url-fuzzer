package messages

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
