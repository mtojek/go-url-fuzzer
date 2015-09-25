package printer

import (
	"log"

	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// Printer receives found relative URLs and prints them to default logging output.
type Printer struct {
	flow.Component

	FoundEntry <-chan messages.FoundEntry
}

// NewPrinter creates new instance of a printer.
func NewPrinter() *Printer {
	return new(Printer)
}

// OnFoundEntry performs printing.
func (p *Printer) OnFoundEntry(foundEntry messages.FoundEntry) {
	log.Println(foundEntry.String())
}
