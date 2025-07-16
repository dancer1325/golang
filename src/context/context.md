* Package context == `Context` type

* `Context` type
  * == deadlines + cancellation signals + OTHER request-scoped values -- TODO: Check in examples --
  * uses
    * API boundaries / SAME process
    * BETWEEN processes 
  * use cases
    * incoming requests -- to a -- server, -> create a `Context`
    * outgoing calls -- to -- servers, -> accept a `Context`
  * | chain of function calls,
    * 👀`Context` is propagated👀 / OPTIONALLY replaced it -- with a -- derived `Context` /
      * created -- via -- `WithCancel`, `WithDeadline`, `WithTimeout`, or `WithValue`
  * ⚠️if a `Context` is canceled -> ALL Contexts / derived from it, ALSO canceled⚠️
  * `Deadline() (deadline time.Time, ok bool)`
    * returns the time | context's work should be canceled
    * if NO deadline is set -> returns `ok==false` 
    * if you call to `Deadline()` successively -> return the SAME results
  * `Done() <-chan struct{}`
    * returns a channel / 
      * if context's work is done -> channel is closed
        * close of the channel could happen asynchronously 
    * if this context can NEVER be canceled -> may return `nil` 
    * if you call to `Done()` successively -> return the SAME results
    * TODO: from "WithCancel arranges"

* 
  ```
  func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {}
  func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {}
  func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {}
  ```
  * `parent Context`
    * == argument
  * `(ctx Context, cancel CancelFunc)`
    * return arguments
    * derived context == child context

* programs / use Context -> ⚠️should follow the rules⚠️
  * Reason:🧠
    * interfaces are consistent | ACROSS packages
    * enable static analysis tools -- to -- check context propagation🧠
  * | `struct`, 
    * ❌NOT store Context❌
    * pass explicitly a `Context` / EACH function / need it
  * `Context` should be the FIRST parameter
  * ALTHOUGH the function permits it, NOT pass `nil` `Context`
  * if you do NOT know the Context to use -> pass `context.TODO`
  * context Values
    * uses for
      * request-scoped data / transits processes
      * APIs
    * NOT use for 
      * passing OPTIONAL parameters -- to -- functions

* 1 Context can be passed -- to -- MULTIPLE functions / run | DIFFERENT goroutines
  * Reason: 🧠Contexts are safe🧠

* `type CancelFunc func()`
  * | operation,
    * abandon its work / ❌NOT wait end it up❌
  * uses
    * by MULTIPLE goroutines SIMULTANEOUSLY
      * AFTER FIRST call -> subsequent calls -- to -- `CancelFunc` do NOTHING

* `func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc) {}`
  * TODO:

* `type CancelCauseFunc func(cause error)`
  * TODO:

* `func Cause(c Context) error {}`
  * | canceled context OR ANY children,
    * if you call -> retrieves the cause
      * & NO cause specified -> returns `ctx.Err()`

* `func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {}`
  * if you cancel this context -> releases resources / associated with it
    * recommendations
      * 👀if the operations / run | this Context, complete -> code should call cancel👀 
        // TODO:
        //	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
        //		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
        //		defer cancel()  // releases resources if slowOperation completes before timeout elapses
        //		return slowOperation(ctx)
        //	}

* `func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {}`
  * returns 
    * `Context`
      * == copy of the parent context / deadline BEFORE OR EQUAL `d`
      * `Context.Done` channel is closed |
        * deadline expires OR
        * returned `CancelFunc` is called OR
        * parent context's Done channel is closed 
  * if the parent's deadline is EARLIER than d -> `WithDeadline(parent, d)` semantically equivalent -- to -- parent
  * if you cancel this context -> releases resources / associated with it
    * | operations / run | this `Context`, complete -> code should call cancel

* `func Background() Context {}`
  * returns
    * `Context`
      * non-nil
      * empty 
        * == NON values
      * NEVER canceled
      * has NO deadline
  * uses
    * main function,
    * initialization,
    * tests,
    * top-level Context -- for -- incoming requests

* TODO: