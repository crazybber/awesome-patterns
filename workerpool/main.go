package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
)

// Task encapsulates a work item that should go in a work
// pool.
type Task struct {
	// Err holds an error that occurred during a task. Its
	// result is only meaningful after Run has been called
	// for the pool that holds it.
	Err error

	f func() error
}

// NewTask initializes a new task based on a given work
// function.
func NewTask(f func() error) *Task {
	return &Task{f: f}
}

// Run runs a Task and does appropriate accounting via a
// given sync.WorkGroup.
func (t *Task) Run(wg *sync.WaitGroup) {
	t.Err = t.f()
	wg.Done()
}

type Pool struct {
	Tasks       []*Task
	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

// Run runs all work within the pool and blocks until it's
// finished.
func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	// all workers return
	close(p.tasksChan)

	p.wg.Wait()
}

// The work loop for any single goroutine.
func (p *Pool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}

func main() {
	f1 := func() error {
		fmt.Println("F1 Ran")
		time.Sleep(1 * time.Second)
		fmt.Println("F1 finished")
		return nil
	}

	f2 := func() error {
		fmt.Println("F2 Ran")
		time.Sleep(2 * time.Second)
		fmt.Println("F2 finished")
		return nil
	}
	f3 := func() error {
		fmt.Println("F3 Ran")
		time.Sleep(1 * time.Second)
		fmt.Println("F3 finished")
		return nil
	}

	f4 := func() error {
		fmt.Println("F4 Ran")
		time.Sleep(2 * time.Second)
		fmt.Println("F4 finished")
		return nil
	}
	tasks := []*Task{
		NewTask(f1),
		NewTask(f2),
		NewTask(f3),
		NewTask(f4),
	}

	p := NewPool(tasks, 4)
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			log.Error(task.Err)
			numErrors++
		}
		if numErrors >= 10 {
			log.Error("Too many errors.")
			break
		}
	}
}
