package main

import (
	"os"
	"testing"

	"net"
	"net/http"

	"log"
	"time"

	"github.com/etix/stoppableListener"
	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "127.0.0.1:10600"
	server := startLocalServer(hostPort)

	setCommandLineArgs("input-data/fuzz_01.txt", "http://"+hostPort)
	sut := newURLFuzzer()

	// when
	configuration := sut.readConfiguration()
	server.Stop <- true

	// then
	assert.NotNil(configuration, "Simple configuration should be read from command line.")
}

func setCommandLineArgs(customArguments ...string) {
	os.Args = os.Args[:len(os.Args)-1] // remove test.v flag
	for _, customArgument := range customArguments {
		os.Args = append(os.Args, customArgument)
	}
}

func startLocalServer(hostPort string) *stoppableListener.StoppableListener {
	listener, _ := net.Listen("tcp", hostPort)
	stoppable := stoppableListener.Handle(listener)
	go http.Serve(stoppable, nil)
	waitUntilReady(hostPort)

	return stoppable
}

func waitUntilReady(hostPort string) {
	var ready bool

	for !ready {
		_, error := http.Get("http://" + hostPort)
		ready = (nil == error)

		log.Printf("Waiting for local server to start: %v, error: %v\n", hostPort, error)
		time.Sleep(10 * time.Millisecond)
	}

	log.Println("Local server started.")
}
