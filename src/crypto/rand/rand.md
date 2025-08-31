* `var Reader io.Reader`
  * == random number generator instance
    * global
    * shared
    * cryptographically secure
  * ⚠️based on OS⚠️ 

* `func Read(b []byte) (n int, err error) {}`
  * == helper function / calls -- , via `io.ReadFull`, -- `Reader.Read`
  * 's return
    * `n == len(b)` -- if & ONLY if `err == nil`

* `func batched(f func([]byte) error, readMax int) func([]byte) error {}`
  * 's return
    * function /
      * populate `[]byte` -- by --
        * calling `f`
        * chunking | subslices / 's length of bytes <= `readMax` 
