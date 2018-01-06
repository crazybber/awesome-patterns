package ch1

import (
	"fmt"
	"strings"
	"sync"
)

var initialString string
var initialBytes []byte
var stringLength int
var finalString string
var lettersProcessed int
var applicationStatus bool
var wg sync.WaitGroup

func getLetters(gQ chan string) {
	for i := range initialBytes {
		gQ <- string(initialBytes[i])
	}
}
func capitalizeLetters(gQ chan string, sQ chan string) {
	for {
		if lettersProcessed >= stringLength {
			applicationStatus = false
			break
		}
		select {
		case letter := <-gQ:
			capitalLetter := strings.ToUpper(letter)
			finalString += capitalLetter
			lettersProcessed++
		}
	}
}
func RunLetter() {
	applicationStatus = true
	getQueue := make(chan string)
	stackQueue := make(chan string)
	initialString = `Four score and seven years ago our fathers
brought forth on this continent, a new nation, conceived in
Liberty, and dedicated to the proposition that all men are
created equal.`
	initialBytes = []byte(initialString)
	stringLength = len(initialString)
	lettersProcessed = 0
	fmt.Println("Let's start capitalizing")
	go getLetters(getQueue)
	capitalizeLetters(getQueue, stackQueue)
	close(getQueue)
	close(stackQueue)
	for {
		if applicationStatus == false {
			fmt.Println("Done")
			fmt.Println(finalString)
			break
		}
	}
}
