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
of the circuit breaker design pattern, only considering Open/Closed 

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

type counters struct {
   ConsecutiveFailures  uint32
   ConsecutiveSuccesses uint32
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

var canRetry = func(cnt counters, failureThreshold uint32) bool {
	backoffLevel := cnt.ConsecutiveFailures - failureThreshold
	// Calculates when should the circuit breaker resume propagating requests
	// to the service
	shouldRetryAt := cnt.LastActivity().Add(time.Second << backoffLevel)
	return time.Now().After(shouldRetryAt)
}

func Breaker(c Circuit, failureThreshold uint32) Circuit {
	cnt := counters{}

	//ctx can be used hold parameters
	return func(ctx context.Context) error {

		if cnt.ConsecutiveFailures >= failureThreshold {
			if !canRetry(cnt, failureThreshold) {
				// Fails fast instead of propagating requests to the circuit since
				// not enough time has passed since the last failure to retry
				return ErrServiceUnavailable
			}
		}
		// Unless the failure threshold is exceeded the wrapped service mimics the
		// old behavior and the difference in behavior is seen after consecutive failures
	         if err := c(ctx); err != nil {
			cnt.Count(FailureState)
			return err
		}

		cnt.Count(SuccessState)
		return nil
	}
}
```

## Related Contents

- [Circuit Breaker Pattern](https://docs.microsoft.com/en-us/previous-versions/msp-n-p/dn589784(v=pandp.10)?redirectedfrom=MSDN)

- [sony/gobreaker](https://github.com/sony/gobreaker) is a well-tested and intuitive circuit breaker implementation for real-world use cases.
