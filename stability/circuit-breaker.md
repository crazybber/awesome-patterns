# Circuit Breaker Pattern

Similar to electrical fuses that prevent fires when a circuit that is connected
to the electrical grid starts drawing a high amount of power which causes the
wires to heat up and combust, the circuit breaker design pattern is a fail-first
mechanism that shuts down the circuit, request/response relationship or a
service in the case of software development, to prevent bigger failures.

**Note:** The words "circuit" and "service" are used synonymously through this
document.

## Implementation

Below is the implementation of a very simple circuit breaker to illustrate the purpose
of the circuit breaker design pattern.

### Operation Counter

`circuit.Counter` is a simple counter that records success and failure states of
a circuit along with a timestamp and calculates the consecutive number of
failures.

```go
package circuit

import (
    "time"
)

type State int

const (
    UnknownState State = iota
    FailureState
    SuccessState
)

type Counter interface {
    Count(State)
    ConsecutiveFailures() uint32
    LastActivity() time.Time
    Reset()
}
```

### Circuit Breaker

Circuit is wrapped using the `circuit.Breaker` closure that keeps an internal operation counter.
It returns a fast error if the circuit has failed consecutively more than the specified threshold.
After a while it retries the request and records it.

**Note:** Context type is used here to carry deadlines, cancellation signals, and
other request-scoped values across API boundaries and between processes.

```go
package circuit

import (
    "context"
    "time"
)

type Circuit func(context.Context) error

func Breaker(c Circuit, failureThreshold uint32) Circuit {
    cnt := NewCounter()
	expired := time.Now()
	currentState := StateClosed
    
   	return func(ctx context.Context) error {


		//handle statue transformation for timeout
		if currentState == StateOpen {
			nowt := time.Now()
			if expired.Before(nowt) || expired.Equal(nowt) {
				currentState = StateHalfOpen 
				cnt.ConsecutiveSuccesses = 0 
			}
		}

		switch currentState {
		case StateOpen:
			return ErrServiceUnavailable 
		case StateHalfOpen:
			if err := c(ctx); err != nil {
				currentState = StateOpen
				expired = time.Now().Add(defaultTimeout) //Reset
				return err
			}
			cnt.Count(SuccessState)
			if cnt.ConsecutiveSuccesses > defaultSuccessThreshold {
				currentState = StateClosed
				cnt.ConsecutiveFailures = 0
			}

		case StateClosed:
			if err := c(ctx); err != nil {
				cnt.Count(FailureState)
			}
		}
		return nil
	}
}
```

## Related Works

- [sony/gobreaker](https://github.com/sony/gobreaker) is a well-tested and intuitive circuit breaker implementation for real-world use cases.
