package flow

import (
	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/components/httpmethod"
	"github.com/mtojek/go-url-fuzzer/flow/components/httprequest"
	"github.com/mtojek/go-url-fuzzer/flow/components/reader"
	"github.com/mtojek/go-url-fuzzer/flow/components/result/broadcaster"
	"github.com/mtojek/go-url-fuzzer/flow/components/result/filewriter"
	"github.com/mtojek/go-url-fuzzer/flow/components/result/printer"
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

	resultBroadcaster := broadcaster.NewResultBroadcaster(configuration)
	resultBroadcaster.Component.Mode = flow.ComponentModePool
	resultBroadcaster.Component.PoolSize = 1

	printer := printer.NewPrinter()
	printer.Component.Mode = flow.ComponentModePool
	printer.Component.PoolSize = 1

	fileWriter := filewriter.NewFileWriter(configuration)
	printer.Component.Mode = flow.ComponentModePool
	printer.Component.PoolSize = 1

	graph.Add(entryProducer, "entryProducer")
	graph.Add(urlChecker, "urlChecker")
	graph.Add(resultBroadcaster, "resultBroadcaster")
	graph.Add(printer, "printer")
	graph.Add(fileWriter, "fileWriter")
	graph.Connect("entryProducer", "Entry", "urlChecker", "Entry")
	graph.Connect("urlChecker", "FoundEntry", "resultBroadcaster", "FoundEntry")
	graph.Connect("resultBroadcaster", "Printer", "printer", "FoundEntry")
	graph.Connect("resultBroadcaster", "FileWriter", "fileWriter", "FoundEntry")
	graph.MapInPort("In", "entryProducer", "RelativeURL")

	var input = make(chan string, fuzzNetworkInputSize)
	graph.SetInPort("In", input)

	return &Fuzz{graph: graph, input: input, configuration: configuration}
}

// Start method runs the flow components network and initiates the input producer.
func (f *Fuzz) Start() {
	flow.RunNet(f.graph)

	abortableFileReader := reader.NewAbortableFileReader(f.configuration)
	isDone := abortableFileReader.Pipe(f.input)
	if isDone {
		<-f.graph.Wait()
	} else {
		f.graph.Stop()
	}
}
