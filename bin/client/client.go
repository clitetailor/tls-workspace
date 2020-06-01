package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	cert, err := ioutil.ReadFile("./etc/ca.crt")
	if err != nil {
		log.Fatalf("failed to read ca cert: %v", err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(cert))

	conf := &tls.Config{
		ServerName: "localhost",
		RootCAs:    certPool,
	}

	conn, err := tls.Dial("tcp", ":1227", conf)
	if err != nil {
		log.Fatalf("failed to dial :1227: %v", err)
	}
	defer conn.Close()

	n, err := conn.Write([]byte("Hello, World!\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	fmt.Print(string(buf[:n]))
}
