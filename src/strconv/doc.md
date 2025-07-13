* Package `strconv`
  * implements 
    * conversions 
      * -- to -- string
      * -- from -- string

# Numeric Conversions
* `Atoi(stringValue)`
  * string is converted -- to -- int
    * requirements of int
      * ⚠️decimal⚠️
* `Itoa(intValue)`
  * int is converted -- to -- string
    * requirements of int
      * ⚠️decimal⚠️

* TODO:
          //
          // [ParseBool], [ParseFloat], [ParseInt], and [ParseUint] convert strings to values:
          //
          //	b, err := strconv.ParseBool("true")
          //	f, err := strconv.ParseFloat("3.1415", 64)
          //	i, err := strconv.ParseInt("-42", 10, 64)
          //	u, err := strconv.ParseUint("42", 10, 64)
          //
          // The parse functions return the widest type (float64, int64, and uint64),
          // but if the size argument specifies a narrower width the result can be
          // converted to that narrower type without data loss:
          //
          //	s := "2147483647" // biggest int32
          //	i64, err := strconv.ParseInt(s, 10, 32)
          //	...
          //	i := int32(i64)
          //
          // [FormatBool], [FormatFloat], [FormatInt], and [FormatUint] convert values to strings:
          //
          //	s := strconv.FormatBool(true)
          //	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
          //	s := strconv.FormatInt(-42, 16)
          //	s := strconv.FormatUint(42, 16)
          //
          // [AppendBool], [AppendFloat], [AppendInt], and [AppendUint] are similar but
          // append the formatted value to a destination slice.

# String Conversions

* `Quote` & `QuoteToASCII`
  * convert 
    * 👀strings -- to -- quoted Go string literals (== "stringValue")👀

* `QuoteToASCII`
  * any non-ASCII Unicode is escaped -- with -- `\u`

* `QuoteRune` & `QuoteRuneToASCII`
  * convert
    * 👀runes -- to -- quoted Go rune literals (== "runeValue")👀 

* `Unquote` & `UnquoteChar`
  * convert
    * 👀quoted Go string & char literals -- to -- unquoted string & char👀
