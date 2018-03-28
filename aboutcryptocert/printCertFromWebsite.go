package main

import (
	"crypto/tls"
	"fmt"
	"github.com/grantae/certinfo"
	"log"
	"flag"
)

var (
	website = flag.String("website","www.jingoal.com:443","the website to access")
)
/**
 * 打印一个网站的cert证书。
 */
func main() {
	flag.Parse()

	// Connect to google.com
	cfg := tls.Config{}
	conn, err := tls.Dial("tcp", *website, &cfg)
	if err != nil {
		log.Fatalln("TLS connection failed: " + err.Error())
	}
	// Grab the last certificate in the chain
	certChain := conn.ConnectionState().PeerCertificates
	log.Println("cert Chain len:", len(certChain))
	//cert := certChain[len(certChain)-1]

	for i, cert := range certChain {
		// Print the certificate
		result, err := certinfo.CertificateText(cert)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n certchain[%d], content: \n %+v \n",i,result)
	}
}