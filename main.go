package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var redirect bool

func init() {
	log.SetFlags(0)
	log.SetPrefix("certspy: ")

	flag.BoolVar(&redirect, "r", false, "")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: certspy [-r] domain\n")
	fmt.Fprintf(os.Stderr, "where:\n")
	fmt.Fprintf(os.Stderr, "       -r\tenable following redirects\n")
}

func main() {
	if len(flag.Args()) < 1 {
		log.Printf("error: must specifiy hostname.\n")
		usage()
		os.Exit(1)
	}

	// use custom http.Client that doesn't respect redirects
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if !redirect {
				return http.ErrUseLastResponse
			} else {
				return nil
			}
		},
	}

	hostname := flag.Arg(0)
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
