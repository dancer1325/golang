* Useful primitives

* `type Mutex struct {…}`
  * allows
    * 👁️checking that 1! goroutine can access to a variable / time 👁️ == mutual exclusion — via wrapping with — `Lock()` & `Unlock()`
  * `func (m *Mutex) Lock()`
  * `func (m *Mutex) Unlock()`
  * cd 'example/' & `go run example.go`