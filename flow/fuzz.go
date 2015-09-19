package flow

import (
	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/components/httpmethod"
	"github.com/mtojek/go-url-fuzzer/flow/components/httprequest"
	"github.com/mtojek/go-url-fuzzer/flow/components/reader"
	"github.com/trustmaster/goflow"
)

const fuzzNetworkInputSize = 256

// Fuzz wraps the flow graph, including network input channel and configuration.
type Fuzz struct {
	graph *flow.Graph
	input chan string

	configuration *configuration.Configuration
}

// NewFuzz creates a fuzz URL flow with defined components and links.
func NewFuzz(configuration *configuration.Configuration) *Fuzz {
	graph := new(flow.Graph)
	graph.InitGraphState()

	entryProducer := httpmethod.NewEntryProducer(configuration)
	entryProducer.Component.Mode = flow.ComponentModePool
	entryProducer.Component.PoolSize = 8

	urlChecker := httprequest.NewURLChecker(configuration)
	urlChecker.Component.Mode = flow.ComponentModePool
	urlChecker.Component.PoolSize = uint8(configuration.WorkersNumber())

	graph.Add(entryProducer, "entryProducer")
	graph.Add(urlChecker, "urlChecker")
	graph.Connect("entryProducer", "Entry", "urlChecker", "Entry")
	graph.MapInPort("In", "entryProducer", "RelativeURL")

	var input = make(chan string, fuzzNetworkInputSize)
	graph.SetInPort("In", input)

	return &Fuzz{graph: graph, input: input, configuration: configuration}
}

// Start method runs the flow components network and initiates the input producer.
func (f *Fuzz) Start() {
	flow.RunNet(f.graph)

	abortableFileReader := reader.NewAbortableFileReader(f.configuration)
	abortableFileReader.Pipe(f.input)

	<-f.graph.Wait()
}
