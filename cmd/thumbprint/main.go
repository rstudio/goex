// Teensy thumbprint thing adapted from https://www.jvt.me/posts/2022/05/06/go-cert-fingerprint/
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rstudio/goex/crypto/tlsex"
)

var (
	VersionString = "???"
)

func main() {
	conf := &tls.Config{}

	flag.BoolVar(&conf.InsecureSkipVerify, "insecure-skip-verify", false, "do not verify the TLS remote")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: thumbprint <network-location>\n")
		flag.CommandLine.PrintDefaults()
	}

	versionFlag := flag.Bool("version", false, "show version and exit")

	flag.Parse()

	if *versionFlag {
		fmt.Println(VersionString)
		return
	}

	if len(os.Args) < 2 {
		log.Fatal("missing required positional argument for network location")
	}

	s, err := tlsex.Thumbprint(os.Args[1], conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
