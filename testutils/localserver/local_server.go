package localserver

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/etix/stoppableListener"
)

// LocalServer provides a simple, local HTTP server.
type LocalServer struct {
	hostPort string
	scheme   string
	listener *stoppableListener.StoppableListener
}

// NewLocalServer creates new instance of a HTTP server. It requires host, port and protocol scheme.
func NewLocalServer(hostPort string, scheme string) *LocalServer {
	return &LocalServer{hostPort: hostPort, scheme: scheme}
}

// Start runs the server.
func (l *LocalServer) Start() {
	listener, _ := net.Listen("tcp", l.hostPort)
	l.listener = stoppableListener.Handle(listener)
	go http.Serve(l.listener, nil)
	l.waitUntilReady()
}

func (l *LocalServer) waitUntilReady() {
	var ready bool

	for !ready {
		_, error := http.Get(l.scheme + "://" + l.hostPort)
		ready = (nil == error)

		log.Printf("Waiting for local server to start: %v, error: %v\n", l.hostPort, error)
		time.Sleep(10 * time.Millisecond)
	}

	log.Println("Local server started.")
}

// Stop method stops the running server.
func (l *LocalServer) Stop() {
	l.listener.Stop <- true
}
