package httpmethod

import "github.com/trustmaster/goflow"

// EntryProducer consumes relative URLs and produces whole entries including mentioned URLs and HTTP methods.
type EntryProducer struct {
	flow.Component

	RelativeURL <-chan string
	Entry       chan<- Entry
}
