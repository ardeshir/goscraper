package main

import (
	"log"
	"net/http"
	"os"
)

func test() {

	// Create the variables for the response and error
	var r *http.Response
	var err error

	// Request index.html from ardeshir.org
	r, err = http.Get("http://ardeshir.org/index.html")

	// If there is a problem accessing the server, kill the program and print the error the console
	if err != nil {
		panic(err)
	}

	// Check the status code returned by the server
	if r.StatusCode == 200 {
		// The request was successful!
		var webPageContent []byte

		// We know the size of the response is ~ 1270
		var bodyLength int = 3000

		// Initialize the byte array to the size of the data
		webPageContent = make([]byte, bodyLength)

		// Read the data from the server
		r.Body.Read(webPageContent)

		var out *os.File
		out, err = os.OpenFile("index.html", os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			panic(err)
		}

		// Write the contents to a file
		out.Write(webPageContent)
		out.Close()

	} else {
		log.Fatal("Failed to retrieve the webpage. Received status code", r.Status)
	}

}
