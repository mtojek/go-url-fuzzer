package flow

import "github.com/trustmaster/goflow"

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
func (fuzz *Fuzz) Start() {
	fuzz.graph.InitGraphState()
}
