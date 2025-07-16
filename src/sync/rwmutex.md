* | runtime/rwmutex.go
  * == copy of it + changes

* `type RWMutex struct {}`
  * == reader/writer mutual exclusion lock
    * lock can be held -- by -- ARBITRARY number of readers OR 1 writer
  * 's zero value == unlocked mutex
  * AFTER FIRST use
    * ❌MUST NOT be copied by value❌
  * if lock hold by >=1 readers & some goroutine calls `RWMutex.Lock` -> concurrent calls to `RWMutex.RLock`, block 
    * | TILL writer has acquired & released the lock
      * == forbid recursive read-locking
      * Reason:🧠ensure lock becomes AVAILABLE -- to the -- writer🧠
  * `func (rw *RWMutex) Lock() {}`
    * locks rw -- for -- writing
      * ⚠️OTHERWISE, race condition⚠️
    * if the lock is ALREADY locked for reading OR writing -> blocks TILL lock is AVAILABLE
  * `func (rw *RWMutex) RLock() {}`
    * locks rw -- for -- reading
      * == ALWAYS read consistent values
      * ⚠️OTHERWISE, race condition⚠️
    * NOT uses
      * recursive read locking
  * `func (rw *RWMutex) RUnlock() {}`
    * undoes 1 `RWMutex.RLock` call
    * ❌NOT affect OTHER SIMULTANEOUS readers❌
    * if rw is NOT locked for reading | RUnlock's entry -> run-time error

* TODO:
     