package factory

import "log"

var speakFuncs = make(map[string]interface{})

func say(funcName string) {
	speakFunc, ok := speakFuncs[funcName]
	if !ok {
		log.Println("speakFunc not exist")
	} else {
		speakFunc.(func())()
	}
}
