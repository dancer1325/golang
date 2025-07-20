* `var logger atomic.Value`
  * == CURRENT logger Interface
  * allows
    * | test startup's init goroutines,
      * avoid race conditions
      * limited visibility BETWEEN goroutines

* `type Interface interface {}`
  * required -- by -- test loggers
  * uses
    * `os` package
  * allows
    * MULTIPLE goroutines call them SIMULTANEOUSLY
  * `func Getenv(name string) {}`
    * if Logger is set -> call `Logger().Getenv`
  * `func Stat(name string) {}`
    * if Logger is set -> call `Stat(file string)`
  * `func Open(name string) {}`
    * if Logger is set -> call `Logger().Open`
  * `Chdir(dir string)`

* `func Logger() Interface {}`
  * 👀returns the CURRENT test logger implementation👀
    * if there is NO logger -> returns `nil`
