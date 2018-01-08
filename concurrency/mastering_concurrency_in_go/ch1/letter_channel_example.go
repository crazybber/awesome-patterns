package ch1

import (
	"fmt"
	"strings"
)

var initialString string
var initialBytes []byte
var stringLength int
var finalString string
var lettersProcessed int
var applicationStatusL bool

func getLetters(gQ chan string) {
	for i := range initialBytes {
		gQ <- string(initialBytes[i])
	}
}
func capitalizeLetters(gQ chan string, sQ chan string) {
	for {
		if lettersProcessed >= stringLength {
			applicationStatusL = false
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
	applicationStatusL = true
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
		if applicationStatusL == false {
			fmt.Println("Done")
			fmt.Println(finalString)
			break
		}
	}
}
