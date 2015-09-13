package main

import (
	"fmt"

	"github.com/mtojek/go-url-fuzzer/configuration"
)

func main() {
	configuration := readConfiguration()
	fmt.Println(configuration)
}

func readConfiguration() *configuration.Configuration {
	configurationFactory := configuration.NewConfigurationFactory()
	return configurationFactory.FromCommandLine()
}
