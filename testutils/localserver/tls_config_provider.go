package localserver

import (
	"crypto/rand"
	"crypto/tls"
	"log"
)

type tlsConfigProvider struct {
}

func newTLSConfigProvider() *tlsConfigProvider {
	return new(tlsConfigProvider)
}

func (t *tlsConfigProvider) Provide(caPemPath, caKeyPath string) *tls.Config {
	cert, error := tls.LoadX509KeyPair(caPemPath, caKeyPath)
	if nil != error {
		log.Fatalln("tls.LoadX509KeyPair(caPemPath, caKeyPath): ", error)
	}

	config := &tls.Config{
		ClientAuth:   tls.NoClientCert,
		Certificates: []tls.Certificate{cert},
	}
	config.Rand = rand.Reader
	return config
}
