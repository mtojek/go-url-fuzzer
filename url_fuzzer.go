package main

import (
	"github.com/mtojek/go-url-fuzzer/configuration"
	"fmt"
)

func main() {
	configurationFactory := configuration.NewConfigurationFactory()
	configuration := configurationFactory.FromCommandLine()

	fmt.Println(configuration)
}
