package main

/* import (
	"fmt"
	"io/ioutil" // helps read/print the webpage
	"net/http"  // package to retrieve the webpage
)

func main() {
	resp, err := http.Get("http://6brand.com.com") // declaring 2 variables at once
	/* The way you handle errors in Go is to expect that functions you
	call will return two things and the second one will be an error. If the
	error is nil then you can continue but if it's not you need to handle it. */

/*fmt.Println("http transport error is: ", err)

	body, err := ioutil.ReadAll(resp.Body) // resp.Body is a ref to a stream of data

	fmt.Println("read error is: ", err)

	fmt.Println(string(body)) // cast from byte array to String

}
*/
