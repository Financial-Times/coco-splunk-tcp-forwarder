package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"io"
	"log"
	"os"
)

var (
	fwdAddress = flag.String("fwdAddress", "", "Address to forward to")
)

func main() {
	flag.Parse()
	con, err := tls.Dial("tcp", *fwdAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		i, err := io.Copy(con, bufio.NewReader(os.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("tcp forwarder exiting. wrote a total of %d bytes while running.", i)
	}
}
