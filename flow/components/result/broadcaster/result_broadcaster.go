package broadcaster

import (
	"log"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// ResultBroadcaster receives found relative URLs and broadcasts to active output components.
type ResultBroadcaster struct {
	flow.Component

	FoundEntry <-chan messages.Entry
}

// NewResultBroadcaster creates new instance of result broadcaster.
func NewResultBroadcaster(configuration *configuration.Configuration) *ResultBroadcaster {
	return &ResultBroadcaster{}
}

// OnFoundEntry performs broadcasting.
func (u *ResultBroadcaster) OnFoundEntry(entry messages.Entry) {
	log.Println(entry)
}
