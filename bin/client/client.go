package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	pem, err := ioutil.ReadFile("./etc/server.pem")
	if err != nil {
		log.Fatalf("failed to read server pem: %v", err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(pem))

	conf := &tls.Config{
		ServerName: "127.0.0.1",
		RootCAs:    certPool,
	}

	conn, err := tls.Dial("tcp", ":443", conf)
	if err != nil {
		log.Fatalf("failed to dial :443: %v", err)
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

	println(string(buf[:n]))
}
