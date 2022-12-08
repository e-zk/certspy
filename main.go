package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("certspy: ")

	// use flag solely for usage handling
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: certspy [domain]\n")
}

func main() {
	if len(os.Args) < 2 {
		log.Printf("error: must specifiy hostname.\n")
		usage()
		os.Exit(1)
	}

	// use custom http.Client that doesn't respect redirects
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	hostname := os.Args[1]
	r, err := c.Get(fmt.Sprintf("https://%s/", hostname))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// get cert and print its alt names
	t := r.TLS.PeerCertificates[0]
	for _, s := range t.DNSNames {
		fmt.Printf("%s\n", s)
	}
}
