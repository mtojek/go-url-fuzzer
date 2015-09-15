package main

import (
	"fmt"

	"github.com/mtojek/go-url-fuzzer/configuration"
)

type urlFuzzer struct{}

func newURLFuzzer() *urlFuzzer {
	return new(urlFuzzer)
}

func (u *urlFuzzer) run() {
	configuration := u.readConfiguration()
	fmt.Println(configuration)
}

func (u *urlFuzzer) readConfiguration() *configuration.Configuration {
	configurationFactory := configuration.NewFactory()
	return configurationFactory.FromCommandLine()
}
