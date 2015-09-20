package localserver

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

type tlsConfigProvider struct {
}

func newTLSConfigProvider() *tlsConfigProvider {
	return new(tlsConfigProvider)
}

func (t *tlsConfigProvider) Provide(caPemPath, caKeyPath string) *tls.Config {
	caPem, _ := ioutil.ReadFile("ca.pem")
	ca, error := x509.ParseCertificate(caPem)
	if nil != error {
		log.Fatalln(error)
	}

	caKey, _ := ioutil.ReadFile("ca.key")
	priv, error := x509.ParsePKCS1PrivateKey(caKey)
	if nil != error {
		log.Fatalln(error)
	}

	pool := x509.NewCertPool()
	pool.AddCert(ca)

	cert := tls.Certificate{
		Certificate: [][]byte{caPem},
		PrivateKey:  priv,
	}

	config := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    pool,
	}
	config.Rand = rand.Reader
	return config
}
