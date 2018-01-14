package main

// The purpose of the work package is to show how you can use an unbuffered channel
// to create a pool of goroutines that will perform and control the amount of work that
// gets done concurrently. This is a better approach than using a buffered channel of
// some arbitrary static size that acts as a queue of work and throwing a bunch of goroutines at it.
// Unbuffered channels provide a guarantee that data has been exchanged
// between two goroutines. This approach of using an unbuffered channel allows the
// user to know when the pool is performing the work, and the channel pushes back
// when it can’t accept any more work because it’s busy. No work is ever lost or stuck in a
// queue that has no guarantee it will ever be worked on.

import "sync"

// Worker must be implemented by types that want to use
// the work pool.
// The Worker interface declares a single method called Task
type Worker interface {
	Task()
}

// Pool provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
// a struct named Pool is declared, which is the type that implements the
// pool of goroutines and will have methods that process the work. The type declares two
// fields, one named work, which is a channel of the Worker interface type, and a sync.WaitGroup named wg.
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New creates a new work pool.
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	// The for range loop blocks until there’s a Worker interface value to receive on the
	// work channel. When a value is received, the Task method is called. Once the work
	// channel is closed, the for range loop ends and the call to Done on the WaitGroup is
	// called. Then the goroutine terminates.
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run submits work to the pool.
// This method is used to submit work into the
// pool. It accepts an interface value of type Worker and sends that value through the
// work channel. Since the work channel is an unbuffered channel, the caller must wait
// for a goroutine from the pool to receive it. This is what we want, because the caller
// needs the guarantee that the work being submitted is being worked on once the call to Run returns.
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown waits for all the goroutines to shutdown.
// The Shutdown method in listing 7.33 does two things. First, it closes the work channel, which causes all of the goroutines in the pool to shut down and call the Done
// method on the WaitGroup. Then the Shutdown method calls the Wait method on the
// WaitGroup, which causes the Shutdown method to wait for all the goroutines to report
// they have terminated.
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
