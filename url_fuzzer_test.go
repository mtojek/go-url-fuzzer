package main

import (
	"os"
	"testing"

	"github.com/mtojek/localserver"
	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	scheme := "http"
	hostPort := "127.0.0.1:10600"
	server := localserver.NewLocalServer(hostPort, scheme)
	server.Start()

	setCommandLineArgs("resources/input-data/fuzz_01.txt", scheme+"://"+hostPort)
	sut := newURLFuzzer()

	// when
	configuration := sut.readConfiguration()
	server.Stop()

	// then
	assert.NotNil(configuration, "Simple configuration should be read from command line.")
}

func setCommandLineArgs(customArguments ...string) {
	os.Args = os.Args[:len(os.Args)-1] // remove test.v flag
	for _, customArgument := range customArguments {
		os.Args = append(os.Args, customArgument)
	}
}
