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
	br := bufio.NewReader(os.Stdin)
	con, err := tls.Dial("tcp", *fwdAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	for {
		str, err := br.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		_, err = con.Write([]byte(str))
		if err != nil {
			log.Println(err)
		}
	}
}
