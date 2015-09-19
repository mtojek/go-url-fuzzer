package httpmethod

import (
	"github.com/trustmaster/goflow"
)

// EntryProducer consumes relative URLs and produces whole entries including mentioned URLs and HTTP methods.
type EntryProducer struct {
	flow.Component

	RelativeURL <-chan string
	Entry       chan<- Entry

	methods []string
}

// NewEntryProducer creates an instance of entry producer.
func NewEntryProducer(configuration entryProducerConfiguration) *EntryProducer {
	methods := configuration.Methods()
	return &EntryProducer{methods: methods}
}

// OnRelativeURL transforms incoming relative URLs into fuzz entries by adding HTTP methods.
func (e *EntryProducer) OnRelativeURL(relativeURL string) {
	for _, httpMethod := range e.methods {
		entry := newEntry(relativeURL, httpMethod)
		e.Entry <- entry
	}
}
