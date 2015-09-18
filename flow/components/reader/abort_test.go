package reader

import (
	"testing"

	"syscall"

	"github.com/stretchr/testify/assert"
)

func TestAbortNotNil(t *testing.T) {
	assert := assert.New(t)

	// when
	sut := newAbort()

	// then
	assert.NotNil(sut, "Abort must return a valid channel.")
}

func TestAbortNotifiesWhenSIGTERM(t *testing.T) {
	assert := assert.New(t)

	// given
	abort := newAbort()

	// when
	sut := abort.signal()
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	signal := <-sut

	// then
	assert.Equal(signal, syscall.SIGTERM, "SIGTERM awaited.")
}

func TestAbortNotifiesWhenSIGINT(t *testing.T) {
	assert := assert.New(t)

	// given
	abort := newAbort()

	// when
	sut := abort.signal()
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	signal := <-sut

	// then
	assert.Equal(signal, syscall.SIGINT, "SIGINT awaited.")
}
