package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	setCommandLineArgs("input-data/fuzz_01.txt", "http://domain.tld/")
	sut := newURLFuzzer()

	// when
	configuration := sut.readConfiguration()

	// then
	assert.NotNil(configuration, "Simple configuration should be read from command line.")
}

func setCommandLineArgs(customArguments ...string) {
	os.Args = os.Args[:len(os.Args)-1] // remove test.v flag
	for _, customArgument := range customArguments {
		os.Args = append(os.Args, customArgument)
	}
}
