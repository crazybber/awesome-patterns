package main

import (
	"errors"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	tickets, timeout := 1, 3*time.Second
	s := New(tickets, timeout)
	if err := s.Aquire(); err != nil {
		panic(err)
	}
	// do important work
	spew.Dump("Test")
	if err := s.Release(); err != nil {
		panic(err)
	}
}

// https://github.com/tmrts/go-patterns/blob/master/synchronization/semaphore.md
var (
	ErrNoTickets      = errors.New("semaphore: could not aquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: Can not release the semaphore without acquiring it first")
)

type Interface interface {
	Aquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Aquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}
