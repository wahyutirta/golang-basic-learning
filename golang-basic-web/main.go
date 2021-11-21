package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/name", handler.NameHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// http.StripPrefix() forwards the handling of the request to one you specify as its parameter,
	// but before that it modifies the request URL by stripping off the specified prefix.

	// So for example in your case if the browser (or an HTTP client) requests the resource:

	// /static/example.txt
	// StripPrefix will cut the /static/ and forward the modified request to the handler returned by http.FileServer()
	// so it will see that the requested resource is

	// /example.txt
	// The Handler returned by http.FileServer() will look for and serve the content of a file relative to the folder (or rather FileSystem)
	// specified as its parameter (you specified "static" to be the root of static files).

	// Now since "example.txt" is in the static folder, you have to specify a relative path to that to get the corrent file path.

	log.Println("Starting web on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
