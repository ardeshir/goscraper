package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

// build http client to fetch
func getClient(url string) {

	var r *http.Response
	var err error

	// Request index.html from ardeshir.org
	r, err = http.Get(url)

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
		out, err = os.OpenFile(path.Base(url), os.O_CREATE|os.O_WRONLY, 0664)
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

func scrapeSite(url string, statusChan chan map[string]string) {
	// Performing scraping operations...
	getClient(url)
	// channel
	statusChan <- map[string]string{url: "DONE"}
}

func main() {

	siteStatus := map[string]string{
		"http://ardeshir.org/index.html": "READY",
		"http://ardeshir.org/intro":      "READY",
		"http://ardeshir.org/aws":        "READY",
	}

	updatesChan := make(chan map[string]string)

	numberCompleted := 0
	for site := range siteStatus {
		siteStatus[site] = "WORKING"
		go scrapeSite(site, updatesChan)
	}

	for update := range updatesChan {
		for url, status := range update {
			siteStatus[url] = status
			numberCompleted++
		}
		if numberCompleted == len(siteStatus) {
			close(updatesChan)
		}
	}
}
