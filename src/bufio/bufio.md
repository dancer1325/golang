* `type Reader struct {}`
    * == Reader / implements buffering
      * Reason:🧠it implements `Read(p []byte) (n int, err error)`🧠
    * `func (b *Reader) ReadString(delim byte) (string, error) {}`
        * 's return
            * `string`
                * data up to & including the `delim`
                    * if it encounters an error BEFORE -> returns the data read BEFORE & error itself
            * `error`
                * `err != nil` -- if and only if -- the returned data does NOT end | `delim`
        * recommendations
            * | simple uses,
                * use `Scanner`

* `func NewReader(rd io.Reader) *Reader {}`
    * 's return
        * `*Reader` / default size

* TODO:

* TODO:

* `type Writer struct {}`
    * == Writer /
        * implements
            * buffering -- for an -- io.Writer object
    * if an error occurs | writing to a `Writer` ->
        * NO MORE data will be accepted
        * ALL subsequent writes & `Writer.Flush` return the error
    * AFTER writing ALL data -> the client should call the `Writer.Flush()`
        * Reason:🧠guarantee ALL data has been forwarded -- to the -- underlying `io.Writer`🧠