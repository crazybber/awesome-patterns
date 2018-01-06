package divide_and_conquer

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func RunDivideAndConquer() {
	type in struct {
		a int
		b int
	}

	type out struct {
		source string
		result int
	}

	evaluators := []Evaluator{
		EvaluatorFunc(func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a + i.b
			return out{"Plus", r}, nil
		}),
		EvaluatorFunc(func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a * i.b
			return out{"Multi", r}, nil
		}),
		EvaluatorFunc(func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a - i.b
			return out{"min", r}, nil
		}),
		EvaluatorFunc(func(inV interface{}) (interface{}, error) {
			i := inV.(in)
			r := i.a / i.b
			return out{"divider", r}, nil
		}),
	}

	r, errors := DivideAndConquer(in{2, 3}, evaluators, 10*time.Millisecond)
	spew.Dump(r, errors)
}

type Evaluator interface {
	Evaluate(data interface{}) (interface{}, error)
	Name() string
}
type EvaluatorFunc func(interface{}) (interface{}, error)

func (ef EvaluatorFunc) Evaluate(in interface{}) (interface{}, error) {
	return ef(in)
}

func (ef EvaluatorFunc) Name() string {
	return runtime.FuncForPC(reflect.ValueOf(ef).Pointer()).Name()
}

func DivideAndConquer(data interface{}, evaluators []Evaluator, timeout time.Duration) ([]interface{}, []error) {
	gather := make(chan interface{}, len(evaluators))
	errors := make(chan error, len(evaluators))
	for _, v := range evaluators {
		go func(e Evaluator) {
			// Why not just use an unbuffered channel? The answer is that we don’t want to leak any goroutines.
			// While the Go runtime is capable of handling thousands or hundreds of thousands of goroutines at a time,
			// each goroutine does use some resources, so you don’t want to leave them hanging around when
			// you don’t have to. If you do, a long-running Go program will start performing poorly
			ch := make(chan interface{}, 1)
			ech := make(chan error, 1)
			go func() {
				result, err := e.Evaluate(data)
				if err != nil {
					errors <- err
				} else {
					ch <- result
				}
			}()
			select {
			case r := <-ch:
				gather <- r
			case err := <-ech:
				errors <- err
			case <-time.After(timeout):
				errors <- fmt.Errorf("%s timeout after %v on %v", e.Name(), timeout, data)
			}
		}(v)
	}
	out := make([]interface{}, 0, len(evaluators))
	errs := make([]error, 0, len(evaluators))
	for range evaluators {
		select {
		case r := <-gather:
			out = append(out, r)
		case e := <-errors:
			errs = append(errs, e)
		}
	}
	return out, errs
}
