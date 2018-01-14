package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Pool manages a set of resources that can be shared safely by
// multiply goroutines. The resource been managed must implement
// the io.Closer interface.
// declares a struct named Pool that allows
// the caller to create as many different pools as needed. Each pool can manage any type
// of resource as long as the type implements the io.Closer interface.
type Pool struct {
	// This mutex is used to keep all the operations against a Pool value-safe for multigoroutine access.
	m sync.Mutex
	// The second field is named resources and is declared as a channel of interface type io.Closer.
	// This channel will be created as a buffered channel and will contain the resources being shared.
	// Because an interface type is being used, the pool can manage any type of resource that
	// implements the io.Closer interface
	resources chan io.Closer
	// The factory field is of a function type. Any function that takes no parameters and
	// returns an io.Closer and an error interface value can be assigned to this field. The
	// purpose of this function is to create a new resource when the pool requires one. This
	// functionality is an implementation detail beyond the scope of the pool package and
	// needs to be implemented and supplied by the user using this package.
	factory func() (io.Closer, error)
	// This field is a flag that indicates the Pool is being shut down or is already shut down.
	closed bool
}

// ErrPoolClosed is returned when an Acquire returns on a closed pool.
// Creating error interface variables is a common practice in Go. This allows the caller
// to identify specific returned error values from any function or method within the package.
var ErrPoolClosed = errors.New("Pool has been closed")

// New creates a pool that manage resources. A pool requires a function
// that can allocate a new resource and the size of the pool.
// The function parameter represents a factory function that creates values of the resource being managed by the pool.
// The second parameter, size, represents the size of the buffered channel created to hold the resources.
func New(fn func() (io.Closer, error), size uint32) (*Pool, error) {
	// The first parameter, fn, is declared as a function type that accepts no parameters and
	// returns an io.Closer and an error interface value
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire retrieves a resource from the pool.
// This method returns a resource from the pool if one is available, or creates a new one for the call.
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// Check for free resources
	// using a select / case statement to check if there’s a resource in the buffered channel.
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		log.Println("Aquire: ", "Shared Resource")
		return r, nil
	default:
		log.Println("Aquire: ", "New Resource")
		return p.factory()
	}
}

// Release place a new resource into the pool.
// After a resource is acquired and no longer needed, it must be released back into
// the pool. This is where the Release method comes in.
func (p *Pool) Release(r io.Closer) {
	// Secure this operation with the Close operation.
	// The use of the mutex serves two purposes. First, it protects the
	// read on the closed flag on line 65 from happening at the same time as a write on this
	// flag in the Close method. Second, we don’t want to attempt to send on a closed channel because this will cause a panic. When the closed
	// field is false, we know the resources channel has been closed
	p.m.Lock()
	defer p.m.Unlock()
	// If the pool is closed, discard the resource
	// the Close method on the resource is called directly when the pool is
	// closed. This is because there’s no way to release the resource back into the pool. At
	// this point the pool has been both closed and flushed.
	if p.closed {
		r.Close()
		return
	}
	select {
	// Attempt to place the new resource on the queue
	case p.resources <- r:
		log.Println("Release: ", "In Queue")
	default:
		// If the queue is already at capacity we close the resource
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close will shutdown the pool and close all existing resources.
// Once the program is finished with the pool, it should call the Close method.
// The method closes and flushes the buffered channel on lines 98 and 101, closing any resources that exist until the channel is
// empty. All the code in this method must be executed by only one goroutine at a time.
// In fact, when this code is being executed, goroutines must also be
// prevented from executing code in the Release method. You’ll understand why this is important soon
func (p *Pool) Close() {
	// Secure this operation with release operation.
	p.m.Lock()
	defer p.m.Unlock()
	// If the pool is already closed, do not do anything.
	if p.closed {
		return
	}
	// set the pool as closed.
	p.closed = true
	// Close the channel before we drain the channel of its
	// resources
	close(p.resources)
	// Close the resources
	for r := range p.resources {
		r.Close()
	}
}

const (
	maxGoroutines   = 25
	pooledResources = 2
)

// dbConnection simulates a resource to share.
type dbConnection struct {
	ID int32
}

// Close implements the io.Closer interface so dbConnection
// can be managed by the pool. Close performs any resource
// release management.
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter provides support for giving each connection a unique id.
var idCounter int32

// createConnection is a factory method that will be called by
// the pool when a new connection is needed.
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	// Create the pool to manage our connections.
	p, err := New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}
	// Perform queries using connections from the pool.
	for query := 0; query < maxGoroutines; query++ {
		// Each goroutine needs its own copy of the query
		// value else they will all be sharing the same query
		// variable.
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	// Wait for the goroutines to finish.
	wg.Wait()
	// Close the pool.
	log.Println("Shutdown Program.")
	p.Close()
}

// performQueries tests the resource pool of connections.
func performQueries(query int, p *Pool) {
	// Acquire a connection from the pool.
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	// Release the connection back to the pool.
	defer p.Release(conn)
	// Wait to simulate a query response.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
