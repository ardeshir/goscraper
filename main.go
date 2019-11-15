package main

func scrapeSite(url string, statusChan chan map[string]string) {

  // Performing scraping operations...

  statusChan <- map[string]string{url: "DONE"}
}

func main() {
  siteStatus := map[string]string{
    "http://ardeshir.org/index.html": "READY",
    "http://ardeshir.org/intro": "READY",
    "http://ardeshir.org/aws": "READY",
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

