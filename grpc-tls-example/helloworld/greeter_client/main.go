package main

import (
	"log"
	"os"

	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = ":50051"
	defaultName = "world"
)

func main() {
	/*certificate, err := tls.LoadX509KeyPair(
		"out/org2_ca1.com.crt",
		"out/org2_ca1.com.key",
	)*/
	certificate, err := tls.LoadX509KeyPair(
		"out/org3_ca2.com.crt",
		"out/org3_ca2.com.key",
	)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile("out/ca1.crt")
	if err != nil {
		log.Fatalf("failed to read ca cert: %s", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		log.Fatal("failed to append certs")
	}
	/*
		bs, err = ioutil.ReadFile("out/ca2.crt")
		if err != nil {
			log.Fatalf("failed to read ca cert: %s", err)
		}

		ok = certPool.AppendCertsFromPEM(bs)
		if !ok {
			log.Fatal("failed to append certs[2]")
		}
	*/
	transportCreds := credentials.NewTLS(&tls.Config{
		ServerName:   "org1_ca1.com",
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	dialOption := grpc.WithTransportCredentials(transportCreds)

	// make your client
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, dialOption)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
