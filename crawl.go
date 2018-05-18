package main

// run with "go run crawl.go"

import (
	"crypto/tls" // for low level transport customizations
	"flag"       // helps parse command line args
	"fmt"        // package for printing to the screen
	"net/http"   // package to retrieve the webpage
	"os"         // gives you access to system calls

	"github.com/jackdanger/collectlinks" // library for parsing links
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify a start page")
		os.Exit(1)
	}

	// This section allows it to search https-secured site by disabling SSL verification
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}

	resp, err := client.Get(args[0]) // declaring 2 variables at once
	/* The way you handle errors in Go is to expect that functions you
	call will return two things and the second one will be an error. If the
	error is nil then you can continue but if it's not you need to handle it. */
	if err != nil {
		return
	}
	defer resp.Body.Close() // gotta close those TCP connections

	links := collectlinks.All(resp.Body)

	for _, link := range links { // 'for' + 'range' in Go is like an enhanced for loop
		fmt.Println(link)
	}
	// _ is a placeholder for where index would be, but since we don't use index we use _

}
