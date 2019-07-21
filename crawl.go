package main

// run with "go run crawl.go"

import (
	"crypto/tls" // for low level transport customizations
	"database/sql"
	"flag"     // helps parse command line args
	"fmt"      // package for printing to the screen
	"net/http" // package to retrieve the webpage
	"net/url"  // standard library to fix urls
	"os"       // gives you access to system calls

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackdanger/collectlinks" // library for parsing links
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()
	var depth = 0

	if len(args) < 1 {
		fmt.Println("Please specify a start page")
		os.Exit(1)
	}

	queue := make(chan string) // New channel that retrieves and delivers strings

	go func() { // run this asyncronously
		queue <- args[0] // put args into the channel
		fmt.Println("async go func")
	}()

	for uri := range queue {
		enqueue(uri, queue, depth)
		fmt.Println("for uri enqueue")
	}
}

func enqueue(uri string, queue chan string, depth int) {
	fmt.Println("Fetching", uri, "Level:", depth)
	visited[uri] = true // record that the page was visited

	// This section allows it to search https-secured site by disabling SSL verification
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}

	resp, err := client.Get(uri) // declaring 2 variables at once
	/* The way you handle errors in Go is to expect that functions you
	call will return two things and the second one will be an error. If the
	error is nil then you can continue but if it's not you need to handle it. */
	if err != nil {
		return
	}
	defer resp.Body.Close() // gotta close those TCP connections

	links := collectlinks.All(resp.Body)

	for _, link := range links { // 'for' + 'range' in Go is like an enhanced for loop
		absolute := fixUrl(link, uri) // fix link before enqueing
		// fmt.Println("for links loop")
		// We'll set invalid URLs to blank strings so let's never send those to the channel.
		if uri != "" {
			if !visited[absolute] {
				go func() {
					queue <- absolute
				}()
			}
		}
	}
}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}
