* Package os
  * 's design
    * Unix-like + Go-like | error handling
      * Go-like | error handling == if call fails -> return values / typer error (!= error numbers | Unix)
  * provides
    * interface -- to -- OS functionality
      * platform-independent == uniform
  * features NOT generally AVAILABLE, appear | package syscall