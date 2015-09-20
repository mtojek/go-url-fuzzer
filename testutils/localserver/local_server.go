package localserver

import (
	"log"
	"net"
	"net/http"
	"time"

	"crypto/tls"

	"github.com/etix/stoppableListener"
)

const (
	defaultCaKey = "server_ca.key"
	defaultCaPem = "server_ca.pem"
)

// LocalServer provides a simple, local HTTP server.
type LocalServer struct {
	hostPort string
	scheme   string
	listener *stoppableListener.StoppableListener

	tlsConfigProvider *tlsConfigProvider
}

// NewLocalServer creates new instance of a HTTP server. It requires host, port and protocol scheme.
func NewLocalServer(hostPort string, scheme string) *LocalServer {
	tlsConfigProvider := newTLSConfigProvider()
	return &LocalServer{hostPort: hostPort, scheme: scheme, tlsConfigProvider: tlsConfigProvider}
}

// Start runs the server basing on chosen protocol scheme.
func (l *LocalServer) Start() {
	if l.scheme == "http" {
		l.StartHTTP()
	} else if l.scheme == "https" {
		l.StartTLS(defaultCaKey, defaultCaPem)
	} else {
		log.Fatalf("Unknown scheme specified: %v\n", l.scheme)
	}
}

// StartHTTP runs the HTTP server.
func (l *LocalServer) StartHTTP() {
	listener, error := net.Listen("tcp", l.hostPort)
	if nil != error {
		log.Fatalln(error)
	}

	l.listener = stoppableListener.Handle(listener)
	l.startServing()
}

// StartTLS runs the TLS server.
func (l *LocalServer) StartTLS(caPemPath, caKeyPath string) {
	config := l.tlsConfigProvider.Provide(caPemPath, caKeyPath)
	listener, error := tls.Listen("tcp", l.hostPort, config)
	if nil != error {
		log.Fatalln(error)
	}

	l.listener = stoppableListener.Handle(listener)
	l.startServing()
}

func (l *LocalServer) startServing() {
	go http.Serve(l.listener, nil)
	l.waitUntilReady()
}

func (l *LocalServer) waitUntilReady() {
	var ready bool

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{
		Transport: tr,
	}

	for !ready {
		if _, error := client.Get(l.scheme + "://" + l.hostPort); nil != error {
			log.Printf("Waiting for local server to start: %v, error: %v\n", l.hostPort, error)
			time.Sleep(10 * time.Millisecond)
		} else {
			ready = true
		}
	}

	log.Printf("Local server started: %v\n", l.hostPort)
}

// Stop method stops the running server.
func (l *LocalServer) Stop() {
	l.listener.Stop <- true
	log.Printf("Local server stopped: %v\n", l.hostPort)
}
