<p align="center">
  <img src="/gopher.png" height="400">
  <h1 align="center">
    Go Patterns
    <br>
    <a href="http://travis-ci.org/tmrts/go-patterns"><img alt="build-status" src="https://img.shields.io/badge/build-passing-brightgreen.svg?style=flat-square" /></a>
    <a href="https://github.com/sindresorhus/awesome" ><img alt="awesome" src="https://img.shields.io/badge/awesome-%E2%9C%93-ff69b4.svg?style=flat-square" /></a>
    <a href="https://github.com/tmrts/go-patterns/blob/master/LICENSE" ><img alt="license" src="https://img.shields.io/badge/license-Apache%20License%202.0-E91E63.svg?style=flat-square" /></a>
  </h1>
</p>

A curated collection of idiomatic design & application patterns for Go language.

## Creational Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Abstract Factory](/creational/abstract_factory.md) | Provides an interface for creating families of releated objects | âœ|
| [Builder](/creational/builder.md) | Builds a complex object using simple objects | âœ|
| [Factory Method](/creational/factory.md) | Defers instantiation of an object to a specialized function for creating instances | âœ|
| [Object Pool](/creational/object-pool.md) | Instantiates and maintains a group of objects instances of the same type | âœ|
| [Singleton](/creational/singleton.md) | Restricts instantiation of a type to one object | âœ|

## Structural Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Bridge](/structural/bridge/main.go) | Decouples an interface from its implementation so that the two can vary independently | âœ|
| [Composite](/structural/composite/main.go) | Encapsulates and provides access to a number of different objects | âœ|
| [Decorator](/structural/decorator.md) | Adds behavior to an object, statically or dynamically | âœ|
| [Facade](/structural/facade/main.go) | Uses one type as an API to a number of others | âœ|
| [Flyweight](/structural/flyweight/main.go) | Reuses existing instances of objects with similar/identical state to minimize resource usage | âœ|
| [Proxy](/structural/proxy.md) | Provides a surrogate for an object to control it's actions | âœ|

## Behavioral Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Chain of Responsibility](/behavioral/chain_of_responsibility/main.go) | Avoids coupling a sender to receiver by giving more than object a chance to handle the request | âœ|
| [Command](/behavioral/command/main.go) | Bundles a command and arguments to call later | âœ|
| [Mediator](/behavioral/mediator/main.go) | Connects objects and acts as a proxy | âœ|
| [Memento](/behavioral/memento/main.go) | Generate an opaque token that can be used to go back to a previous state | âœ|
| [Observer](/behavioral/observer.md) | Provide a callback for notification of events/changes to data | âœ|
| [Registry](/behavioral/registry.md) | Keep track of all subclasses of a given class | âœ|
| [State](/behavioral/state/main.go) | Encapsulates varying behavior for the same object based on its internal state | âœ|
| [Strategy](/behavioral/strategy.md) | Enables an algorithm's behavior to be selected at runtime | âœ|
| [Template](/behavioral/template/main.go) | Defines a skeleton class which defers some methods to subclasses | âœ|
| [Visitor](/behavioral/visitor/main.go) | Separates an algorithm from an object on which it operates | âœ|

## Synchronization Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Condition Variable](/synchronization/condition_variable.md) | Provides a mechanism for threads to temporarily give up access in order to wait for some condition | âœ|
| [Lock/Mutex](/synchronization/mutex.md) | Enforces mutual exclusion limit on a resource to gain exclusive access | âœ|
| [Monitor](/synchronization/monitor.md) | Combination of mutex and condition variable patterns | âœ|
| [Read-Write Lock](/synchronization/read_write_lock.md) | Allows parallel read access, but only exclusive access on write operations to a resource | âœ|
| [Semaphore](/synchronization/semaphore.md) | Allows controlling access to a common resource | âœ|

## Concurrency Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [N-Barrier](/concurrency/n_barrier/main.go) | Prevents a process from proceeding until all N processes reach to the barrier | âœ|
| [Bounded Parallelism](/concurrency/bounded/bounded_parallelism.md) | Completes large number of independent tasks with resource limits | âœ|
| [Broadcast](/concurrency/broadcast.md) | Transfers a message to all recipients simultaneously | âœ|
| [Coroutines](/concurrency/coroutine.md) | Subroutines that allow suspending and resuming execution at certain locations | âœ|
| [Generators](/concurrency/generator.md) | Yields a sequence of values one at a time | âœ|
| [Reactor](/concurrency/reactor.md) | Demultiplexes service requests delivered concurrently to a service handler and dispatches them syncronously to the associated request handlers | âœ|
| [Parallelism](/concurrency/parallelism.md) | Completes large number of independent tasks | âœ|
| [Producer Consumer](/concurrency/producer_consumer.md) | Separates tasks from task executions | âœ|

## Messaging Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Fan-In](/messaging/fan_in.md) | Funnels tasks to a work sink (e.g. server) | âœ|
| [Fan-Out](/messaging/fan_out.md) | Distributes tasks among workers (e.g. producer) | âœ|
| [Futures & Promises](/messaging/futures_promises.md) | Acts as a place-holder of a result that is initially unknown for synchronization purposes | âœ|
| [Publish/Subscribe](/messaging/publish_subscribe.md) | Passes information to a collection of recipients who subscribed to a topic | âœ|
| [Push & Pull](/messaging/push_pull.md) | Distributes messages to multiple workers, arranged in a pipeline | âœ|

## Stability Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Bulkheads](/stability/bulkhead.md)  | Enforces a principle of failure containment (i.e. prevents cascading failures) | âœ|
| [Circuit-Breaker](/stability/circuit-breaker.md) | Stops the flow of the requests when requests are likely to fail | âœ|
| [Deadline](/stability/deadline.md) | Allows clients to stop waiting for a response once the probability of response becomes low (e.g. after waiting 10 seconds for a page refresh) | âœ|
| [Fail-Fast](/stability/fail_fast.md) | Checks the availability of required resources at the start of a request and fails if the requirements are not satisfied | âœ|
| [Handshaking](/stability/handshaking.md) | Asks a component if it can take any more load, if it can't, the request is declined | âœ|
| [Steady-State](/stability/steady_state.md) | For every service that accumulates a resource, some other service must recycle that resource | âœ|

## Profiling Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Timing Functions](/profiling/timing.md) | Wraps a function and logs the execution | âœ|

## Idioms

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Functional Options](/idiom/functional-options.md) | Allows creating clean APIs with sane defaults and idiomatic overrides | âœ|

## Anti-Patterns

| Pattern | Description | Status |
|:-------:|:----------- |:------:|
| [Cascading Failures](/anti-patterns/cascading_failures.md) | A failure in a system of interconnected parts in which the failure of a part causes a domino effect | âœ|
