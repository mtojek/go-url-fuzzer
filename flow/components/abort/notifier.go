package abort

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Notifier informs about received signal SIGINT or SIGTERM.
type Notifier struct {
	waitGroup sync.WaitGroup
	receiver  chan os.Signal
}

// NewNotifier creates an instance of a new notifier.
func NewNotifier() *Notifier {
	receiver := make(chan os.Signal, 1)
	return &Notifier{receiver: receiver}
}

// Notify waits until a signal is received.
func (n *Notifier) Notify() bool {
	signal.Notify(n.receiver, syscall.SIGINT, syscall.SIGTERM)

	n.waitGroup.Add(1)
	go n.awaitSignal()
	n.waitGroup.Wait()

	log.Printf("Signal received. Aborting.")
	return true
}

func (n *Notifier) awaitSignal() {
	defer n.waitGroup.Done()
	<-n.receiver
}
