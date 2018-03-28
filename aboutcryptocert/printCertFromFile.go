package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/grantae/certinfo"
	"io/ioutil"
	"log"
	"flag"
)


var (
	pemfile = flag.String("pemfile","/tmp/a.pem","pem file path")
)

func main() {
	flag.Parse()

	// Read and parse the PEM certificate file
	pemData, err := ioutil.ReadFile(*pemfile)
	if err != nil {
		log.Fatal(err)
	}
	block, rest := pem.Decode([]byte(pemData))
	if block == nil || len(rest) > 0 {
		log.Fatal("Certificate decoding error")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// Print the certificate
	result, err := certinfo.CertificateText(cert)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}