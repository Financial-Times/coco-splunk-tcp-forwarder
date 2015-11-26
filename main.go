package main

import (
	"bufio"
	"crypto/tls"
	"io"
	"log"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	con, err := tls.Dial("tcp", "ingest.splunk.glb.ft.com:8443", nil)
	if err != nil {
		log.Fatal(err)
	}
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
