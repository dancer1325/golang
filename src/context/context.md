* Package context == `Context` type

* `Context` type
  * ACROSS API boundaries & BETWEEN processes
    * carries
      * deadlines
      * cancellation signals
      * OTHER request-scoped values 
  * uses
    * incoming requests -- to a -- server, -> create a `Context`
    * outgoing calls -- to -- servers, -> accept a `Context`
  * | chain of function calls,
    * 👀`Context` is propagated👀 / OPTIONALLY replaced it -- with a -- derived `Context` /
      * created -- via -- `WithCancel`, `WithDeadline`, `WithTimeout`, or `WithValue`
  * ⚠️if a `Context` is canceled -> ALL Contexts / derived from it, ALSO canceled⚠️

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

* `type CancelFunc func()`
  * | operation,
    * abandon its work / ❌NOT wait end it up❌
  * uses
    * by MULTIPLE goroutines SIMULTANEOUSLY
      * AFTER FIRST call -> subsequent calls -- to -- `CancelFunc` do NOTHING

* TODO: