package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGophers(t *testing.T) {
	var gophers = []string{"tom", "peter", "john", "brown"}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(gophers))

	for i := 0; i < len(gophers); i++ {
		go func(wg *sync.WaitGroup, name string) {
			sleepyGopherSnore(name)
			wg.Done()
		}(&waitGroup, gophers[i])
	}

	waitGroup.Wait()

	fmt.Println("All done")
}

func sleepyGopherSnore(name string) {
	fmt.Println(name, ": ... start snore")
	time.Sleep(2 * time.Second)
	fmt.Println(name, ": ... snore")
}
