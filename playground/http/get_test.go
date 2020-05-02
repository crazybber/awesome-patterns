package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	counts := 2

	var wg sync.WaitGroup
	wg.Add(counts)

	for i := 0; i < counts; i++ {
		go func() {
			defer wg.Done()
			getValue()
		}()
	}
	wg.Wait()
}

func getValue() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	var url string
	if x > 5 {
		url = "http://localhost:48082/api/v1/device/5bae2ef4f37ba14693a5e4fc/command/5bae2ef4f37ba14693a5e4eb"
	} else {
		url = "http://localhost:48082/api/v1/device/5bae2d1bf37ba14693a5e4e9/command/5bae2d05f37ba14693a5e4e2"
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resBody))
}
