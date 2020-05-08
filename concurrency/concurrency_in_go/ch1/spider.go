package ch1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var applicationStatus bool
var urls []string
var urlsProcessed int
var foundUrls []string
var fullText string
var totalURLCount int
var wg sync.WaitGroup
var v1 int

func readURLs(statusChannel chan int, textChannel chan string) {
	time.Sleep(time.Millisecond * 1)
	fmt.Println("Grabbing", len(urls), "urls")
	for i := 0; i < totalURLCount; i++ {
		fmt.Println("Url", i, urls[i])
		resp, _ := http.Get(urls[i])
		text, err := ioutil.ReadAll(resp.Body)
		textChannel <- string(text)
		if err != nil {
			fmt.Println("No HTML body")
		}
		statusChannel <- 0
	}
}

func addToScrapedText(textChannel chan string, processChannel chan bool) {
	for {
		select {
		case pC := <-processChannel:
			if pC == true {
				// hang on
			}
			if pC == false {
				close(textChannel)
				close(processChannel)
			}
		case tC := <-textChannel:
			fullText += tC
		}
	}
}

func evaluateStatus(statusChannel chan int, textChannel chan string, processChannel chan bool) {
	for {
		select {
		case status := <-statusChannel:
			fmt.Print(urlsProcessed, totalURLCount)
			urlsProcessed++
			if status == 0 {
				fmt.Println("Got url")
			}
			if status == 1 {
				close(statusChannel)
			}
			if urlsProcessed == totalURLCount {
				fmt.Println("Read all top-level URLs")
				processChannel <- false
				applicationStatus = false
			}
		}
	}
}

func RunSpider() {
	applicationStatus = true
	statusChannel := make(chan int)
	textChannel := make(chan string)
	processChannel := make(chan bool)
	totalURLCount = 0
	urls = append(urls, "http://www.mastergoco.com/index1.html")
	urls = append(urls, "http://www.mastergoco.com/index2.html")
	urls = append(urls, "http://www.mastergoco.com/index3.html")
	urls = append(urls, "http://www.mastergoco.com/index4.html")
	urls = append(urls, "http://www.mastergoco.com/index5.html")
	fmt.Println("Starting spider")
	urlsProcessed = 0
	totalURLCount = len(urls)
	go evaluateStatus(statusChannel, textChannel, processChannel)
	go readURLs(statusChannel, textChannel)
	go addToScrapedText(textChannel, processChannel)
	for {
		if applicationStatus == false {
			fmt.Println(fullText)
			fmt.Println("Done!")
			break
		}
		select {
		case sC := <-statusChannel:
			fmt.Println("Message on StatusChannel", sC)
		}
	}
}
