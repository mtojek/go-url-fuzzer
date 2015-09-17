package reader

import (
	"os"
	"os/signal"
	"syscall"
)

type abort struct{}

func newAbort() *abort {
	return new(abort)
}

func (a *abort) signal() chan os.Signal {
	receiver := make(chan os.Signal, 1)
	signal.Notify(receiver, syscall.SIGINT, syscall.SIGTERM)
	return receiver
}
