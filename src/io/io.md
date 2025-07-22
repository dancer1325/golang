* TODO:

* `type Reader interface {}`
  * `Read(p []byte) (n int, err error)`
    * reads <= `p`'s `len(p)` bytes
      * if SOME data is AVAILABLE, BUT NOT `len(p)` bytes -> Read returns IMMEDIATELY what is available
        * IMMEDIATELY != waiting for MORE 
    * returns 
      * number of bytes read (`0 <= n <= len(p)`)
      * any error encountered
    * if `n < len(p)` -> `p` could be used -- as -- scratch space | call
    * if AFTER successfully reading `n > 0` bytes, it encounters an error OR EOF condition -> ways to handle -- by --
      * producer
        * | SAME call, returns number of bytes read + error
        * | 
          * SAME call, return the bytes read 
          * SUBSEQUENT call, return the error
      * callers
        * recommended steps
          * process `n > 0` bytes returned
          * consider `err`
    * if `len(p) == 0` -> Read should ALWAYS return `n == 0`
      * if some error condition is known -> may return a `non-nil` error  
    * 's implementations
      * recommendations
        * discourage -- from -- returning 0 byte count + `nil` error
          * EXCEPT TO `len(p) == 0`
        * ❌MUST NOT retain `p`❌
    * 's callers
      * recommendations
        * 👀treat a return of 0 & `nil` == NOTHING happened👀

* TODO:

* `func ReadFull(r Reader, buf []byte) (n int, err error) {}`
  * reads EXACTLY `r`'s `len(buf)` bytes & 👀store | `buf`👀
  * returns
    * `n int` == number of bytes copied
      * == n == len(buf) if and only if `err == nil`
    * `err error`
      * Reason:🧠fewer bytes were read🧠
        * if NO bytes were read -> error == EOF
        * if EOF happens, | AFTER reading some BUT NOT ALL bytes -> returns ErrUnexpectedEOF
  * if `r` returns an error & have read >= `len(buf)` bytes -> error is dropped

* `type Writer interface {}`
  * `Write(p []byte) (n int, err error)`
    * writes `p`'s `len(p)` bytes | underlying data stream
    * returns 
      * number of bytes written -- from -- `p` (0 <= n <= len(p)) &
      * error encountered / caused the write to stop early
    * MUST
      * if it returns `n < len(p)` -> return a non-nil error, 
      * ❌NOT 
        * modify the slice data
          * EVEN temporarily
        * retain `p`❌

* TODO: