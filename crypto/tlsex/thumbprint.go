package tlsex

import (
	"crypto/sha1"
	"crypto/tls"
	"fmt"
)

// Thumbprint accepts a network address suitable for use with tls.Dial such as
// "www.posit.co:443" and an optional configuration which will default to the zero value
// of tls.Config and returns the SHA1 signature in hex of the address' root certificate.
func Thumbprint(address string, conf *tls.Config) (string, error) {
	if conf == nil {
		conf = &tls.Config{}
	}

	conn, err := tls.Dial("tcp", address, conf)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	peerCerts := conn.ConnectionState().PeerCertificates
	rootCert := peerCerts[len(peerCerts)-1]
	rootSha1 := sha1.Sum(rootCert.Raw)

	return fmt.Sprintf("%x", rootSha1), nil
}
