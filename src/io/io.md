* TODO:

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