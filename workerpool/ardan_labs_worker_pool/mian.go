package main

// article https://www.ardanlabs.com/blog/2013/05/thread-pooling-in-go-programming.html
import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type MyWork struct {
	Name      string
	BirthYear int
	WP        *WorkPool
}

func (mw *MyWork) DoWork(workRoutine int) {
	fmt.Printf("%s : %d\n", mw.Name, mw.BirthYear)
	fmt.Printf("Q:%d R:%d\n", mw.WP.QueuedWork(), mw.WP.ActiveRoutines())

	// Simulate some delay
	time.Sleep(100 * time.Millisecond)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// we create a thread pool where the number of routines to use is based on the number of cores we have on the machine.
	// This means we have a routine for each core. You can’t do any more work if each core is busy.
	// Again, performance testing will determine what this number should be. The second parameter is the size of the queue.
	// In this case I have made the queue large enough to handle all the requests coming in.
	workPool := New(runtime.NumCPU(), 800)

	shutdown := false // Race Condition, Sorry

	go func() {
		for i := 0; i < 1000; i++ {
			work := MyWork{
				Name:      "A" + strconv.Itoa(i),
				BirthYear: i,
				WP:        workPool,
			}

			if err := workPool.PostWork("routine", &work); err != nil {
				fmt.Printf("ERROR: %s\n", err)
				time.Sleep(100 * time.Millisecond)
			}

			if shutdown == true {
				return
			}
		}
	}()

	fmt.Println("Hit any key to exit")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	shutdown = true

	fmt.Println("Shutting Down")
	workPool.Shutdown("routine")
}
