package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("certspy: ")
}

func usage() {
	fmt.Printf("usage: certspy [domain]\n")
}

func main() {
	if len(os.Args) < 2 {
		log.Printf("error: must specifiy hostname.\n")
		usage()
		os.Exit(1)
	}

	hostname := os.Args[1]
	r, err := http.Get(fmt.Sprintf("https://%s/", hostname))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// get cert and print its alt names
	t := r.TLS.PeerCertificates[0]
	for _, s := range t.DNSNames {
		fmt.Printf("%s\n", s)
	}
}
