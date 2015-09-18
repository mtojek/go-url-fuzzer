package flow

import (
	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/components/reader"
	"github.com/mtojek/goflow"
)

// Fuzz wraps flow.Graph abstraction.
type Fuzz struct {
	graph *flow.Graph
}

// NewFuzz creates a fuzz URL flow.
func NewFuzz() *Fuzz {
	graph := new(flow.Graph)
	return &Fuzz{graph: graph}
}

// Start methods starts the flow.
func (f *Fuzz) Start(configuration *configuration.Configuration) {
	f.graph.InitGraphState()

	var ch = make(chan string, (2<<17)+10) // TODO remove

	abortableFileReader := reader.NewAbortableFileReader(configuration)
	abortableFileReader.Pipe(ch)
}
