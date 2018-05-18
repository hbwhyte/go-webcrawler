package main

// run with "go run crawl.go"

import (
	"flag" // helps parse command line args
	"fmt"  // package for printing to the screen
	"os"   // gives you access to system calls
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify a start page")
		os.Exit(1)
	}
}
