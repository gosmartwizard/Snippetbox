package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "\n<h1> Welcome to the Go web programming </h1>")

	w.Write([]byte("Hello from Snippetbox"))
}
func main() {
	// Use the http.NewServeMux() function to initialize a new servemux
	mux := http.NewServeMux()

	// register the home function as the handler for the "/" URL pattern.
	mux.HandleFunc("/", home)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4949")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Print("Starting server on :4949")
	err := http.ListenAndServe(":4949", mux)
	log.Fatal(err)
}
