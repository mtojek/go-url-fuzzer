package main

import (
	"fmt"

	"github.com/mtojek/go-url-fuzzer/configuration"
)

func main() {
	configurationFactory := configuration.NewConfigurationFactory()
	configuration := configurationFactory.FromCommandLine()

	fmt.Println(configuration)
}
