package main

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool manages a set of resources that can be shared safely by
// multiply goroutines. The resource been managed must implement
// the io.Closer interface.
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed is returned when an Acquire returns on a closed pool.
var ErrPoolClosed = errors.New("Pool has been closed")

// New creates a pool that manage resources. A pool requires a function
// that can allocate a new resource and the size of the pool.
func New(fn func() (io.Closer, error), size uint32) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire retrieves a resource from the pool.
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// Check for free resources
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
func (p *Pool) Release(r io.Closer) {
	// Secure this operation with the Close operation.
	p.m.Lock()
	defer p.m.Unlock()
	// If the pool is closed, discard the resource
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
