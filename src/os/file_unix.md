* TODO:

* `type file struct {}`
  * == `*File`'s real representation
    * `*`
      * == extra level of indirection
      * == extra entity
      * Reason:🧠ensures that NO clients of os can overwrite this data🧠
        * ⚠️OTHERWISE, the finalizer could close the wrong file descriptor⚠️
  * TODO:

* TODO: