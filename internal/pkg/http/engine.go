package http

import (
	"airbnb-messaging-be/internal/pkg/env"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func DefaultSameSite() http.SameSite {
	if env.CONFIG.Stage != string(env.StageLocal) {
		return http.SameSiteStrictMode
	} else {
		return http.SameSiteNoneMode
	}
}

func CreateTlsConfiguration(certFile, keyFile, caFile string, tlsSkipVerify bool) (t *tls.Config) {
	// "certificate", "", "The optional certificate file for client authentication"
	// "key", "", "The optional key file for client authentication"
	// "ca", "", "The optional certificate authority file for TLS client authentication"
	// "tls-skip-verify", false, "Whether to skip TLS server cert verification"

	if certFile != "" && keyFile != "" && caFile != "" {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}

		caCert, err := os.ReadFile(caFile)
		if err != nil {
			log.Fatal(err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		t = &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: tlsSkipVerify,
		}
	}
	// will be nil by default if nothing is provided
	return t
}
