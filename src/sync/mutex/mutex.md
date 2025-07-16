* `func throw(string)` & `func fatal(string)`
  * provided -- by -- runtime (TODO: ❓)

* `type Mutex struct {…}`
  * == 👀mutual exclusion lock👀
    * == ⚠️1! goroutine can access -- to -- shared resource / EACH time⚠️ 
    * -- via -- `.Lock()` & somethingDone & `.Unlock()`
  * 's 0 value
    * == unlocked mutex
  * AFTER FIRST use
    * ❌MUST NOT be copied by value❌
  * 👀n'th call to `Mutex.Unlock` synchronizes BEFORE m'th call to `Mutex.Lock`, for any `n < m`👀
    * requirement
      * ⚠️guarantee goroutine execution order⚠️
  * `func (m *Mutex) Lock() {}`
    * 👀locks mutex👀
      * if the lock is ALREADY in use -> calling goroutine blocks UNTIL the mutex is AVAILABLE
    * ⚠️locked `Mutex` is NOT associated -- with a -- particular goroutine⚠️
      * == 1 goroutine can lock a Mutex & another goroutine can unlock it  
  * `func (m *Mutex) Unlock() {}`
    * unlocks mutex
    * if mutex is NOT locked & you call it -> run-time error / 
      * ⚠️IMPOSSIBLE to cath⚠️
  * `func (m *Mutex) TryLock() bool {}`
    * tries to lock `m` & reports whether it succeeded
      * if it was ALREADY locked -> returns `false`
    * rarely used

* TODO:

