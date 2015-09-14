package main

import (
	"fmt"

	"github.com/mtojek/go-url-fuzzer/configuration"
)

type urlFuzzer struct{}

func newURLFuzzer() *urlFuzzer {
	return new(urlFuzzer)
}

func (urlFuzzer *urlFuzzer) run() {
	configuration := urlFuzzer.readConfiguration()
	fmt.Println(configuration)
}

func (urlFuzzer *urlFuzzer) readConfiguration() *configuration.Configuration {
	configurationFactory := configuration.NewFactory()
	return configurationFactory.FromCommandLine()
}
