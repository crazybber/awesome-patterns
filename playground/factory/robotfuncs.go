package factory

import "log"

func sayHi() {
	log.Print("Hi~")
}

func sayHello() {
	log.Print("Hello~")
}

func init() {
	speakFuncs["hi"] = sayHi
	speakFuncs["hello"] = sayHello
}
