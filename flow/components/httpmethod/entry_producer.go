package httpmethod

import "github.com/trustmaster/goflow"

// EntryProducer processes relative URLs and
type EntryProducer struct {
	flow.Component

	RelativeURL <-chan string
	Entry       chan<- Entry
}
