# Introduction

* Go programming language's reference manual
  * != [Go1.18- programming language's reference manual](go1.17_spec.html)
    * WITHOUT generics

* Go
  * language's
    * goal
      * general-purpose
      * systems programming
    * characteristics
      * strongly typed
      * garbage-collected
      * explicit support -- for -- concurrent programming
  * 's programs
    * constructed -- from --packages

## Notation

* syntax
  * compact
  * simple to parse
  * specified -- following -- [Extended Backus-Naur Form (EBNF)](https://en.wikipedia.org/wiki/Wirth_syntax_notation)

    ```
    Syntax      = { Production } .
    Production  = production_name "=" [ Expression ] "." .
    Expression  = Term { "|" Term } .
    Term        = Factor { Factor } .
    Factor      = production_name | token [ "…" token ] | Group | Option | Repetition .
    Group       = "(" Expression ")" .
    Option      = "[" Expression "]" .
    Repetition  = "{" Expression "}" .
    ```
    * `Production`
      * == expressions / constructed -- from -- terms & operators / increasing precedence
        ```
        |           alternation
        ()          grouping
        []          option (0 or 1 times)
        {}          repetition (0 -- to -- n times)
        ```
      * lowercase production names
        * uses
          * identify lexical (terminal) tokens

* Non-terminals
  * CamelCase
* Lexical tokens
  * enclosed | 
    * double quotes `""`
    * back quotes ``

* `a … b`
  * == set of characters a -- to -- b / as alternatives

* `…`
  * horizontal ellipsis 
  * uses
    * informally denote various
      * enumerations OR
      * code snippets
  * != `...`
  * ❌NOT Go language's token❌

* TODO: 
<p>
A link of the form [<a href="#Language_versions">Go 1.xx</a>] indicates that a described
language feature (or some aspect of it) was changed or added with language version 1.xx and
thus requires at minimum that language version to build.
For details, see the <a href="#Language_versions">linked section</a>
in the <a href="#Appendix">appendix</a>.
</p>

# Source code representation

<p>
Source code is Unicode text encoded in
<a href="https://en.wikipedia.org/wiki/UTF-8">UTF-8</a>. The text is not
canonicalized, so a single accented code point is distinct from the
same character constructed from combining an accent and a letter;
those are treated as two code points.  For simplicity, this document
will use the unqualified term <i>character</i> to refer to a Unicode code point
in the source text.
</p>
<p>
Each code point is distinct; for instance, uppercase and lowercase letters
are different characters.
</p>
<p>
Implementation restriction: For compatibility with other tools, a
compiler may disallow the NUL character (U+0000) in the source text.
</p>
<p>
Implementation restriction: For compatibility with other tools, a
compiler may ignore a UTF-8-encoded byte order mark
(U+FEFF) if it is the first Unicode code point in the source text.
A byte order mark may be disallowed anywhere else in the source.
</p>

<h3 id="Characters">Characters</h3>

<p>
The following terms are used to denote specific Unicode character categories:
</p>
<pre class="ebnf">
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point categorized as "Letter" */ .
unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */ .
</pre>

<p>
In <a href="https://www.unicode.org/versions/Unicode8.0.0/">The Unicode Standard 8.0</a>,
Section 4.5 "General Category" defines a set of character categories.
Go treats all characters in any of the Letter categories Lu, Ll, Lt, Lm, or Lo
as Unicode letters, and those in the Number category Nd as Unicode digits.
</p>

<h3 id="Letters_and_digits">Letters and digits</h3>

<p>
The underscore character <code>_</code> (U+005F) is considered a lowercase letter.
</p>
<pre class="ebnf">
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
</pre>

<h2 id="Lexical_elements">Lexical elements</h2>

<h3 id="Comments">Comments</h3>

<p>
Comments serve as program documentation. There are two forms:
</p>

<ol>
<li>
<i>Line comments</i> start with the character sequence <code>//</code>
and stop at the end of the line.
</li>
<li>
<i>General comments</i> start with the character sequence <code>/*</code>
and stop with the first subsequent character sequence <code>*/</code>.
</li>
</ol>

<p>
A comment cannot start inside a <a href="#Rune_literals">rune</a> or
<a href="#String_literals">string literal</a>, or inside a comment.
A general comment containing no newlines acts like a space.
Any other comment acts like a newline.
</p>

<h3 id="Tokens">Tokens</h3>

<p>
Tokens form the vocabulary of the Go language.
There are four classes: <i>identifiers</i>, <i>keywords</i>, <i>operators
and punctuation</i>, and <i>literals</i>.  <i>White space</i>, formed from
spaces (U+0020), horizontal tabs (U+0009),
carriage returns (U+000D), and newlines (U+000A),
is ignored except as it separates tokens
that would otherwise combine into a single token. Also, a newline or end of file
may trigger the insertion of a <a href="#Semicolons">semicolon</a>.
While breaking the input into tokens,
the next token is the longest sequence of characters that form a
valid token.
</p>

<h3 id="Semicolons">Semicolons</h3>

<p>
The formal syntax uses semicolons <code>";"</code> as terminators in
a number of productions. Go programs may omit most of these semicolons
using the following two rules:
</p>

<ol>
<li>
When the input is broken into tokens, a semicolon is automatically inserted
into the token stream immediately after a line's final token if that token is
<ul>
	<li>an
	    <a href="#Identifiers">identifier</a>
	</li>

	<li>an
	    <a href="#Integer_literals">integer</a>,
	    <a href="#Floating-point_literals">floating-point</a>,
	    <a href="#Imaginary_literals">imaginary</a>,
	    <a href="#Rune_literals">rune</a>, or
	    <a href="#String_literals">string</a> literal
	</li>

	<li>one of the <a href="#Keywords">keywords</a>
	    <code>break</code>,
	    <code>continue</code>,
	    <code>fallthrough</code>, or
	    <code>return</code>
	</li>

	<li>one of the <a href="#Operators_and_punctuation">operators and punctuation</a>
	    <code>++</code>,
	    <code>--</code>,
	    <code>)</code>,
	    <code>]</code>, or
	    <code>}</code>
	</li>
</ul>
</li>

<li>
To allow complex statements to occupy a single line, a semicolon
may be omitted before a closing <code>")"</code> or <code>"}"</code>.
</li>
</ol>

<p>
To reflect idiomatic use, code examples in this document elide semicolons
using these rules.
</p>


<h3 id="Identifiers">Identifiers</h3>

<p>
Identifiers name program entities such as variables and types.
An identifier is a sequence of one or more letters and digits.
The first character in an identifier must be a letter.
</p>
<pre class="ebnf">
identifier = letter { letter | unicode_digit } .
</pre>
<pre>
a
_x9
ThisVariableIsExported
αβ
</pre>

<p>
Some identifiers are <a href="#Predeclared_identifiers">predeclared</a>.
</p>


## Keywords

* reserver keywords
  * == ❌NOT ALLOWED to use -- as -- identifiers❌

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## Operators and punctuation

<p>
The following character sequences represent <a href="#Operators">operators</a>
(including <a href="#Assignment_statements">assignment operators</a>) and punctuation
[<a href="#Go_1.18">Go 1.18</a>]:
</p>
<pre class="grammar">
+    &amp;     +=    &amp;=     &amp;&amp;    ==    !=    (    )
-    |     -=    |=     ||    &lt;     &lt;=    [    ]
*    ^     *=    ^=     &lt;-    &gt;     &gt;=    {    }
/    &lt;&lt;    /=    &lt;&lt;=    ++    =     :=    ,    ;
%    &gt;&gt;    %=    &gt;&gt;=    --    !     ...   .    :
     &amp;^          &amp;^=          ~
</pre>

## Integer literals

* integer literal
  * := sequence of digits / 
    * == integer constant
* 
  An optional prefix sets a non-decimal base: <code>0b</code> or <code>0B</code>
  for binary, <code>0</code>, <code>0o</code>, or <code>0O</code> for octal,
  and <code>0x</code> or <code>0X</code> for hexadecimal
  [<a href="#Go_1.13">Go 1.13</a>].
  A single <code>0</code> is considered a decimal zero.
  In hexadecimal literals, letters <code>a</code> through <code>f</code>
  and <code>A</code> through <code>F</code> represent values 10 through 15.

<p>
For readability, an underscore character <code>_</code> may appear after
a base prefix or between successive digits; such underscores do not change
the literal's value.
</p>
<pre class="ebnf">
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .

decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
binary_digits  = binary_digit { [ "_" ] binary_digit } .
octal_digits   = octal_digit { [ "_" ] octal_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
</pre>

<pre>
42
4_2
0600
0_600
0o600
0O600       // second character is capital letter 'O'
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // an identifier, not an integer literal
42_         // invalid: _ must separate successive digits
4__2        // invalid: only one _ at a time
0_xBadFace  // invalid: _ must separate successive digits
</pre>

## Floating-point literals

<p>
A floating-point literal is a decimal or hexadecimal representation of a
<a href="#Constants">floating-point constant</a>.
</p>

<p>
A decimal floating-point literal consists of an integer part (decimal digits),
a decimal point, a fractional part (decimal digits), and an exponent part
(<code>e</code> or <code>E</code> followed by an optional sign and decimal digits).
One of the integer part or the fractional part may be elided; one of the decimal point
or the exponent part may be elided.
An exponent value exp scales the mantissa (integer and fractional part) by 10<sup>exp</sup>.
</p>

<p>
A hexadecimal floating-point literal consists of a <code>0x</code> or <code>0X</code>
prefix, an integer part (hexadecimal digits), a radix point, a fractional part (hexadecimal digits),
and an exponent part (<code>p</code> or <code>P</code> followed by an optional sign and decimal digits).
One of the integer part or the fractional part may be elided; the radix point may be elided as well,
but the exponent part is required. (This syntax matches the one given in IEEE 754-2008 §5.12.3.)
An exponent value exp scales the mantissa (integer and fractional part) by 2<sup>exp</sup>
[<a href="#Go_1.13">Go 1.13</a>].
</p>

<p>
For readability, an underscore character <code>_</code> may appear after
a base prefix or between successive digits; such underscores do not change
the literal value.
</p>

<pre class="ebnf">
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
</pre>

<pre>
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0

0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (integer subtraction)

0x.p1        // invalid: mantissa has no digits
1p-2         // invalid: p exponent requires hexadecimal mantissa
0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent
1_.5         // invalid: _ must separate successive digits
1._5         // invalid: _ must separate successive digits
1.5_e1       // invalid: _ must separate successive digits
1.5e_1       // invalid: _ must separate successive digits
1.5e1_       // invalid: _ must separate successive digits
</pre>


<h3 id="Imaginary_literals">Imaginary literals</h3>

<p>
An imaginary literal represents the imaginary part of a
<a href="#Constants">complex constant</a>.
It consists of an <a href="#Integer_literals">integer</a> or
<a href="#Floating-point_literals">floating-point</a> literal
followed by the lowercase letter <code>i</code>.
The value of an imaginary literal is the value of the respective
integer or floating-point literal multiplied by the imaginary unit <i>i</i>
[<a href="#Go_1.13">Go 1.13</a>]
</p>

<pre class="ebnf">
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
</pre>

<p>
For backward compatibility, an imaginary literal's integer part consisting
entirely of decimal digits (and possibly underscores) is considered a decimal
integer, even if it starts with a leading <code>0</code>.
</p>

<pre>
0i
0123i         // == 123i for backward-compatibility
0o123i        // == 0o123 * 1i == 83i
0xabci        // == 0xabc * 1i == 2748i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
0x1p-2i       // == 0x1p-2 * 1i == 0.25i
</pre>


<h3 id="Rune_literals">Rune literals</h3>

<p>
A rune literal represents a <a href="#Constants">rune constant</a>,
an integer value identifying a Unicode code point.
A rune literal is expressed as one or more characters enclosed in single quotes,
as in <code>'x'</code> or <code>'\n'</code>.
Within the quotes, any character may appear except newline and unescaped single
quote. A single quoted character represents the Unicode value
of the character itself,
while multi-character sequences beginning with a backslash encode
values in various formats.
</p>

<p>
The simplest form represents the single character within the quotes;
since Go source text is Unicode characters encoded in UTF-8, multiple
UTF-8-encoded bytes may represent a single integer value.  For
instance, the literal <code>'a'</code> holds a single byte representing
a literal <code>a</code>, Unicode U+0061, value <code>0x61</code>, while
<code>'ä'</code> holds two bytes (<code>0xc3</code> <code>0xa4</code>) representing
a literal <code>a</code>-dieresis, U+00E4, value <code>0xe4</code>.
</p>

<p>
Several backslash escapes allow arbitrary values to be encoded as
ASCII text.  There are four ways to represent the integer value
as a numeric constant: <code>\x</code> followed by exactly two hexadecimal
digits; <code>\u</code> followed by exactly four hexadecimal digits;
<code>\U</code> followed by exactly eight hexadecimal digits, and a
plain backslash <code>\</code> followed by exactly three octal digits.
In each case the value of the literal is the value represented by
the digits in the corresponding base.
</p>

<p>
Although these representations all result in an integer, they have
different valid ranges.  Octal escapes must represent a value between
0 and 255 inclusive.  Hexadecimal escapes satisfy this condition
by construction. The escapes <code>\u</code> and <code>\U</code>
represent Unicode code points so within them some values are illegal,
in particular those above <code>0x10FFFF</code> and surrogate halves.
</p>

<p>
After a backslash, certain single-character escapes represent special values:
</p>

<pre class="grammar">
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000B vertical tab
\\   U+005C backslash
\'   U+0027 single quote  (valid escape only within rune literals)
\"   U+0022 double quote  (valid escape only within string literals)
</pre>

<p>
An unrecognized character following a backslash in a rune literal is illegal.
</p>

<pre class="ebnf">
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
</pre>

<pre>
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // rune literal containing single quote character
'aa'         // illegal: too many characters
'\k'         // illegal: k is not recognized after a backslash
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\400'       // illegal: octal value over 255
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
</pre>


## String literals

<p>
A string literal represents a <a href="#Constants">string constant</a>
obtained from concatenating a sequence of characters. There are two forms:
raw string literals and interpreted string literals.
</p>

<p>
Raw string literals are character sequences between back quotes, as in
<code>`foo`</code>.  Within the quotes, any character may appear except
back quote. The value of a raw string literal is the
string composed of the uninterpreted (implicitly UTF-8-encoded) characters
between the quotes;
in particular, backslashes have no special meaning and the string may
contain newlines.
Carriage return characters ('\r') inside raw string literals
are discarded from the raw string value.
</p>

<p>
Interpreted string literals are character sequences between double
quotes, as in <code>&quot;bar&quot;</code>.
Within the quotes, any character may appear except newline and unescaped double quote.
The text between the quotes forms the
value of the literal, with backslash escapes interpreted as they
are in <a href="#Rune_literals">rune literals</a> (except that <code>\'</code> is illegal and
<code>\"</code> is legal), with the same restrictions.
The three-digit octal (<code>\</code><i>nnn</i>)
and two-digit hexadecimal (<code>\x</code><i>nn</i>) escapes represent individual
<i>bytes</i> of the resulting string; all other escapes represent
the (possibly multi-byte) UTF-8 encoding of individual <i>characters</i>.
Thus inside a string literal <code>\377</code> and <code>\xFF</code> represent
a single byte of value <code>0xFF</code>=255, while <code>ÿ</code>,
<code>\u00FF</code>, <code>\U000000FF</code> and <code>\xc3\xbf</code> represent
the two bytes <code>0xc3</code> <code>0xbf</code> of the UTF-8 encoding of character
U+00FF.
</p>

```
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
```

<pre>
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
</pre>

<p>
These examples all represent the same string:
</p>

<pre>
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
</pre>

<p>
If the source code represents a character as two code points, such as
a combining form involving an accent and a letter, the result will be
an error if placed in a rune literal (it is not a single code
point), and will appear as two code points if placed in a string
literal.
</p>


# Constants

* types
  * boolean constants
  * numeric constants
    * rune constants
    * integer constants
    * floating-point constants
    * complex constants
  * string constants

* constant value
  * types
    * rune literal
    * integer literal
    * floating-point literal
    * imaginary literal
    * string literal
    * identifier / == a constant
    * constant expression
    * conversion / 's result == constant
    * built-in functions / 's result == constant
      * _Example:_ 
        * `min(someConstant)`
        * `max(someConstant)`
        * `unsafe.Sizeof(certainValues)`
        * applied to <a href="#Package_unsafe"></a>,
        <code>cap</code> or <code>len</code> applied to
        <a href="#Length_and_capacity">some expressions</a>,
        <code>real</code> and <code>imag</code> applied to a complex constant
        and <code>complex</code> applied to numeric constants.
    
    The boolean truth values are represented by the predeclared constants
    <code>true</code> and <code>false</code>. The predeclared identifier
    <a href="#Iota">iota</a> denotes an integer constant.

<p>
In general, complex constants are a form of
<a href="#Constant_expressions">constant expression</a>
and are discussed in that section.
</p>

<p>
Numeric constants represent exact values of arbitrary precision and do not overflow.
Consequently, there are no constants denoting the IEEE-754 negative zero, infinity,
and not-a-number values.
</p>

<p>
Constants may be <a href="#Types">typed</a> or <i>untyped</i>.
Literal constants, <code>true</code>, <code>false</code>, <code>iota</code>,
and certain <a href="#Constant_expressions">constant expressions</a>
containing only untyped constant operands are untyped.
</p>

<p>
A constant may be given a type explicitly by a <a href="#Constant_declarations">constant declaration</a>
or <a href="#Conversions">conversion</a>, or implicitly when used in a
<a href="#Variable_declarations">variable declaration</a> or an
<a href="#Assignment_statements">assignment statement</a> or as an
operand in an <a href="#Expressions">expression</a>.
It is an error if the constant value
cannot be <a href="#Representability">represented</a> as a value of the respective type.
If the type is a type parameter, the constant is converted into a non-constant
value of the type parameter.
</p>

<p>
An untyped constant has a <i>default type</i> which is the type to which the
constant is implicitly converted in contexts where a typed value is required,
for instance, in a <a href="#Short_variable_declarations">short variable declaration</a>
such as <code>i := 0</code> where there is no explicit type.
The default type of an untyped constant is <code>bool</code>, <code>rune</code>,
<code>int</code>, <code>float64</code>, <code>complex128</code>, or <code>string</code>
respectively, depending on whether it is a boolean, rune, integer, floating-point,
complex, or string constant.
</p>

<p>
Implementation restriction: Although numeric constants have arbitrary
precision in the language, a compiler may implement them using an
internal representation with limited precision.  That said, every
implementation must:
</p>

<ul>
	<li>Represent integer constants with at least 256 bits.</li>

	<li>Represent floating-point constants, including the parts of
	    a complex constant, with a mantissa of at least 256 bits
	    and a signed binary exponent of at least 16 bits.</li>

	<li>Give an error if unable to represent an integer constant
	    precisely.</li>

	<li>Give an error if unable to represent a floating-point or
	    complex constant due to overflow.</li>

	<li>Round to the nearest representable constant if unable to
	    represent a floating-point or complex constant due to limits
	    on precision.</li>
</ul>

<p>
These requirements apply both to literal constants and to the result
of evaluating <a href="#Constant_expressions">constant
expressions</a>.
</p>


# Variables

* == storage location
* uses
  * hold a value /
    * ⚠️ALLOWED values is determined -- by the -- variable's [type](#variables)⚠️

* entities / reserve storage -- for a -- named variable
  * [variable declaration](#variable-declarations)
  * [function declaration](#function-declarations)'s
    * parameters
    * results
  * [function literal](#function-literals) 

Calling the built-in function <a href="#Allocation"><code>new</code></a>
or taking the address of a <a href="#Composite_literals">composite literal</a>
allocates storage for a variable at run time.
Such an anonymous variable is referred to via a (possibly implicit)
<a href="#Address_operators">pointer indirection</a>.


<p>
<i>Structured</i> variables of <a href="#Array_types">array</a>, <a href="#Slice_types">slice</a>,
and <a href="#Struct_types">struct</a> types have elements and fields that may
be <a href="#Address_operators">addressed</a> individually. Each such element
acts like a variable.
</p>

<p>
The <i>static type</i> (or just <i>type</i>) of a variable is the
type given in its declaration, the type provided in the
<code>new</code> call or composite literal, or the type of
an element of a structured variable.
Variables of interface type also have a distinct <i>dynamic type</i>,
which is the (non-interface) type of the value assigned to the variable at run time
(unless the value is the predeclared identifier <code>nil</code>,
which has no type).
The dynamic type may vary during execution but values stored in interface
variables are always <a href="#Assignability">assignable</a>
to the static type of the variable.
</p>

<pre>
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
</pre>

<p>
A variable's value is retrieved by referring to the variable in an
<a href="#Expressions">expression</a>; it is the most recent value
<a href="#Assignment_statements">assigned</a> to the variable.
If a variable has not yet been assigned a value, its value is the
<a href="#The_zero_value">zero value</a> for its type.
</p>


# Types

* type
  * == 👀set of values together + operations + values' methods👀

```
Type      = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
  # if the type is generic -> specify [ TypeArgs ]
  # TypeLit == type -- from -- EXISTING types
TypeName  = identifier | QualifiedIdent .
TypeArgs  = "[" TypeList [ "," ] "]" .
TypeList  = Type { "," Type } .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
            SliceType | MapType | ChannelType .
```

<p>
The language <a href="#Predeclared_identifiers">predeclares</a> certain type names.
Others are introduced with <a href="#Type_declarations">type declarations</a>
or <a href="#Type_parameter_declarations">type parameter lists</a>.
<i>Composite types</i>&mdash;array, struct, pointer, function,
interface, slice, map, and channel types&mdash;may be constructed using
type literals.
</p>

<p>
Predeclared types, defined types, and type parameters are called <i>named types</i>.
An alias denotes a named type if the type given in the alias declaration is a named type.
</p>

## Boolean types

<p>
A <i>boolean type</i> represents the set of Boolean truth values
denoted by the predeclared constants <code>true</code>
and <code>false</code>. The predeclared boolean type is <code>bool</code>;
it is a <a href="#Type_definitions">defined type</a>.
</p>

<h3 id="Numeric_types">Numeric types</h3>

<p>
An <i>integer</i>, <i>floating-point</i>, or <i>complex</i> type
represents the set of integer, floating-point, or complex values, respectively.
They are collectively called <i>numeric types</i>.
The predeclared architecture-independent numeric types are:
</p>

<pre class="grammar">
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
</pre>

<p>
The value of an <i>n</i>-bit integer is <i>n</i> bits wide and represented using
<a href="https://en.wikipedia.org/wiki/Two's_complement">two's complement arithmetic</a>.
</p>

<p>
There is also a set of predeclared integer types with implementation-specific sizes:
</p>

<pre class="grammar">
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
</pre>

<p>
To avoid portability issues all numeric types are <a href="#Type_definitions">defined
types</a> and thus distinct except
<code>byte</code>, which is an <a href="#Alias_declarations">alias</a> for <code>uint8</code>, and
<code>rune</code>, which is an alias for <code>int32</code>.
Explicit conversions
are required when different numeric types are mixed in an expression
or assignment. For instance, <code>int32</code> and <code>int</code>
are not the same type even though they may have the same size on a
particular architecture.
</p>

## String types

<p>
A <i>string type</i> represents the set of string values.
A string value is a (possibly empty) sequence of bytes.
The number of bytes is called the length of the string and is never negative.
Strings are immutable: once created,
it is impossible to change the contents of a string.
The predeclared string type is <code>string</code>;
it is a <a href="#Type_definitions">defined type</a>.
</p>

<p>
The length of a string <code>s</code> can be discovered using
the built-in function <a href="#Length_and_capacity"><code>len</code></a>.
The length is a compile-time constant if the string is a constant.
A string's bytes can be accessed by integer <a href="#Index_expressions">indices</a>
0 through <code>len(s)-1</code>.
It is illegal to take the address of such an element; if
<code>s[i]</code> is the <code>i</code>'th byte of a
string, <code>&amp;s[i]</code> is invalid.
</p>


## Array types

<p>
An array is a numbered sequence of elements of a single
type, called the element type.
The number of elements is called the length of the array and is never negative.
</p>

<pre class="ebnf">
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
</pre>

<p>
The length is part of the array's type; it must evaluate to a
non-negative <a href="#Constants">constant</a>
<a href="#Representability">representable</a> by a value
of type <code>int</code>.
The length of array <code>a</code> can be discovered
using the built-in function <a href="#Length_and_capacity"><code>len</code></a>.
The elements can be addressed by integer <a href="#Index_expressions">indices</a>
0 through <code>len(a)-1</code>.
Array types are always one-dimensional but may be composed to form
multi-dimensional types.
</p>

<pre>
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
</pre>

<p>
An array type <code>T</code> may not have an element of type <code>T</code>,
or of a type containing <code>T</code> as a component, directly or indirectly,
if those containing types are only array or struct types.
</p>

<pre>
// invalid array types
type (
	T1 [10]T1                 // element type of T1 is T1
	T2 [10]struct{ f T2 }     // T2 contains T2 as component of a struct
	T3 [10]T4                 // T3 contains T3 as component of a struct in T4
	T4 struct{ f T3 }         // T4 contains T4 as component of array T3 in a struct
)

// valid array types
type (
	T5 [10]*T5                // T5 contains T5 as component of a pointer
	T6 [10]func() T6          // T6 contains T6 as component of a function type
	T7 [10]struct{ f []T7 }   // T7 contains T7 as component of a slice in a struct
)
</pre>

## Slice types

<p>
A slice is a descriptor for a contiguous segment of an <i>underlying array</i> and
provides access to a numbered sequence of elements from that array.
A slice type denotes the set of all slices of arrays of its element type.
The number of elements is called the length of the slice and is never negative.
The value of an uninitialized slice is <code>nil</code>.
</p>

<pre class="ebnf">
SliceType = "[" "]" ElementType .
</pre>

<p>
The length of a slice <code>s</code> can be discovered by the built-in function
<a href="#Length_and_capacity"><code>len</code></a>; unlike with arrays it may change during
execution.  The elements can be addressed by integer <a href="#Index_expressions">indices</a>
0 through <code>len(s)-1</code>.  The slice index of a
given element may be less than the index of the same element in the
underlying array.
</p>
<p>
A slice, once initialized, is always associated with an underlying
array that holds its elements.  A slice therefore shares storage
with its array and with other slices of the same array; by contrast,
distinct arrays always represent distinct storage.
</p>
<p>
The array underlying a slice may extend past the end of the slice.
The <i>capacity</i> is a measure of that extent: it is the sum of
the length of the slice and the length of the array beyond the slice;
a slice of length up to that capacity can be created by
<a href="#Slice_expressions"><i>slicing</i></a> a new one from the original slice.
The capacity of a slice <code>a</code> can be discovered using the
built-in function <a href="#Length_and_capacity"><code>cap(a)</code></a>.
</p>

<p>
A new, initialized slice value for a given element type <code>T</code> may be
made using the built-in function
<a href="#Making_slices_maps_and_channels"><code>make</code></a>,
which takes a slice type
and parameters specifying the length and optionally the capacity.
A slice created with <code>make</code> always allocates a new, hidden array
to which the returned slice value refers. That is, executing
</p>

<pre>
make([]T, length, capacity)
</pre>

<p>
produces the same slice as allocating an array and <a href="#Slice_expressions">slicing</a>
it, so these two expressions are equivalent:
</p>

<pre>
make([]int, 50, 100)
new([100]int)[0:50]
</pre>

<p>
Like arrays, slices are always one-dimensional but may be composed to construct
higher-dimensional objects.
With arrays of arrays, the inner arrays are, by construction, always the same length;
however with slices of slices (or arrays of slices), the inner lengths may vary dynamically.
Moreover, the inner slices must be initialized individually.
</p>

## Struct types

* struct
  * == sequence of named elements == sequence of fields /
    * EACH field == name + type
    * ⚠️NON-[blank](#blank-identifier----_---) field names MUST be [unique](#uniqueness-of-identifiers) | EACH struct⚠️

```go
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
	// IdentifierList  ==  explicitly    OR  EmbeddedField  ==  implicitly 
EmbeddedField = [ "*" ] TypeName [ TypeArgs ] .
Tag           = string_lit .
```

* embedded field
  * == field / type declared & ❌NO explicit field name❌
    * unqualified type name == field name
  * ways to specify it
    * type name OR 
    * pointer -- to a -- NON-interface type name

* promoted field
  * := struct's (`x`) embedded field's field OR method (`f`) / `x.f` == selector of that field OR method

* TODO:
<p>
Promoted fields act like ordinary fields
of a struct except that they cannot be used as field names in
<a href="#Composite_literals">composite literals</a> of the struct.
</p>

<p>
Given a struct type <code>S</code> and a <a href="#Types">named type</a>
<code>T</code>, promoted methods are included in the method set of the struct as follows:
</p>
<ul>
	<li>
	If <code>S</code> contains an embedded field <code>T</code>,
	the <a href="#Method_sets">method sets</a> of <code>S</code>
	and <code>*S</code> both include promoted methods with receiver
	<code>T</code>. The method set of <code>*S</code> also
	includes promoted methods with receiver <code>*T</code>.
	</li>

	<li>
	If <code>S</code> contains an embedded field <code>*T</code>,
	the method sets of <code>S</code> and <code>*S</code> both
	include promoted methods with receiver <code>T</code> or
	<code>*T</code>.
	</li>
</ul>

* `Tag           = string_lit .`
  * OPTIONAL
  * == field declaration's ALL fields' attribute
  * if `""` == absent tag
  * if you want to make it visible -> through [reflection interface](https://pkg.go.dev/reflect#StructTag)
  * uses
    * structs' [type identity](#type-identity)

<p>
A struct type <code>T</code> may not contain a field of type <code>T</code>,
or of a type containing <code>T</code> as a component, directly or indirectly,
if those containing types are only array or struct types.
</p>

<pre>
// invalid struct types
type (
	T1 struct{ T1 }            // T1 contains a field of T1
	T2 struct{ f [10]T2 }      // T2 contains T2 as component of an array
	T3 struct{ T4 }            // T3 contains T3 as component of an array in struct T4
	T4 struct{ f [10]T3 }      // T4 contains T4 as component of struct T3 in an array
)

// valid struct types
type (
	T5 struct{ f *T5 }         // T5 contains T5 as component of a pointer
	T6 struct{ f func() T6 }   // T6 contains T6 as component of a function type
	T7 struct{ f [10][]T7 }    // T7 contains T7 as component of a slice in an array
)
</pre>

## Pointer types

* pointer type
  * == ⭐️ALL pointers -- to -- given variables of certain type⭐️
  * ❌if it's NOT initialized -> 's value == `nil`❌
  * syntax
    ```go
    PointerType = "*" BaseType .
    BaseType    = Type .
    ```
    * _Example:_
      ```go
      *Point
      *[4]int
      ```

* pointer's base type
  * := given type's variables

* _Examples:_ [here](examples/pointer-types.go)

## Function types

* := ALL functions / SAME parameter & result types
* function type variable uninitialized 's value == `nil`

```go
FunctionType   = "func" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```

<p>
Within a list of parameters or results, the names (IdentifierList)
must either all be present or all be absent. If present, each name
stands for one item (parameter or result) of the specified type and
all non-<a href="#Blank_identifier">blank</a> names in the signature
must be <a href="#Uniqueness_of_identifiers">unique</a>.
If absent, each type stands for one item of that type.
Parameter and result
lists are always parenthesized except that if there is exactly
one unnamed result it may be written as an unparenthesized type.
</p>

<p>
The final incoming parameter in a function signature may have
a type prefixed with <code>...</code>.
A function with such a parameter is called <i>variadic</i> and
may be invoked with zero or more arguments for that parameter.
</p>

```go
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

## Interface types

* == type set
* variable of interface type
  * can store a value / type included | type set
    * == ["that type implements the interface"](#implementing-an-interface)
  * if it's uninitialized -> == `nil`

```go
InterfaceType  = "interface" "{" { InterfaceElem ";" } "}" .
InterfaceElem  = MethodElem | TypeElem .
MethodElem     = MethodName Signature .
MethodName     = identifier .
TypeElem       = TypeTerm { "|" TypeTerm } .
TypeTerm       = Type | UnderlyingType .
UnderlyingType = "~" Type .
```

* == list of interface elements /
  * interface element is
    * method OR
    * type element
      * == UNION of >=1 type terms

### Basic interfaces

* basic interfaces
  * := interfaces / type sets are defined -- entirely by a -- list of methods 
  * MOST basic form
    * == EMPTY list of methods
  * 's type set
    * == set of types / implement ALL those methods

* [method set](#method-sets)
  * == methods / specified -- by the -- interface

<pre>
// A simple File interface.
interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}
</pre>

<p>
The name of each explicitly specified method must be <a href="#Uniqueness_of_identifiers">unique</a>
and not <a href="#Blank_identifier">blank</a>.
</p>

<pre>
interface {
	String() string
	String() string  // illegal: String not unique
	_(x int)         // illegal: method must have non-blank name
}
</pre>

<p>
More than one type may implement an interface.
For instance, if two types <code>S1</code> and <code>S2</code>
have the method set
</p>

<pre>
func (p T) Read(p []byte) (n int, err error)
func (p T) Write(p []byte) (n int, err error)
func (p T) Close() error
</pre>

<p>
(where <code>T</code> stands for either <code>S1</code> or <code>S2</code>)
then the <code>File</code> interface is implemented by both <code>S1</code> and
<code>S2</code>, regardless of what other methods
<code>S1</code> and <code>S2</code> may have or share.
</p>

* every type 
  * / (TODO:) member of the type set of an interface implements that interface.
  * may implement SEVERAL DISTINT interfaces

* empty interface -- `interface {}` --
  * == set of ALL (NON-interface) types 
  * implemented -- by -- ALL types
  * `any`
    * 💡alias of it💡
    * [Go 1.18](#go-118)

<p>
Similarly, consider this interface specification,
which appears within a <a href="#Type_declarations">type declaration</a>
to define an interface called <code>Locker</code>:
</p>

<pre>
type Locker interface {
	Lock()
	Unlock()
}
</pre>

<p>
If <code>S1</code> and <code>S2</code> also implement
</p>

<pre>
func (p T) Lock() { … }
func (p T) Unlock() { … }
</pre>

<p>
they implement the <code>Locker</code> interface as well
as the <code>File</code> interface.
</p>

### Embedded interfaces

<p>
In a slightly more general form
an interface <code>T</code> may use a (possibly qualified) interface type
name <code>E</code> as an interface element. This is called
<i>embedding</i> interface <code>E</code> in <code>T</code>
[<a href="#Go_1.14">Go 1.14</a>].
The type set of <code>T</code> is the <i>intersection</i> of the type sets
defined by <code>T</code>'s explicitly declared methods and the type sets
of <code>T</code>’s embedded interfaces.
In other words, the type set of <code>T</code> is the set of all types that implement all the
explicitly declared methods of <code>T</code> and also all the methods of
<code>E</code>
[<a href="#Go_1.18">Go 1.18</a>].
</p>

<pre>
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader  // includes methods of Reader in ReadWriter's method set
	Writer  // includes methods of Writer in ReadWriter's method set
}
</pre>

<p>
When embedding interfaces, methods with the
<a href="#Uniqueness_of_identifiers">same</a> names must
have <a href="#Type_identity">identical</a> signatures.
</p>

<pre>
type ReadCloser interface {
	Reader   // includes methods of Reader in ReadCloser's method set
	Close()  // illegal: signatures of Reader.Close and Close are different
}
</pre>

<h4 id="General_interfaces">General interfaces</h4>

<p>
In their most general form, an interface element may also be an arbitrary type term
<code>T</code>, or a term of the form <code>~T</code> specifying the underlying type <code>T</code>,
or a union of terms <code>t<sub>1</sub>|t<sub>2</sub>|…|t<sub>n</sub></code>
[<a href="#Go_1.18">Go 1.18</a>].
Together with method specifications, these elements enable the precise
definition of an interface's type set as follows:
</p>

<ul>
	<li>The type set of the empty interface is the set of all non-interface types.
	</li>

	<li>The type set of a non-empty interface is the intersection of the type sets
		of its interface elements.
	</li>

	<li>The type set of a method specification is the set of all non-interface types
		whose method sets include that method.
	</li>

	<li>The type set of a non-interface type term is the set consisting
		of just that type.
	</li>

	<li>The type set of a term of the form <code>~T</code>
		is the set of all types whose underlying type is <code>T</code>.
	</li>

	<li>The type set of a <i>union</i> of terms
		<code>t<sub>1</sub>|t<sub>2</sub>|…|t<sub>n</sub></code>
		is the union of the type sets of the terms.
	</li>
</ul>

<p>
The quantification "the set of all non-interface types" refers not just to all (non-interface)
types declared in the program at hand, but all possible types in all possible programs, and
hence is infinite.
Similarly, given the set of all non-interface types that implement a particular method, the
intersection of the method sets of those types will contain exactly that method, even if all
types in the program at hand always pair that method with another method.
</p>

<p>
By construction, an interface's type set never contains an interface type.
</p>

<pre>
// An interface representing only the type int.
interface {
	int
}

// An interface representing all types with underlying type int.
interface {
	~int
}

// An interface representing all types with underlying type int that implement the String method.
interface {
	~int
	String() string
}

// An interface representing an empty type set: there is no type that is both an int and a string.
interface {
	int
	string
}
</pre>

<p>
In a term of the form <code>~T</code>, the underlying type of <code>T</code>
must be itself, and <code>T</code> cannot be an interface.
</p>

<pre>
type MyInt int

interface {
	~[]byte  // the underlying type of []byte is itself
	~MyInt   // illegal: the underlying type of MyInt is not MyInt
	~error   // illegal: error is an interface
}
</pre>

<p>
Union elements denote unions of type sets:
</p>

<pre>
// The Float interface represents all floating-point types
// (including any named types whose underlying types are
// either float32 or float64).
type Float interface {
	~float32 | ~float64
}
</pre>

<p>
The type <code>T</code> in a term of the form <code>T</code> or <code>~T</code> cannot
be a <a href="#Type_parameter_declarations">type parameter</a>, and the type sets of all
non-interface terms must be pairwise disjoint (the pairwise intersection of the type sets must be empty).
Given a type parameter <code>P</code>:
</p>

<pre>
interface {
	P                // illegal: P is a type parameter
	int | ~P         // illegal: P is a type parameter
	~int | MyInt     // illegal: the type sets for ~int and MyInt are not disjoint (~int includes MyInt)
	float32 | Float  // overlapping type sets but Float is an interface
}
</pre>

<p>
Implementation restriction:
A union (with more than one term) cannot contain the
<a href="#Predeclared_identifiers">predeclared identifier</a> <code>comparable</code>
or interfaces that specify methods, or embed <code>comparable</code> or interfaces
that specify methods.
</p>

<p>
Interfaces that are not <a href="#Basic_interfaces">basic</a> may only be used as type
constraints, or as elements of other interfaces used as constraints.
They cannot be the types of values or variables, or components of other,
non-interface types.
</p>

<pre>
var x Float                     // illegal: Float is not a basic interface

var x interface{} = Float(nil)  // illegal

type Floatish struct {
	f Float                 // illegal
}
</pre>

<p>
An interface type <code>T</code> may not embed a type element
that is, contains, or embeds <code>T</code>, directly or indirectly.
</p>

<pre>
// illegal: Bad may not embed itself
type Bad interface {
	Bad
}

// illegal: Bad1 may not embed itself using Bad2
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}

// illegal: Bad3 may not embed a union containing Bad3
type Bad3 interface {
	~int | ~string | Bad3
}

// illegal: Bad4 may not embed an array containing Bad4 as element type
type Bad4 interface {
	[10]Bad4
}
</pre>

### Implementing an interface

* let's be type `T` & interface `I`
  * 👀requirements / `T` implements `I`👀
    * `T`
      * != interface & == element -- of the -- `I`'s type set, OR
      * == interface & `T`'s type set == subset of the `I`'s type set
  * 👀requirements / value of type `T` implements `I`👀
    * `T` implements the interface

## Map types

<p>
A map is an unordered group of elements of one type, called the
element type, indexed by a set of unique <i>keys</i> of another type,
called the key type.
The value of an uninitialized map is <code>nil</code>.
</p>

<pre class="ebnf">
MapType     = "map" "[" KeyType "]" ElementType .
KeyType     = Type .
</pre>

<p>
The <a href="#Comparison_operators">comparison operators</a>
<code>==</code> and <code>!=</code> must be fully defined
for operands of the key type; thus the key type must not be a function, map, or
slice.
If the key type is an interface type, these
comparison operators must be defined for the dynamic key values;
failure will cause a <a href="#Run_time_panics">run-time panic</a>.
</p>

<pre>
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
</pre>

<p>
The number of map elements is called its length.
For a map <code>m</code>, it can be discovered using the
built-in function <a href="#Length_and_capacity"><code>len</code></a>
and may change during execution. Elements may be added during execution
using <a href="#Assignment_statements">assignments</a> and retrieved with
<a href="#Index_expressions">index expressions</a>; they may be removed with the
<a href="#Deletion_of_map_elements"><code>delete</code></a> and
<a href="#Clear"><code>clear</code></a> built-in function.
</p>

<p>
A new, empty map value is made using the built-in
function <a href="#Making_slices_maps_and_channels"><code>make</code></a>,
which takes the map type and an optional capacity hint as arguments:
</p>

<pre>
make(map[string]int)
make(map[string]int, 100)
</pre>

<p>
The initial capacity does not bound its size:
maps grow to accommodate the number of items
stored in them, with the exception of <code>nil</code> maps.
A <code>nil</code> map is equivalent to an empty map except that no elements
may be added.
</p>

## Channel types

* channel
  * allows
    * goroutines can
      * send values / specific type
      * receive values / specific type

        <a href="#Send_statements">sending</a> and
        <a href="#Receive_operator">receiving</a>
        values of a specified element type.
        The value of an uninitialized channel is <code>nil</code>.

<pre class="ebnf">
ChannelType = ( "chan" | "chan" "&lt;-" | "&lt;-" "chan" ) ElementType .
</pre>

<p>
The optional <code>&lt;-</code> operator specifies the channel <i>direction</i>,
<i>send</i> or <i>receive</i>. If a direction is given, the channel is <i>directional</i>,
otherwise it is <i>bidirectional</i>.
A channel may be constrained only to send or only to receive by
<a href="#Assignment_statements">assignment</a> or
explicit <a href="#Conversions">conversion</a>.
</p>

<pre>
chan T          // can be used to send and receive values of type T
chan&lt;- float64  // can only be used to send float64s
&lt;-chan int      // can only be used to receive ints
</pre>

<p>
The <code>&lt;-</code> operator associates with the leftmost <code>chan</code>
possible:
</p>

<pre>
chan&lt;- chan int    // same as chan&lt;- (chan int)
chan&lt;- &lt;-chan int  // same as chan&lt;- (&lt;-chan int)
&lt;-chan &lt;-chan int  // same as &lt;-chan (&lt;-chan int)
chan (&lt;-chan int)
</pre>

<p>
A new, initialized channel
value can be made using the built-in function
<a href="#Making_slices_maps_and_channels"><code>make</code></a>,
which takes the channel type and an optional <i>capacity</i> as arguments:
</p>

<pre>
make(chan int, 100)
</pre>

<p>
The capacity, in number of elements, sets the size of the buffer in the channel.
If the capacity is zero or absent, the channel is unbuffered and communication
succeeds only when both a sender and receiver are ready. Otherwise, the channel
is buffered and communication succeeds without blocking if the buffer
is not full (sends) or not empty (receives).
A <code>nil</code> channel is never ready for communication.
</p>

<p>
A channel may be closed with the built-in function
<a href="#Close"><code>close</code></a>.
The multi-valued assignment form of the
<a href="#Receive_operator">receive operator</a>
reports whether a received value was sent before
the channel was closed.
</p>

<p>
A single channel may be used in
<a href="#Send_statements">send statements</a>,
<a href="#Receive_operator">receive operations</a>,
and calls to the built-in functions
<a href="#Length_and_capacity"><code>cap</code></a> and
<a href="#Length_and_capacity"><code>len</code></a>
by any number of goroutines without further synchronization.
Channels act as first-in-first-out queues.
For example, if one goroutine sends values on a channel
and a second goroutine receives them, the values are
received in the order sent.
</p>

<h2 id="Properties_of_types_and_values">Properties of types and values</h2>

## Underlying types

* == 💡real type / Go uses internally💡
* EXIST underlying type / EACH type
  * if type == 
    * built-in
      * boolean -> underlying type == boolean
      * numeric -> underlying type == numeric
      * string -> underlying type == string
    * custom type / -- based on --
      * built-in types -> underlying type == built-in types 
      * type literals -> underlying type == type literal

* TODO:
For a type parameter that is the underlying type of its
<a href="#Type_constraints">type constraint</a>, which is always an interface.


<pre>
func f[P any](x P) { … }
</pre>

<p>
The underlying type of <code>P</code> is <code>interface{}</code>.
</p>

## Core types

* 👀EXIST core type / EACH non-interface type👀
  * == underlying type
  
<a href="#Underlying_types">underlying type</a> of <code>T</code>.

<p>
An interface <code>T</code> has a core type if one of the following
conditions is satisfied:
</p>

<ol>
<li>
There is a single type <code>U</code> which is the <a href="#Underlying_types">underlying type</a>
of all types in the <a href="#Interface_types">type set</a> of <code>T</code>; or
</li>
<li>
the type set of <code>T</code> contains only <a href="#Channel_types">channel types</a>
with identical element type <code>E</code>, and all directional channels have the same
direction.
</li>
</ol>

<p>
No other interfaces have a core type.
</p>

<p>
The core type of an interface is, depending on the condition that is satisfied, either:
</p>

<ol>
<li>
the type <code>U</code>; or
</li>
<li>
the type <code>chan E</code> if <code>T</code> contains only bidirectional
channels, or the type <code>chan&lt;- E</code> or <code>&lt;-chan E</code>
depending on the direction of the directional channels present.
</li>
</ol>

<p>
By definition, a core type is never a <a href="#Type_definitions">defined type</a>,
<a href="#Type_parameter_declarations">type parameter</a>, or
<a href="#Interface_types">interface type</a>.
</p>

<p>
Examples of interfaces with core types:
</p>

<pre>
type Celsius float32
type Kelvin  float32

interface{ int }                          // int
interface{ Celsius|Kelvin }               // float32
interface{ ~chan int }                    // chan int
interface{ ~chan int|~chan&lt;- int }        // chan&lt;- int
interface{ ~[]*data; String() string }    // []*data
</pre>

<p>
Examples of interfaces without core types:
</p>

<pre>
interface{}                               // no single underlying type
interface{ Celsius|float64 }              // no single underlying type
interface{ chan int | chan&lt;- string }     // channels have different element types
interface{ &lt;-chan int | chan&lt;- int }      // directional channels have different directions
</pre>

<p>
Some operations (<a href="#Slice_expressions">slice expressions</a>,
<a href="#Appending_and_copying_slices"><code>append</code> and <code>copy</code></a>)
rely on a slightly more loose form of core types which accept byte slices and strings.
Specifically, if there are exactly two types, <code>[]byte</code> and <code>string</code>,
which are the underlying types of all types in the type set of interface <code>T</code>,
the core type of <code>T</code> is called <code>bytestring</code>.
</p>

<p>
Examples of interfaces with <code>bytestring</code> core types:
</p>

<pre>
interface{ int }                          // int (same as ordinary core type)
interface{ []byte | string }              // bytestring
interface{ ~[]byte | myString }           // bytestring
</pre>

<p>
Note that <code>bytestring</code> is not a real type; it cannot be used to declare
variables or compose other types. It exists solely to describe the behavior of some
operations that read from a sequence of bytes, which may be a byte slice or a string.
</p>

## Type identity

<p>
Two types are either <i>identical</i> or <i>different</i>.
</p>

<p>
A <a href="#Types">named type</a> is always different from any other type.
Otherwise, two types are identical if their <a href="#Types">underlying</a> type literals are
structurally equivalent; that is, they have the same literal structure and corresponding
components have identical types. In detail:
</p>

<ul>
	<li>Two array types are identical if they have identical element types and
	    the same array length.</li>

	<li>Two slice types are identical if they have identical element types.</li>

	<li>Two struct types are identical if they have the same sequence of fields,
	    and if corresponding fields have the same names, and identical types,
	    and identical tags.
	    <a href="#Exported_identifiers">Non-exported</a> field names from different
	    packages are always different.</li>

	<li>Two pointer types are identical if they have identical base types.</li>

	<li>Two function types are identical if they have the same number of parameters
	    and result values, corresponding parameter and result types are
	    identical, and either both functions are variadic or neither is.
	    Parameter and result names are not required to match.</li>

	<li>Two interface types are identical if they define the same type set.
	</li>

	<li>Two map types are identical if they have identical key and element types.</li>

	<li>Two channel types are identical if they have identical element types and
	    the same direction.</li>

	<li>Two <a href="#Instantiations">instantiated</a> types are identical if
	    their defined types and all type arguments are identical.
	</li>
</ul>

<p>
Given the declarations
</p>

<pre>
type (
	A0 = []string
	A1 = A0
	A2 = struct{ a, b int }
	A3 = int
	A4 = func(A3, float64) *A0
	A5 = func(x int, _ float64) *[]string

	B0 A0
	B1 []string
	B2 struct{ a, b int }
	B3 struct{ a, c int }
	B4 func(int, float64) *B0
	B5 func(x int, y float64) *A1

	C0 = B0
	D0[P1, P2 any] struct{ x P1; y P2 }
	E0 = D0[int, string]
)
</pre>

<p>
these types are identical:
</p>

<pre>
A0, A1, and []string
A2 and struct{ a, b int }
A3 and int
A4, func(int, float64) *[]string, and A5

B0 and C0
D0[int, string] and E0
[]int and []int
struct{ a, b *B5 } and struct{ a, b *B5 }
func(x int, y float64) *[]string, func(int, float64) (result *[]string), and A5
</pre>

<p>
<code>B0</code> and <code>B1</code> are different because they are new types
created by distinct <a href="#Type_definitions">type definitions</a>;
<code>func(int, float64) *B0</code> and <code>func(x int, y float64) *[]string</code>
are different because <code>B0</code> is different from <code>[]string</code>;
and <code>P1</code> and <code>P2</code> are different because they are different
type parameters.
<code>D0[int, string]</code> and <code>struct{ x int; y string }</code> are
different because the former is an <a href="#Instantiations">instantiated</a>
defined type while the latter is a type literal
(but they are still <a href="#Assignability">assignable</a>).
</p>

## Assignability

* value `x` / type `V` & variable / type `T`
  * 👀`x` is assignable -- to -- that variable, if >= 1 condition fulfill👀
    * `V` == `T`
      * == identical
    * `V` & `T` have identical [underlying types](#underlying-types) BUT NOT type parameters & (`V` or `T` are NOT named type)
    * `V` & `T` == channel types / identical element types + `V` == bidirectional channel + (>=1 of `V` or `T` NOT a named type)
    * TODO:
<li>
<code>T</code> is an interface type, but not a type parameter, and
<code>x</code> <a href="#Implementing_an_interface">implements</a> <code>T</code>.
</li>
<li>
<code>x</code> is the predeclared identifier <code>nil</code> and <code>T</code>
is a pointer, function, slice, map, channel, or interface type,
but not a type parameter.
</li>
<li>
<code>x</code> is an untyped <a href="#Constants">constant</a>
<a href="#Representability">representable</a>
by a value of type <code>T</code>.
</li>
</ul>

<p>
Additionally, if <code>x</code>'s type <code>V</code> or <code>T</code> are type parameters, <code>x</code>
is assignable to a variable of type <code>T</code> if one of the following conditions applies:
</p>

<ul>
<li>
<code>x</code> is the predeclared identifier <code>nil</code>, <code>T</code> is
a type parameter, and <code>x</code> is assignable to each type in
<code>T</code>'s type set.
</li>
<li>
<code>V</code> is not a <a href="#Types">named type</a>, <code>T</code> is
a type parameter, and <code>x</code> is assignable to each type in
<code>T</code>'s type set.
</li>
<li>
<code>V</code> is a type parameter and <code>T</code> is not a named type,
and values of each type in <code>V</code>'s type set are assignable
to <code>T</code>.
</li>
</ul>

<h3 id="Representability">Representability</h3>

<p>
A <a href="#Constants">constant</a> <code>x</code> is <i>representable</i>
by a value of type <code>T</code>,
where <code>T</code> is not a <a href="#Type_parameter_declarations">type parameter</a>,
if one of the following conditions applies:
</p>

<ul>
<li>
<code>x</code> is in the set of values <a href="#Types">determined</a> by <code>T</code>.
</li>

<li>
<code>T</code> is a <a href="#Numeric_types">floating-point type</a> and <code>x</code> can be rounded to <code>T</code>'s
precision without overflow. Rounding uses IEEE 754 round-to-even rules but with an IEEE
negative zero further simplified to an unsigned zero. Note that constant values never result
in an IEEE negative zero, NaN, or infinity.
</li>

<li>
<code>T</code> is a complex type, and <code>x</code>'s
<a href="#Complex_numbers">components</a> <code>real(x)</code> and <code>imag(x)</code>
are representable by values of <code>T</code>'s component type (<code>float32</code> or
<code>float64</code>).
</li>
</ul>

<p>
If <code>T</code> is a type parameter,
<code>x</code> is representable by a value of type <code>T</code> if <code>x</code> is representable
by a value of each type in <code>T</code>'s type set.
</p>

<pre>
x                   T           x is representable by a value of T because

'a'                 byte        97 is in the set of byte values
97                  rune        rune is an alias for int32, and 97 is in the set of 32-bit integers
"foo"               string      "foo" is in the set of string values
1024                int16       1024 is in the set of 16-bit integers
42.0                byte        42 is in the set of unsigned 8-bit integers
1e10                uint64      10000000000 is in the set of unsigned 64-bit integers
2.718281828459045   float32     2.718281828459045 rounds to 2.7182817 which is in the set of float32 values
-1e-1000            float64     -1e-1000 rounds to IEEE -0.0 which is further simplified to 0.0
0i                  int         0 is an integer value
(42 + 0i)           float32     42.0 (with zero imaginary part) is in the set of float32 values
</pre>

<pre>
x                   T           x is not representable by a value of T because

0                   bool        0 is not in the set of boolean values
'a'                 string      'a' is a rune, it is not in the set of string values
1024                byte        1024 is not in the set of unsigned 8-bit integers
-1                  uint16      -1 is not in the set of unsigned 16-bit integers
1.1                 int         1.1 is not an integer value
42i                 float32     (0 + 42i) is not in the set of float32 values
1e1000              float64     1e1000 overflows to IEEE +Inf after rounding
</pre>

## Method sets

<p>
The <i>method set</i> of a type determines the methods that can be
<a href="#Calls">called</a> on an <a href="#Operands">operand</a> of that type.
Every type has a (possibly empty) method set associated with it:
</p>

<ul>
<li>The method set of a <a href="#Type_definitions">defined type</a> <code>T</code> consists of all
<a href="#Method_declarations">methods</a> declared with receiver type <code>T</code>.
</li>

<li>
The method set of a pointer to a defined type <code>T</code>
(where <code>T</code> is neither a pointer nor an interface)
is the set of all methods declared with receiver <code>*T</code> or <code>T</code>.
</li>

<li>The method set of an <a href="#Interface_types">interface type</a> is the intersection
of the method sets of each type in the interface's <a href="#Interface_types">type set</a>
(the resulting method set is usually just the set of declared methods in the interface).
</li>
</ul>

<p>
Further rules apply to structs (and pointer to structs) containing embedded fields,
as described in the section on <a href="#Struct_types">struct types</a>.
Any other type has an empty method set.
</p>

<p>
In a method set, each method must have a
<a href="#Uniqueness_of_identifiers">unique</a>
non-<a href="#Blank_identifier">blank</a> <a href="#MethodName">method name</a>.
</p>

# Blocks

```go
Block = "{" StatementList "}" .
StatementList = { Statement ";" } .
```

* == (POSSIBLY) empty sequence of declarations & statements | `{}`

* types of blocks
  * explicit
  * implicit
    * universe block
      * == ALL Go source built-in
    * package block
      * == ALL [package](#packages)'s Go source
    * file's block
      * == ALL Go source text | that file
    * `if`, `for`, `switch`'s implicit blocks

* nest & influence [scoping](#declarations-and-scope)

# Declarations and scope

* declaration
  * binds a non-<a href="#Blank_identifier">blank</a> identifier to a
  <a href="#Constant_declarations">constant</a>,
  <a href="#Type_declarations">type</a>,
  <a href="#Type_parameter_declarations">type parameter</a>,
  <a href="#Variable_declarations">variable</a>,
  <a href="#Function_declarations">function</a>,
  <a href="#Labeled_statements">label</a>, or
  <a href="#Import_declarations">package</a>.
  Every identifier in a program must be declared.
  No identifier may be declared twice in the same block, and
  no identifier may be declared in both the file and package block.


<p>
The <a href="#Blank_identifier">blank identifier</a> may be used like any other identifier
in a declaration, but it does not introduce a binding and thus is not declared.
In the package block, the identifier <code>init</code> may only be used for
<a href="#Package_initialization"><code>init</code> function</a> declarations,
and like the blank identifier it does not introduce a new binding.
</p>

<pre class="ebnf">
Declaration   = ConstDecl | TypeDecl | VarDecl .
TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .
</pre>

<p>
The <i>scope</i> of a declared identifier is the extent of source text in which
the identifier denotes the specified constant, type, variable, function, label, or package.
</p>

<p>
Go is lexically scoped using <a href="#Blocks">blocks</a>:
</p>

<ol>
	<li>The scope of a <a href="#Predeclared_identifiers">predeclared identifier</a> is the universe block.</li>

	<li>The scope of an identifier denoting a constant, type, variable,
	    or function (but not method) declared at top level (outside any
	    function) is the package block.</li>

	<li>The scope of the package name of an imported package is the file block
	    of the file containing the import declaration.</li>

	<li>The scope of an identifier denoting a method receiver, function parameter,
	    or result variable is the function body.</li>

	<li>The scope of an identifier denoting a type parameter of a function
	    or declared by a method receiver begins after the name of the function
	    and ends at the end of the function body.</li>

	<li>The scope of an identifier denoting a type parameter of a type
	    begins after the name of the type and ends at the end
	    of the TypeSpec.</li>

	<li>The scope of a constant or variable identifier declared
	    inside a function begins at the end of the ConstSpec or VarSpec
	    (ShortVarDecl for short variable declarations)
	    and ends at the end of the innermost containing block.</li>

	<li>The scope of a type identifier declared inside a function
	    begins at the identifier in the TypeSpec
	    and ends at the end of the innermost containing block.</li>
</ol>

<p>
An identifier declared in a block may be redeclared in an inner block.
While the identifier of the inner declaration is in scope, it denotes
the entity declared by the inner declaration.
</p>

<p>
The <a href="#Package_clause">package clause</a> is not a declaration; the package name
does not appear in any scope. Its purpose is to identify the files belonging
to the same <a href="#Packages">package</a> and to specify the default package name for import
declarations.
</p>


<h3 id="Label_scopes">Label scopes</h3>

<p>
Labels are declared by <a href="#Labeled_statements">labeled statements</a> and are
used in the <a href="#Break_statements">"break"</a>,
<a href="#Continue_statements">"continue"</a>, and
<a href="#Goto_statements">"goto"</a> statements.
It is illegal to define a label that is never used.
In contrast to other identifiers, labels are not block scoped and do
not conflict with identifiers that are not labels. The scope of a label
is the body of the function in which it is declared and excludes
the body of any nested function.
</p>


## Blank identifier -- `_` --

* blank identifier
  * `_`
  * uses
    * 👀anonymous placeholder👀
      * != regular (non-blank) identifier
      * == ignore values / NOT need
    * | [declarations](#declarations-and-scope)
      * == [operand](#operands)
    * | [assignment statements](#assignment-statements)

## Predeclared identifiers

* identifiers / implicitly declared | 
<a href="#Blocks">universe block</a>
[<a href="#Go_1.18">Go 1.18</a>]
[<a href="#Go_1.21">Go 1.21</a>]:

<pre class="grammar">
Types:
	any bool byte comparable
	complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap clear close complex copy delete imag len
	make max min new panic print println real recover
</pre>

## Exported identifiers

<p>
An identifier may be <i>exported</i> to permit access to it from another package.
An identifier is exported if both:
</p>
<ol>
	<li>the first character of the identifier's name is a Unicode uppercase
	letter (Unicode character category Lu); and</li>
	<li>the identifier is declared in the <a href="#Blocks">package block</a>
	or it is a <a href="#Struct_types">field name</a> or
	<a href="#MethodName">method name</a>.</li>
</ol>
<p>
All other identifiers are not exported.
</p>

## Uniqueness of identifiers

<p>
Given a set of identifiers, an identifier is called <i>unique</i> if it is
<i>different</i> from every other in the set.
Two identifiers are different if they are spelled differently, or if they
appear in different <a href="#Packages">packages</a> and are not
<a href="#Exported_identifiers">exported</a>. Otherwise, they are the same.
</p>

## Constant declarations

<p>
A constant declaration binds a list of identifiers (the names of
the constants) to the values of a list of <a href="#Constant_expressions">constant expressions</a>.
The number of identifiers must be equal
to the number of expressions, and the <i>n</i>th identifier on
the left is bound to the value of the <i>n</i>th expression on the
right.
</p>

```go
ConstDecl      = "const" ( ConstSpec | "(" { ConstSpec ";" } ")" ) .
ConstSpec      = IdentifierList [ [ Type ] "=" ExpressionList ] .

IdentifierList = identifier { "," identifier } .
ExpressionList = Expression { "," Expression } .
```

<p>
If the type is present, all constants take the type specified, and
the expressions must be <a href="#Assignability">assignable</a> to that type,
which must not be a type parameter.
If the type is omitted, the constants take the
individual types of the corresponding expressions.
If the expression values are untyped <a href="#Constants">constants</a>,
the declared constants remain untyped and the constant identifiers
denote the constant values. For instance, if the expression is a
floating-point literal, the constant identifier denotes a floating-point
constant, even if the literal's fractional part is zero.
</p>

<pre>
const Pi float64 = 3.14159265358979323846
const zero = 0.0         // untyped floating-point constant
const (
	size int64 = 1024
	eof        = -1  // untyped integer constant
)
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
</pre>

<p>
Within a parenthesized <code>const</code> declaration list the
expression list may be omitted from any but the first ConstSpec.
Such an empty list is equivalent to the textual substitution of the
first preceding non-empty expression list and its type if any.
Omitting the list of expressions is therefore equivalent to
repeating the previous list.  The number of identifiers must be equal
to the number of expressions in the previous list.
Together with the <a href="#Iota"><code>iota</code> constant generator</a>
this mechanism permits light-weight declaration of sequential values:
</p>

<pre>
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays  // this constant is not exported
)
</pre>


## Iota

* := predeclared identifier / 
  * == successive constants
    * untyped
    * integer

* uses
  * | [constant declaration](#constant-declarations), 
  
* its value
  * | constant declaration,
    * ConstSpec's respective index /
      * index starts -- from -- 0

* use cases
  * construct a set of related constants

<pre>
const (
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const (
	a = 1 &lt;&lt; iota  // a == 1  (iota == 0)
	b = 1 &lt;&lt; iota  // b == 2  (iota == 1)
	c = 3          // c == 3  (iota == 2, unused)
	d = 1 &lt;&lt; iota  // d == 8  (iota == 3)
)

const (
	u         = iota * 42  // u == 0     (untyped integer constant)
	v float64 = iota * 42  // v == 42.0  (float64 constant)
	w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0
const y = iota  // y == 0
</pre>

<p>
By definition, multiple uses of <code>iota</code> in the same ConstSpec all have the same value:
</p>

<pre>
const (
	bit0, mask0 = 1 &lt;&lt; iota, 1&lt;&lt;iota - 1  // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                           // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                  //                        (iota == 2, unused)
	bit3, mask3                           // bit3 == 8, mask3 == 7  (iota == 3)
)
</pre>

<p>
This last example exploits the <a href="#Constant_declarations">implicit repetition</a>
of the last non-empty expression list.
</p>


## Type declarations

* type declaration
  * 👀binds an identifier (name) -- to a -- [type](#types)👀
  * forms
    * alias declarations
    * type definitions
  ```
  TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) 
  TypeSpec = AliasDecl | TypeDef
  ```

### Alias declarations

* alias declaration
  * binds an identifier -- to the -- given type
    * [Go 1.9](#go-19)
  ```
  AliasDecl = identifier "=" Type 
  ```

```
type (
	nodeList = []*Node  // nodeList type == []*Node type
	Polar    = polar    // Polar type == polar type
)
```

### Type definitions

<p>
A type definition creates a new, distinct type with the same
<a href="#Underlying_types">underlying type</a> and operations as the given type
and binds an identifier, the <i>type name</i>, to it.
</p>

<pre class="ebnf">
TypeDef = identifier [ TypeParameters ] Type .
</pre>

<p>
The new type is called a <i>defined type</i>.
It is <a href="#Type_identity">different</a> from any other type,
including the type it is created from.
</p>

<pre>
type (
	Point struct{ x, y float64 }  // Point and struct{ x, y float64 } are different types
	polar Point                   // polar and Point denote different types
)

type TreeNode struct {
	left, right *TreeNode
	value any
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
</pre>

<p>
A defined type may have <a href="#Method_declarations">methods</a> associated with it.
It does not inherit any methods bound to the given type,
but the <a href="#Method_sets">method set</a>
of an interface type or of elements of a composite type remains unchanged:
</p>

<pre>
// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of PtrMutex's underlying type *Mutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex.
type PrintableMutex struct {
	Mutex
}

// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block
</pre>

<p>
Type definitions may be used to define different boolean, numeric,
or string types and associate methods with them:
</p>

<pre>
type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT%+dh", tz)
}
</pre>

<p>
If the type definition specifies <a href="#Type_parameter_declarations">type parameters</a>,
the type name denotes a <i>generic type</i>.
Generic types must be <a href="#Instantiations">instantiated</a> when they
are used.
</p>

<pre>
type List[T any] struct {
	next  *List[T]
	value T
}
</pre>

<p>
In a type definition the given type cannot be a type parameter.
</p>

<pre>
type T[P any] P    // illegal: P is a type parameter

func f[T any]() {
	type L T   // illegal: T is a type parameter declared by the enclosing function
}
</pre>

<p>
A generic type may also have <a href="#Method_declarations">methods</a> associated with it.
In this case, the method receivers must declare the same number of type parameters as
present in the generic type definition.
</p>

<pre>
// The method Len returns the number of elements in the linked list l.
func (l *List[T]) Len() int  { … }
</pre>

## Type parameter declarations

```
TypeParameters  = "[" TypeParamList [ "," ] "]" .
TypeParamList   = TypeParamDecl { "," TypeParamDecl } .
TypeParamDecl   = IdentifierList TypeConstraint .
```

* type parameter list
  * uses | 
    * generic function
      * ⚠️!= function parameter list ⚠️
    * generic type declaration
  * how to use?
    * | instantiate generic function OR generic type,
      * 👀replaced -- with the -- type argument👀

* `IdentifierList`
  * ⚠️non-blank names, MUST be unique ⚠️

* TODO:
<p>
A parsing ambiguity arises when the type parameter list for a generic type
declares a single type parameter <code>P</code> with a constraint <code>C</code>
such that the text <code>P C</code> forms a valid expression:
</p>

<pre>
type T[P *C] …
type T[P (C)] …
type T[P *C|Q] …
…
</pre>

<p>
In these rare cases, the type parameter list is indistinguishable from an
expression and the type declaration is parsed as an array type declaration.
To resolve the ambiguity, embed the constraint in an
<a href="#Interface_types">interface</a> or use a trailing comma:
</p>

<pre>
type T[P interface{*C}] …
type T[P *C,] …
</pre>

<p>
Type parameters may also be declared by the receiver specification
of a <a href="#Method_declarations">method declaration</a> associated
with a generic type.
</p>

<p>
Within a type parameter list of a generic type <code>T</code>, a type constraint
may not (directly, or indirectly through the type parameter list of another
generic type) refer to <code>T</code>.
</p>

<pre>
type T1[P T1[P]] …                    // illegal: T1 refers to itself
type T2[P interface{ T2[int] }] …     // illegal: T2 refers to itself
type T3[P interface{ m(T3[int])}] …   // illegal: T3 refers to itself
type T4[P T5[P]] …                    // illegal: T4 refers to T5 and
type T5[P T4[P]] …                    //          T5 refers to T4

type T6[P int] struct{ f *T6[P] }     // ok: reference to T6 is not in type parameter list
</pre>

### Type constraints

```go
TypeConstraint = TypeElem .
```

* == interface / 
  * define ALLOWED type parameter's type arguments 
  * controls the operations -- supported by -- that type parameter's values
  * 👀if interface literal's form == `interface{E}` / `E` == embedded type element -> `interface{ … }` may be omitted👀
* use cases
  * 👀[type parameter](#type-parameter-declarations)'s meta-type👀

* TODO:
<p>
The <a href="#Predeclared_identifiers">predeclared</a>
<a href="#Interface_types">interface type</a> <code>comparable</code>
denotes the set of all non-interface types that are
<a href="#Comparison_operators">strictly comparable</a>
[<a href="#Go_1.18">Go 1.18</a>].
</p>

<p>
Even though interfaces that are not type parameters are <a href="#Comparison_operators">comparable</a>,
they are not strictly comparable and therefore they do not implement <code>comparable</code>.
However, they <a href="#Satisfying_a_type_constraint">satisfy</a> <code>comparable</code>.
</p>

<pre>
int                          // implements comparable (int is strictly comparable)
[]byte                       // does not implement comparable (slices cannot be compared)
interface{}                  // does not implement comparable (see above)
interface{ ~int | ~string }  // type parameter only: implements comparable (int, string types are strictly comparable)
interface{ comparable }      // type parameter only: implements comparable (comparable implements itself)
interface{ ~int | ~[]byte }  // type parameter only: does not implement comparable (slices are not comparable)
interface{ ~struct{ any } }  // type parameter only: does not implement comparable (field any is not strictly comparable)
</pre>

<p>
The <code>comparable</code> interface and interfaces that (directly or indirectly) embed
<code>comparable</code> may only be used as type constraints. They cannot be the types of
values or variables, or components of other, non-interface types.
</p>

<h4 id="Satisfying_a_type_constraint">Satisfying a type constraint</h4>

<p>
A type argument <code>T</code><i> satisfies</i> a type constraint <code>C</code>
if <code>T</code> is an element of the type set defined by <code>C</code>; i.e.,
if <code>T</code> <a href="#Implementing_an_interface">implements</a> <code>C</code>.
As an exception, a <a href="#Comparison_operators">strictly comparable</a>
type constraint may also be satisfied by a <a href="#Comparison_operators">comparable</a>
(not necessarily strictly comparable) type argument
[<a href="#Go_1.20">Go 1.20</a>].
More precisely:
</p>

<p>
A type T <i>satisfies</i> a constraint <code>C</code> if
</p>

<ul>
<li>
	<code>T</code> <a href="#Implementing_an_interface">implements</a> <code>C</code>; or
</li>
<li>
	<code>C</code> can be written in the form <code>interface{ comparable; E }</code>,
	where <code>E</code> is a <a href="#Basic_interfaces">basic interface</a> and
	<code>T</code> is <a href="#Comparison_operators">comparable</a> and implements <code>E</code>.
</li>
</ul>

<pre>
type argument      type constraint                // constraint satisfaction

int                interface{ ~int }              // satisfied: int implements interface{ ~int }
string             comparable                     // satisfied: string implements comparable (string is strictly comparable)
[]byte             comparable                     // not satisfied: slices are not comparable
any                interface{ comparable; int }   // not satisfied: any does not implement interface{ int }
any                comparable                     // satisfied: any is comparable and implements the basic interface any
struct{f any}      comparable                     // satisfied: struct{f any} is comparable and implements the basic interface any
any                interface{ comparable; m() }   // not satisfied: any does not implement the basic interface interface{ m() }
interface{ m() }   interface{ comparable; m() }   // satisfied: interface{ m() } is comparable and implements the basic interface interface{ m() }
</pre>

<p>
Because of the exception in the constraint satisfaction rule, comparing operands of type parameter type
may panic at run-time (even though comparable type parameters are always strictly comparable).
</p>

## Variable declarations

<p>
A variable declaration creates one or more <a href="#Variables">variables</a>,
binds corresponding identifiers to them, and gives each a type and an initial value.
</p>

<pre class="ebnf">
VarDecl     = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec     = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
</pre>

<pre>
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"
</pre>

<p>
If a list of expressions is given, the variables are initialized
with the expressions following the rules for <a href="#Assignment_statements">assignment statements</a>.
Otherwise, each variable is initialized to its <a href="#The_zero_value">zero value</a>.
</p>

<p>
If a type is present, each variable is given that type.
Otherwise, each variable is given the type of the corresponding
initialization value in the assignment.
If that value is an untyped constant, it is first implicitly
<a href="#Conversions">converted</a> to its <a href="#Constants">default type</a>;
if it is an untyped boolean value, it is first implicitly converted to type <code>bool</code>.
The predeclared value <code>nil</code> cannot be used to initialize a variable
with no explicit type.
</p>

<pre>
var d = math.Sin(0.5)  // d is float64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal
</pre>

<p>
Implementation restriction: A compiler may make it illegal to declare a variable
inside a <a href="#Function_declarations">function body</a> if the variable is
never used.
</p>

<h3 id="Short_variable_declarations">Short variable declarations</h3>

<p>
A <i>short variable declaration</i> uses the syntax:
</p>

<pre class="ebnf">
ShortVarDecl = IdentifierList ":=" ExpressionList .
</pre>

<p>
It is shorthand for a regular <a href="#Variable_declarations">variable declaration</a>
with initializer expressions but no types:
</p>

<pre class="grammar">
"var" IdentifierList "=" ExpressionList .
</pre>

<pre>
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w, _ := os.Pipe()  // os.Pipe() returns a connected pair of Files and an error, if any
_, y, _ := coord(p)   // coord() returns three values; only interested in y coordinate
</pre>

<p>
Unlike regular variable declarations, a short variable declaration may <i>redeclare</i>
variables provided they were originally declared earlier in the same block
(or the parameter lists if the block is the function body) with the same type,
and at least one of the non-<a href="#Blank_identifier">blank</a> variables is new.
As a consequence, redeclaration can only appear in a multi-variable short declaration.
Redeclaration does not introduce a new variable; it just assigns a new value to the original.
The non-blank variable names on the left side of <code>:=</code>
must be <a href="#Uniqueness_of_identifiers">unique</a>.
</p>

<pre>
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
x, y, x := 1, 2, 3                        // illegal: x repeated on left side of :=
</pre>

<p>
Short variable declarations may appear only inside functions.
In some contexts such as the initializers for
<a href="#If_statements">"if"</a>,
<a href="#For_statements">"for"</a>, or
<a href="#Switch_statements">"switch"</a> statements,
they can be used to declare local temporary variables.
</p>

## Function declarations

* allows
  * function name is -- bound to a -- function

```
FunctionDecl = "func" FunctionName [ TypeParameters ] Signature [ FunctionBody ] .
FunctionName = identifier .
FunctionBody = Block .
```

* if the `Signature` has result parameters -> `[ FunctionBody ]` MUST end in [terminating statement](#terminating-statements)
* cases / NO declared `FunctionBody`
  * declare interface

* TODO: 
<p>
If the function declaration specifies <a href="#Type_parameter_declarations">type parameters</a>,
the function name denotes a <i>generic function</i>.
A generic function must be <a href="#Instantiations">instantiated</a> before it can be
called or used as a value.
</p>

<pre>
func min[T ~int|~float64](x, y T) T {
	if x &lt; y {
		return x
	}
	return y
}
</pre>

<p>
A function declaration without type parameters may omit the body.
Such a declaration provides the signature for a function implemented outside Go,
such as an assembly routine.
</p>

<pre>
func flushICache(begin, end uintptr)  // implemented externally
</pre>

## Method declarations

<p>
A method is a <a href="#Function_declarations">function</a> with a <i>receiver</i>.
A method declaration binds an identifier, the <i>method name</i>, to a method,
and associates the method with the receiver's <i>base type</i>.
</p>

<pre class="ebnf">
MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
Receiver   = Parameters .
</pre>

<p>
The receiver is specified via an extra parameter section preceding the method
name. That parameter section must declare a single non-variadic parameter, the receiver.
Its type must be a <a href="#Type_definitions">defined</a> type <code>T</code> or a
pointer to a defined type <code>T</code>, possibly followed by a list of type parameter
names <code>[P1, P2, …]</code> enclosed in square brackets.
<code>T</code> is called the receiver <i>base type</i>. A receiver base type cannot be
a pointer or interface type and it must be defined in the same package as the method.
The method is said to be <i>bound</i> to its receiver base type and the method name
is visible only within <a href="#Selectors">selectors</a> for type <code>T</code>
or <code>*T</code>.
</p>

<p>
A non-<a href="#Blank_identifier">blank</a> receiver identifier must be
<a href="#Uniqueness_of_identifiers">unique</a> in the method signature.
If the receiver's value is not referenced inside the body of the method,
its identifier may be omitted in the declaration. The same applies in
general to parameters of functions and methods.
</p>

<p>
For a base type, the non-blank names of methods bound to it must be unique.
If the base type is a <a href="#Struct_types">struct type</a>,
the non-blank method and field names must be distinct.
</p>

<p>
Given defined type <code>Point</code> the declarations
</p>

<pre>
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
</pre>

<p>
bind the methods <code>Length</code> and <code>Scale</code>,
with receiver type <code>*Point</code>,
to the base type <code>Point</code>.
</p>

<p>
If the receiver base type is a <a href="#Type_declarations">generic type</a>, the
receiver specification must declare corresponding type parameters for the method
to use. This makes the receiver type parameters available to the method.
Syntactically, this type parameter declaration looks like an
<a href="#Instantiations">instantiation</a> of the receiver base type: the type
arguments must be identifiers denoting the type parameters being declared, one
for each type parameter of the receiver base type.
The type parameter names do not need to match their corresponding parameter names in the
receiver base type definition, and all non-blank parameter names must be unique in the
receiver parameter section and the method signature.
The receiver type parameter constraints are implied by the receiver base type definition:
corresponding type parameters have corresponding constraints.
</p>

<pre>
type Pair[A, B any] struct {
	a A
	b B
}

func (p Pair[A, B]) Swap() Pair[B, A]  { … }  // receiver declares A, B
func (p Pair[First, _]) First() First  { … }  // receiver declares First, corresponds to A in Pair
</pre>

# Expressions

<p>
An expression specifies the computation of a value by applying
operators and functions to operands.
</p>

## Operands

* Operands
  * := 👀elementary values | expression👀

```go
Operand     = Literal | OperandName [ TypeArgs ] | "(" Expression ")" .
Literal     = BasicLit | CompositeLit | FunctionLit .
BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
OperandName = identifier | QualifiedIdent .
```

*  if `OperandName` == [generic function](#function-declarations) -> may be `OperandName [ TypeArgs ]`
  * == instantiated function

* [blank identifier](#blank-identifier)
  * ⚠️may appear -- as an -- operand, ONLY | assignment statement's left-hand side⚠️
    ```go
    _ = somethingRightHandSide
    ```

* implementation restriction
  * ❌if an operand's type == type parameter / empty type set -> 
    * compiler need NOT report an error❌ 
    * functions can NOT be instantiated❌

## Qualified identifiers

```go
QualifiedIdent = PackageName "." identifier .
	// PackageName & identifier MUST NOT be blank identifier 
```

* requirements
  * [import](#import-declarations) the package
  * identifier MUST be 
    * [exported](#exported-identifiers)
    * declared | [package's block](#blocks)

## Composite literals

<p>
Composite literals construct new composite values each time they are evaluated.
They consist of the type of the literal followed by a brace-bound list of elements.
Each element may optionally be preceded by a corresponding key.
</p>

<pre class="ebnf">
CompositeLit  = LiteralType LiteralValue .
LiteralType   = StructType | ArrayType | "[" "..." "]" ElementType |
                SliceType | MapType | TypeName [ TypeArgs ] .
LiteralValue  = "{" [ ElementList [ "," ] ] "}" .
ElementList   = KeyedElement { "," KeyedElement } .
KeyedElement  = [ Key ":" ] Element .
Key           = FieldName | Expression | LiteralValue .
FieldName     = identifier .
Element       = Expression | LiteralValue .
</pre>

<p>
The LiteralType's <a href="#Core_types">core type</a> <code>T</code>
must be a struct, array, slice, or map type
(the syntax enforces this constraint except when the type is given
as a TypeName).
The types of the elements and keys must be <a href="#Assignability">assignable</a>
to the respective field, element, and key types of type <code>T</code>;
there is no additional conversion.
The key is interpreted as a field name for struct literals,
an index for array and slice literals, and a key for map literals.
For map literals, all elements must have a key. It is an error
to specify multiple elements with the same field name or
constant key value. For non-constant map keys, see the section on
<a href="#Order_of_evaluation">evaluation order</a>.
</p>

<p>
For struct literals the following rules apply:
</p>
<ul>
	<li>A key must be a field name declared in the struct type.
	</li>
	<li>An element list that does not contain any keys must
	    list an element for each struct field in the
	    order in which the fields are declared.
	</li>
	<li>If any element has a key, every element must have a key.
	</li>
	<li>An element list that contains keys does not need to
	    have an element for each struct field. Omitted fields
	    get the zero value for that field.
	</li>
	<li>A literal may omit the element list; such a literal evaluates
	    to the zero value for its type.
	</li>
	<li>It is an error to specify an element for a non-exported
	    field of a struct belonging to a different package.
	</li>
</ul>

<p>
Given the declarations
</p>
<pre>
type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }
</pre>

<p>
one may write
</p>

<pre>
origin := Point3D{}                            // zero value for Point3D
line := Line{origin, Point3D{y: -4, z: 12.3}}  // zero value for line.q.x
</pre>

<p>
For array and slice literals the following rules apply:
</p>
<ul>
	<li>Each element has an associated integer index marking
	    its position in the array.
	</li>
	<li>An element with a key uses the key as its index. The
	    key must be a non-negative constant
	    <a href="#Representability">representable</a> by
	    a value of type <code>int</code>; and if it is typed
	    it must be of <a href="#Numeric_types">integer type</a>.
	</li>
	<li>An element without a key uses the previous element's index plus one.
	    If the first element has no key, its index is zero.
	</li>
</ul>

<p>
<a href="#Address_operators">Taking the address</a> of a composite literal
generates a pointer to a unique <a href="#Variables">variable</a> initialized
with the literal's value.
</p>

<pre>
var pointer *Point3D = &amp;Point3D{y: 1000}
</pre>

<p>
Note that the <a href="#The_zero_value">zero value</a> for a slice or map
type is not the same as an initialized but empty value of the same type.
Consequently, taking the address of an empty slice or map composite literal
does not have the same effect as allocating a new slice or map value with
<a href="#Allocation">new</a>.
</p>

<pre>
p1 := &amp;[]int{}    // p1 points to an initialized, empty slice with value []int{} and length 0
p2 := new([]int)  // p2 points to an uninitialized slice with value nil and length 0
</pre>

<p>
The length of an array literal is the length specified in the literal type.
If fewer elements than the length are provided in the literal, the missing
elements are set to the zero value for the array element type.
It is an error to provide elements with index values outside the index range
of the array. The notation <code>...</code> specifies an array length equal
to the maximum element index plus one.
</p>

<pre>
buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2
</pre>

<p>
A slice literal describes the entire underlying array literal.
Thus the length and capacity of a slice literal are the maximum
element index plus one. A slice literal has the form
</p>

<pre>
[]T{x1, x2, … xn}
</pre>

<p>
and is shorthand for a slice operation applied to an array:
</p>

<pre>
tmp := [n]T{x1, x2, … xn}
tmp[0 : n]
</pre>

<p>
Within a composite literal of array, slice, or map type <code>T</code>,
elements or map keys that are themselves composite literals may elide the respective
literal type if it is identical to the element or key type of <code>T</code>.
Similarly, elements or keys that are addresses of composite literals may elide
the <code>&amp;T</code> when the element or key type is <code>*T</code>.
</p>

<pre>
[...]Point{{1.5, -3.5}, {0, 0}}     // same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}
[][]int{{1, 2, 3}, {4, 5}}          // same as [][]int{[]int{1, 2, 3}, []int{4, 5}}
[][]Point{{{0, 1}, {1, 2}}}         // same as [][]Point{[]Point{Point{0, 1}, Point{1, 2}}}
map[string]Point{"orig": {0, 0}}    // same as map[string]Point{"orig": Point{0, 0}}
map[Point]string{{0, 0}: "orig"}    // same as map[Point]string{Point{0, 0}: "orig"}

type PPoint *Point
[2]*Point{{1.5, -3.5}, {}}          // same as [2]*Point{&amp;Point{1.5, -3.5}, &amp;Point{}}
[2]PPoint{{1.5, -3.5}, {}}          // same as [2]PPoint{PPoint(&amp;Point{1.5, -3.5}), PPoint(&amp;Point{})}
</pre>

<p>
A parsing ambiguity arises when a composite literal using the
TypeName form of the LiteralType appears as an operand between the
<a href="#Keywords">keyword</a> and the opening brace of the block
of an "if", "for", or "switch" statement, and the composite literal
is not enclosed in parentheses, square brackets, or curly braces.
In this rare case, the opening brace of the literal is erroneously parsed
as the one introducing the block of statements. To resolve the ambiguity,
the composite literal must appear within parentheses.
</p>

<pre>
if x == (T{a,b,c}[i]) { … }
if (x == T{a,b,c}[i]) { … }
</pre>

<p>
Examples of valid array, slice, and map literals:
</p>

<pre>
// list of prime numbers
primes := []int{2, 3, 5, 7, 9, 2147483647}

// vowels[ch] is true if ch is a vowel
vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}

// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
</pre>


## Function literals

<p>
A function literal represents an anonymous <a href="#Function_declarations">function</a>.
Function literals cannot declare type parameters.
</p>

<pre class="ebnf">
FunctionLit = "func" Signature FunctionBody .
</pre>

<pre>
func(a, b int, z float64) bool { return a*b &lt; int(z) }
</pre>

<p>
A function literal can be assigned to a variable or invoked directly.
</p>

<pre>
f := func(x, y int) int { return x + y }
func(ch chan int) { ch &lt;- ACK }(replyChan)
</pre>

<p>
Function literals are <i>closures</i>: they may refer to variables
defined in a surrounding function. Those variables are then shared between
the surrounding function and the function literal, and they survive as long
as they are accessible.
</p>


## Primary expressions

* := operands -- for -- unary & binary expressions

```go
PrimaryExpr =
	Operand |
	Conversion |
	MethodExpr |
	PrimaryExpr Selector |
	PrimaryExpr Index |
	PrimaryExpr Slice |
	PrimaryExpr TypeAssertion |
	PrimaryExpr Arguments .

Selector       = "." identifier .
Index          = "[" Expression [ "," ] "]" .
Slice          = "[" [ Expression ] ":" [ Expression ] "]" |
                 "[" [ Expression ] ":" Expression ":" Expression "]" .
TypeAssertion  = "." "(" Type ")" .
Arguments      = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
```

## Selectors

<p>
For a <a href="#Primary_expressions">primary expression</a> <code>x</code>
that is not a <a href="#Package_clause">package name</a>, the
<i>selector expression</i>
</p>

<pre>
x.f
</pre>

<p>
denotes the field or method <code>f</code> of the value <code>x</code>
(or sometimes <code>*x</code>; see below).
The identifier <code>f</code> is called the (field or method) <i>selector</i>;
it must not be the <a href="#Blank_identifier">blank identifier</a>.
The type of the selector expression is the type of <code>f</code>.
If <code>x</code> is a package name, see the section on
<a href="#Qualified_identifiers">qualified identifiers</a>.
</p>

<p>
A selector <code>f</code> may denote a field or method <code>f</code> of
a type <code>T</code>, or it may refer
to a field or method <code>f</code> of a nested
<a href="#Struct_types">embedded field</a> of <code>T</code>.
The number of embedded fields traversed
to reach <code>f</code> is called its <i>depth</i> in <code>T</code>.
The depth of a field or method <code>f</code>
declared in <code>T</code> is zero.
The depth of a field or method <code>f</code> declared in
an embedded field <code>A</code> in <code>T</code> is the
depth of <code>f</code> in <code>A</code> plus one.
</p>

<p>
The following rules apply to selectors:
</p>

<ol>
<li>
For a value <code>x</code> of type <code>T</code> or <code>*T</code>
where <code>T</code> is not a pointer or interface type,
<code>x.f</code> denotes the field or method at the shallowest depth
in <code>T</code> where there is such an <code>f</code>.
If there is not exactly <a href="#Uniqueness_of_identifiers">one <code>f</code></a>
with shallowest depth, the selector expression is illegal.
</li>

<li>
For a value <code>x</code> of type <code>I</code> where <code>I</code>
is an interface type, <code>x.f</code> denotes the actual method with name
<code>f</code> of the dynamic value of <code>x</code>.
If there is no method with name <code>f</code> in the
<a href="#Method_sets">method set</a> of <code>I</code>, the selector
expression is illegal.
</li>

<li>
As an exception, if the type of <code>x</code> is a <a href="#Type_definitions">defined</a>
pointer type and <code>(*x).f</code> is a valid selector expression denoting a field
(but not a method), <code>x.f</code> is shorthand for <code>(*x).f</code>.
</li>

<li>
In all other cases, <code>x.f</code> is illegal.
</li>

<li>
If <code>x</code> is of pointer type and has the value
<code>nil</code> and <code>x.f</code> denotes a struct field,
assigning to or evaluating <code>x.f</code>
causes a <a href="#Run_time_panics">run-time panic</a>.
</li>

<li>
If <code>x</code> is of interface type and has the value
<code>nil</code>, <a href="#Calls">calling</a> or
<a href="#Method_values">evaluating</a> the method <code>x.f</code>
causes a <a href="#Run_time_panics">run-time panic</a>.
</li>
</ol>

<p>
For example, given the declarations:
</p>

<pre>
type T0 struct {
	x int
}

func (*T0) M0()

type T1 struct {
	y int
}

func (T1) M1()

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2()

type Q *T2

var t T2     // with t.T0 != nil
var p *T2    // with p != nil and (*p).T0 != nil
var q Q = p
</pre>

<p>
one may write:
</p>

<pre>
t.z          // t.z
t.y          // t.T1.y
t.x          // (*t.T0).x

p.z          // (*p).z
p.y          // (*p).T1.y
p.x          // (*(*p).T0).x

q.x          // (*(*q).T0).x        (*q).x is a valid field selector

p.M0()       // ((*p).T0).M0()      M0 expects *T0 receiver
p.M1()       // ((*p).T1).M1()      M1 expects T1 receiver
p.M2()       // p.M2()              M2 expects *T2 receiver
t.M2()       // (&amp;t).M2()           M2 expects *T2 receiver, see section on Calls
</pre>

<p>
but the following is invalid:
</p>

<pre>
q.M0()       // (*q).M0 is valid but not a field selector
</pre>


## Method expressions

<p>
If <code>M</code> is in the <a href="#Method_sets">method set</a> of type <code>T</code>,
<code>T.M</code> is a function that is callable as a regular function
with the same arguments as <code>M</code> prefixed by an additional
argument that is the receiver of the method.
</p>

<pre class="ebnf">
MethodExpr    = ReceiverType "." MethodName .
ReceiverType  = Type .
</pre>

<p>
Consider a struct type <code>T</code> with two methods,
<code>Mv</code>, whose receiver is of type <code>T</code>, and
<code>Mp</code>, whose receiver is of type <code>*T</code>.
</p>

<pre>
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
</pre>

<p>
The expression
</p>

<pre>
T.Mv
</pre>

<p>
yields a function equivalent to <code>Mv</code> but
with an explicit receiver as its first argument; it has signature
</p>

<pre>
func(tv T, a int) int
</pre>

<p>
That function may be called normally with an explicit receiver, so
these five invocations are equivalent:
</p>

<pre>
t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)
</pre>

<p>
Similarly, the expression
</p>

<pre>
(*T).Mp
</pre>

<p>
yields a function value representing <code>Mp</code> with signature
</p>

<pre>
func(tp *T, f float32) float32
</pre>

<p>
For a method with a value receiver, one can derive a function
with an explicit pointer receiver, so
</p>

<pre>
(*T).Mv
</pre>

<p>
yields a function value representing <code>Mv</code> with signature
</p>

<pre>
func(tv *T, a int) int
</pre>

<p>
Such a function indirects through the receiver to create a value
to pass as the receiver to the underlying method;
the method does not overwrite the value whose address is passed in
the function call.
</p>

<p>
The final case, a value-receiver function for a pointer-receiver method,
is illegal because pointer-receiver methods are not in the method set
of the value type.
</p>

<p>
Function values derived from methods are called with function call syntax;
the receiver is provided as the first argument to the call.
That is, given <code>f := T.Mv</code>, <code>f</code> is invoked
as <code>f(t, 7)</code> not <code>t.f(7)</code>.
To construct a function that binds the receiver, use a
<a href="#Function_literals">function literal</a> or
<a href="#Method_values">method value</a>.
</p>

<p>
It is legal to derive a function value from a method of an interface type.
The resulting function takes an explicit receiver of that interface type.
</p>

## Method values

<p>
If the expression <code>x</code> has static type <code>T</code> and
<code>M</code> is in the <a href="#Method_sets">method set</a> of type <code>T</code>,
<code>x.M</code> is called a <i>method value</i>.
The method value <code>x.M</code> is a function value that is callable
with the same arguments as a method call of <code>x.M</code>.
The expression <code>x</code> is evaluated and saved during the evaluation of the
method value; the saved copy is then used as the receiver in any calls,
which may be executed later.
</p>

<pre>
type S struct { *T }
type T int
func (t T) M() { print(t) }

t := new(T)
s := S{T: t}
f := t.M                    // receiver *t is evaluated and stored in f
g := s.M                    // receiver *(s.T) is evaluated and stored in g
*t = 42                     // does not affect stored receivers in f and g
</pre>

<p>
The type <code>T</code> may be an interface or non-interface type.
</p>

<p>
As in the discussion of <a href="#Method_expressions">method expressions</a> above,
consider a struct type <code>T</code> with two methods,
<code>Mv</code>, whose receiver is of type <code>T</code>, and
<code>Mp</code>, whose receiver is of type <code>*T</code>.
</p>

<pre>
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
var pt *T
func makeT() T
</pre>

<p>
The expression
</p>

<pre>
t.Mv
</pre>

<p>
yields a function value of type
</p>

<pre>
func(int) int
</pre>

<p>
These two invocations are equivalent:
</p>

<pre>
t.Mv(7)
f := t.Mv; f(7)
</pre>

<p>
Similarly, the expression
</p>

<pre>
pt.Mp
</pre>

<p>
yields a function value of type
</p>

<pre>
func(float32) float32
</pre>

<p>
As with <a href="#Selectors">selectors</a>, a reference to a non-interface method with a value receiver
using a pointer will automatically dereference that pointer: <code>pt.Mv</code> is equivalent to <code>(*pt).Mv</code>.
</p>

<p>
As with <a href="#Calls">method calls</a>, a reference to a non-interface method with a pointer receiver
using an addressable value will automatically take the address of that value: <code>t.Mp</code> is equivalent to <code>(&amp;t).Mp</code>.
</p>

<pre>
f := t.Mv; f(7)   // like t.Mv(7)
f := pt.Mp; f(7)  // like pt.Mp(7)
f := pt.Mv; f(7)  // like (*pt).Mv(7)
f := t.Mp; f(7)   // like (&amp;t).Mp(7)
f := makeT().Mp   // invalid: result of makeT() is not addressable
</pre>

<p>
Although the examples above use non-interface types, it is also legal to create a method value
from a value of interface type.
</p>

<pre>
var i interface { M(int) } = myVal
f := i.M; f(7)  // like i.M(7)
</pre>


## Index expressions

<p>
A primary expression of the form
</p>

<pre>
a[x]
</pre>

<p>
denotes the element of the array, pointer to array, slice, string or map <code>a</code> indexed by <code>x</code>.
The value <code>x</code> is called the <i>index</i> or <i>map key</i>, respectively.
The following rules apply:
</p>

<p>
If <code>a</code> is neither a map nor a type parameter:
</p>
<ul>
	<li>the index <code>x</code> must be an untyped constant or its
	    <a href="#Core_types">core type</a> must be an <a href="#Numeric_types">integer</a></li>
	<li>a constant index must be non-negative and
	    <a href="#Representability">representable</a> by a value of type <code>int</code></li>
	<li>a constant index that is untyped is given type <code>int</code></li>
	<li>the index <code>x</code> is <i>in range</i> if <code>0 &lt;= x &lt; len(a)</code>,
	    otherwise it is <i>out of range</i></li>
</ul>

<p>
For <code>a</code> of <a href="#Array_types">array type</a> <code>A</code>:
</p>
<ul>
	<li>a <a href="#Constants">constant</a> index must be in range</li>
	<li>if <code>x</code> is out of range at run time,
	    a <a href="#Run_time_panics">run-time panic</a> occurs</li>
	<li><code>a[x]</code> is the array element at index <code>x</code> and the type of
	    <code>a[x]</code> is the element type of <code>A</code></li>
</ul>

<p>
For <code>a</code> of <a href="#Pointer_types">pointer</a> to array type:
</p>
<ul>
	<li><code>a[x]</code> is shorthand for <code>(*a)[x]</code></li>
</ul>

<p>
For <code>a</code> of <a href="#Slice_types">slice type</a> <code>S</code>:
</p>
<ul>
	<li>if <code>x</code> is out of range at run time,
	    a <a href="#Run_time_panics">run-time panic</a> occurs</li>
	<li><code>a[x]</code> is the slice element at index <code>x</code> and the type of
	    <code>a[x]</code> is the element type of <code>S</code></li>
</ul>

<p>
For <code>a</code> of <a href="#String_types">string type</a>:
</p>
<ul>
	<li>a <a href="#Constants">constant</a> index must be in range
	    if the string <code>a</code> is also constant</li>
	<li>if <code>x</code> is out of range at run time,
	    a <a href="#Run_time_panics">run-time panic</a> occurs</li>
	<li><code>a[x]</code> is the non-constant byte value at index <code>x</code> and the type of
	    <code>a[x]</code> is <code>byte</code></li>
	<li><code>a[x]</code> may not be assigned to</li>
</ul>

<p>
For <code>a</code> of <a href="#Map_types">map type</a> <code>M</code>:
</p>
<ul>
	<li><code>x</code>'s type must be
	    <a href="#Assignability">assignable</a>
	    to the key type of <code>M</code></li>
	<li>if the map contains an entry with key <code>x</code>,
	    <code>a[x]</code> is the map element with key <code>x</code>
	    and the type of <code>a[x]</code> is the element type of <code>M</code></li>
	<li>if the map is <code>nil</code> or does not contain such an entry,
	    <code>a[x]</code> is the <a href="#The_zero_value">zero value</a>
	    for the element type of <code>M</code></li>
</ul>

<p>
For <code>a</code> of <a href="#Type_parameter_declarations">type parameter type</a> <code>P</code>:
</p>
<ul>
	<li>The index expression <code>a[x]</code> must be valid for values
	    of all types in <code>P</code>'s type set.</li>
	<li>The element types of all types in <code>P</code>'s type set must be identical.
	    In this context, the element type of a string type is <code>byte</code>.</li>
	<li>If there is a map type in the type set of <code>P</code>,
	    all types in that type set must be map types, and the respective key types
	    must be all identical.</li>
	<li><code>a[x]</code> is the array, slice, or string element at index <code>x</code>,
	    or the map element with key <code>x</code> of the type argument
	    that <code>P</code> is instantiated with, and the type of <code>a[x]</code> is
	    the type of the (identical) element types.</li>
	<li><code>a[x]</code> may not be assigned to if <code>P</code>'s type set
	    includes string types.</li>
</ul>

<p>
Otherwise <code>a[x]</code> is illegal.
</p>

<p>
An index expression on a map <code>a</code> of type <code>map[K]V</code>
used in an <a href="#Assignment_statements">assignment statement</a> or initialization of the special form
</p>

<pre>
v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]
</pre>

<p>
yields an additional untyped boolean value. The value of <code>ok</code> is
<code>true</code> if the key <code>x</code> is present in the map, and
<code>false</code> otherwise.
</p>

<p>
Assigning to an element of a <code>nil</code> map causes a
<a href="#Run_time_panics">run-time panic</a>.
</p>


## Slice expressions

<p>
Slice expressions construct a substring or slice from a string, array, pointer
to array, or slice. There are two variants: a simple form that specifies a low
and high bound, and a full form that also specifies a bound on the capacity.
</p>

### Simple slice expressions

<p>
The primary expression
</p>

<pre>
a[low : high]
</pre>

<p>
constructs a substring or slice. The <a href="#Core_types">core type</a> of
<code>a</code> must be a string, array, pointer to array, slice, or a
<a href="#Core_types"><code>bytestring</code></a>.
The <i>indices</i> <code>low</code> and
<code>high</code> select which elements of operand <code>a</code> appear
in the result. The result has indices starting at 0 and length equal to
<code>high</code>&nbsp;-&nbsp;<code>low</code>.
After slicing the array <code>a</code>
</p>

<pre>
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
</pre>

<p>
the slice <code>s</code> has type <code>[]int</code>, length 3, capacity 4, and elements
</p>

<pre>
s[0] == 2
s[1] == 3
s[2] == 4
</pre>

<p>
For convenience, any of the indices may be omitted. A missing <code>low</code>
index defaults to zero; a missing <code>high</code> index defaults to the length of the
sliced operand:
</p>

<pre>
a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]
</pre>

<p>
If <code>a</code> is a pointer to an array, <code>a[low : high]</code> is shorthand for
<code>(*a)[low : high]</code>.
</p>

<p>
For arrays or strings, the indices are <i>in range</i> if
<code>0</code> &lt;= <code>low</code> &lt;= <code>high</code> &lt;= <code>len(a)</code>,
otherwise they are <i>out of range</i>.
For slices, the upper index bound is the slice capacity <code>cap(a)</code> rather than the length.
A <a href="#Constants">constant</a> index must be non-negative and
<a href="#Representability">representable</a> by a value of type
<code>int</code>; for arrays or constant strings, constant indices must also be in range.
If both indices are constant, they must satisfy <code>low &lt;= high</code>.
If the indices are out of range at run time, a <a href="#Run_time_panics">run-time panic</a> occurs.
</p>

<p>
Except for <a href="#Constants">untyped strings</a>, if the sliced operand is a string or slice,
the result of the slice operation is a non-constant value of the same type as the operand.
For untyped string operands the result is a non-constant value of type <code>string</code>.
If the sliced operand is an array, it must be <a href="#Address_operators">addressable</a>
and the result of the slice operation is a slice with the same element type as the array.
</p>

<p>
If the sliced operand of a valid slice expression is a <code>nil</code> slice, the result
is a <code>nil</code> slice. Otherwise, if the result is a slice, it shares its underlying
array with the operand.
</p>

<pre>
var a [10]int
s1 := a[3:7]   // underlying array of s1 is array a; &amp;s1[2] == &amp;a[5]
s2 := s1[1:4]  // underlying array of s2 is underlying array of s1 which is array a; &amp;s2[1] == &amp;a[5]
s2[1] = 42     // s2[1] == s1[2] == a[5] == 42; they all refer to the same underlying array element

var s []int
s3 := s[:0]    // s3 == nil
</pre>


### Full slice expressions

<p>
The primary expression
</p>

<pre>
a[low : high : max]
</pre>

<p>
constructs a slice of the same type, and with the same length and elements as the simple slice
expression <code>a[low : high]</code>. Additionally, it controls the resulting slice's capacity
by setting it to <code>max - low</code>. Only the first index may be omitted; it defaults to 0.
The <a href="#Core_types">core type</a> of <code>a</code> must be an array, pointer to array,
or slice (but not a string).
After slicing the array <code>a</code>
</p>

<pre>
a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]
</pre>

<p>
the slice <code>t</code> has type <code>[]int</code>, length 2, capacity 4, and elements
</p>

<pre>
t[0] == 2
t[1] == 3
</pre>

<p>
As for simple slice expressions, if <code>a</code> is a pointer to an array,
<code>a[low : high : max]</code> is shorthand for <code>(*a)[low : high : max]</code>.
If the sliced operand is an array, it must be <a href="#Address_operators">addressable</a>.
</p>

<p>
The indices are <i>in range</i> if <code>0 &lt;= low &lt;= high &lt;= max &lt;= cap(a)</code>,
otherwise they are <i>out of range</i>.
A <a href="#Constants">constant</a> index must be non-negative and
<a href="#Representability">representable</a> by a value of type
<code>int</code>; for arrays, constant indices must also be in range.
If multiple indices are constant, the constants that are present must be in range relative to each
other.
If the indices are out of range at run time, a <a href="#Run_time_panics">run-time panic</a> occurs.
</p>

## Type assertions

<p>
For an expression <code>x</code> of <a href="#Interface_types">interface type</a>,
but not a <a href="#Type_parameter_declarations">type parameter</a>, and a type <code>T</code>,
the primary expression
</p>

<pre>
x.(T)
</pre>

<p>
asserts that <code>x</code> is not <code>nil</code>
and that the value stored in <code>x</code> is of type <code>T</code>.
The notation <code>x.(T)</code> is called a <i>type assertion</i>.
</p>
<p>
More precisely, if <code>T</code> is not an interface type, <code>x.(T)</code> asserts
that the dynamic type of <code>x</code> is <a href="#Type_identity">identical</a>
to the type <code>T</code>.
In this case, <code>T</code> must <a href="#Method_sets">implement</a> the (interface) type of <code>x</code>;
otherwise the type assertion is invalid since it is not possible for <code>x</code>
to store a value of type <code>T</code>.
If <code>T</code> is an interface type, <code>x.(T)</code> asserts that the dynamic type
of <code>x</code> <a href="#Implementing_an_interface">implements</a> the interface <code>T</code>.
</p>
<p>
If the type assertion holds, the value of the expression is the value
stored in <code>x</code> and its type is <code>T</code>. If the type assertion is false,
a <a href="#Run_time_panics">run-time panic</a> occurs.
In other words, even though the dynamic type of <code>x</code>
is known only at run time, the type of <code>x.(T)</code> is
known to be <code>T</code> in a correct program.
</p>

<pre>
var x interface{} = 7          // x has dynamic type int and value 7
i := x.(int)                   // i has type int and value 7

type I interface { m() }

func f(y I) {
	s := y.(string)        // illegal: string does not implement I (missing method m)
	r := y.(io.Reader)     // r has type io.Reader and the dynamic type of y must implement both I and io.Reader
	…
}
</pre>

<p>
A type assertion used in an <a href="#Assignment_statements">assignment statement</a> or initialization of the special form
</p>

<pre>
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
var v, ok interface{} = x.(T) // dynamic types of v and ok are T and bool
</pre>

<p>
yields an additional untyped boolean value. The value of <code>ok</code> is <code>true</code>
if the assertion holds. Otherwise it is <code>false</code> and the value of <code>v</code> is
the <a href="#The_zero_value">zero value</a> for type <code>T</code>.
No <a href="#Run_time_panics">run-time panic</a> occurs in this case.
</p>


## Calls

* EACH arguments' expression
  * restrictions 
    * assignable -- to the -- function's parameter types
    * ⚠️evaluated BEFORE calling the function⚠️

<p>
If <code>f</code> denotes a generic function, it must be
<a href="#Instantiations">instantiated</a> before it can be called
or used as a function value.
</p>

<p>
In a function call, the function value and arguments are evaluated in
<a href="#Order_of_evaluation">the usual order</a>.
After they are evaluated, the parameters of the call are passed by value to the function
and the called function begins execution.
The return parameters of the function are passed by value
back to the caller when the function returns.
</p>

<p>
Calling a <code>nil</code> function value
causes a <a href="#Run_time_panics">run-time panic</a>.
</p>

<p>
As a special case, if the return values of a function or method
<code>g</code> are equal in number and individually
assignable to the parameters of another function or method
<code>f</code>, then the call <code>f(g(<i>parameters_of_g</i>))</code>
will invoke <code>f</code> after binding the return values of
<code>g</code> to the parameters of <code>f</code> in order.  The call
of <code>f</code> must contain no parameters other than the call of <code>g</code>,
and <code>g</code> must have at least one return value.
If <code>f</code> has a final <code>...</code> parameter, it is
assigned the return values of <code>g</code> that remain after
assignment of regular parameters.
</p>

<pre>
func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func Join(s, t string) string {
	return s + t
}

if Join(Split(value, len(value)/2)) != value {
	log.Panic("test fails")
}
</pre>

<p>
A method call <code>x.m()</code> is valid if the <a href="#Method_sets">method set</a>
of (the type of) <code>x</code> contains <code>m</code> and the
argument list can be assigned to the parameter list of <code>m</code>.
If <code>x</code> is <a href="#Address_operators">addressable</a> and <code>&amp;x</code>'s method
set contains <code>m</code>, <code>x.m()</code> is shorthand
for <code>(&amp;x).m()</code>:
</p>

<pre>
var p Point
p.Scale(3.5)
</pre>

<p>
There is no distinct method type and there are no method literals.
</p>

## Passing arguments to `...` parameters

<p>
If <code>f</code> is <a href="#Function_types">variadic</a> with a final
parameter <code>p</code> of type <code>...T</code>, then within <code>f</code>
the type of <code>p</code> is equivalent to type <code>[]T</code>.
If <code>f</code> is invoked with no actual arguments for <code>p</code>,
the value passed to <code>p</code> is <code>nil</code>.
Otherwise, the value passed is a new slice
of type <code>[]T</code> with a new underlying array whose successive elements
are the actual arguments, which all must be <a href="#Assignability">assignable</a>
to <code>T</code>. The length and capacity of the slice is therefore
the number of arguments bound to <code>p</code> and may differ for each
call site.
</p>

<p>
Given the function and calls
</p>
<pre>
func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")
</pre>

<p>
within <code>Greeting</code>, <code>who</code> will have the value
<code>nil</code> in the first call, and
<code>[]string{"Joe", "Anna", "Eileen"}</code> in the second.
</p>

<p>
If the final argument is assignable to a slice type <code>[]T</code> and
is followed by <code>...</code>, it is passed unchanged as the value
for a <code>...T</code> parameter. In this case no new slice is created.
</p>

<p>
Given the slice <code>s</code> and call
</p>

<pre>
s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)
</pre>

<p>
within <code>Greeting</code>, <code>who</code> will have the same value as <code>s</code>
with the same underlying array.
</p>

## Instantiations

<p>
A generic function or type is <i>instantiated</i> by substituting <i>type arguments</i>
for the type parameters [<a href="#Go_1.18">Go 1.18</a>].
Instantiation proceeds in two steps:
</p>

<ol>
<li>
Each type argument is substituted for its corresponding type parameter in the generic
declaration.
This substitution happens across the entire function or type declaration,
including the type parameter list itself and any types in that list.
</li>

<li>
After substitution, each type argument must <a href="#Satisfying_a_type_constraint">satisfy</a>
the <a href="#Type_parameter_declarations">constraint</a> (instantiated, if necessary)
of the corresponding type parameter. Otherwise instantiation fails.
</li>
</ol>

<p>
Instantiating a type results in a new non-generic <a href="#Types">named type</a>;
instantiating a function produces a new non-generic function.
</p>

<pre>
type parameter list    type arguments    after substitution

[P any]                int               int satisfies any
[S ~[]E, E any]        []int, int        []int satisfies ~[]int, int satisfies any
[P io.Writer]          string            illegal: string doesn't satisfy io.Writer
[P comparable]         any               any satisfies (but does not implement) comparable
</pre>

<p>
When using a generic function, type arguments may be provided explicitly,
or they may be partially or completely <a href="#Type_inference">inferred</a>
from the context in which the function is used.
Provided that they can be inferred, type argument lists may be omitted entirely if the function is:
</p>

<ul>
<li>
	<a href="#Calls">called</a> with ordinary arguments,
</li>
<li>
	<a href="#Assignment_statements">assigned</a> to a variable with a known type
</li>
<li>
	<a href="#Calls">passed as an argument</a> to another function, or
</li>
<li>
	<a href="#Return_statements">returned as a result</a>.
</li>
</ul>

<p>
In all other cases, a (possibly partial) type argument list must be present.
If a type argument list is absent or partial, all missing type arguments
must be inferrable from the context in which the function is used.
</p>

<pre>
// sum returns the sum (concatenation, for strings) of its arguments.
func sum[T ~int | ~float64 | ~string](x... T) T { … }

x := sum                       // illegal: the type of x is unknown
intSum := sum[int]             // intSum has type func(x... int) int
a := intSum(2, 3)              // a has value 5 of type int
b := sum[float64](2.0, 3)      // b has value 5.0 of type float64
c := sum(b, -1)                // c has value 4.0 of type float64

type sumFunc func(x... string) string
var f sumFunc = sum            // same as var f sumFunc = sum[string]
f = sum                        // same as f = sum[string]
</pre>

<p>
A partial type argument list cannot be empty; at least the first argument must be present.
The list is a prefix of the full list of type arguments, leaving the remaining arguments
to be inferred. Loosely speaking, type arguments may be omitted from "right to left".
</p>

<pre>
func apply[S ~[]E, E any](s S, f func(E) E) S { … }

f0 := apply[]                  // illegal: type argument list cannot be empty
f1 := apply[[]int]             // type argument for S explicitly provided, type argument for E inferred
f2 := apply[[]string, string]  // both type arguments explicitly provided

var bytes []byte
r := apply(bytes, func(byte) byte { … })  // both type arguments inferred from the function arguments
</pre>

<p>
For a generic type, all type arguments must always be provided explicitly.
</p>

## Type inference

<p>
A use of a generic function may omit some or all type arguments if they can be
<i>inferred</i> from the context within which the function is used, including
the constraints of the function's type parameters.
Type inference succeeds if it can infer the missing type arguments
and <a href="#Instantiations">instantiation</a> succeeds with the
inferred type arguments.
Otherwise, type inference fails and the program is invalid.
</p>

<p>
Type inference uses the type relationships between pairs of types for inference:
For instance, a function argument must be <a href="#Assignability">assignable</a>
to its respective function parameter; this establishes a relationship between the
type of the argument and the type of the parameter.
If either of these two types contains type parameters, type inference looks for the
type arguments to substitute the type parameters with such that the assignability
relationship is satisfied.
Similarly, type inference uses the fact that a type argument must
<a href="#Satisfying_a_type_constraint">satisfy</a> the constraint of its respective
type parameter.
</p>

<p>
Each such pair of matched types corresponds to a <i>type equation</i> containing
one or multiple type parameters, from one or possibly multiple generic functions.
Inferring the missing type arguments means solving the resulting set of type
equations for the respective type parameters.
</p>

<p>
For example, given
</p>

<pre>
// dedup returns a copy of the argument slice with any duplicate entries removed.
func dedup[S ~[]E, E comparable](S) S { … }

type Slice []int
var s Slice
s = dedup(s)   // same as s = dedup[Slice, int](s)
</pre>

<p>
the variable <code>s</code> of type <code>Slice</code> must be assignable to
the function parameter type <code>S</code> for the program to be valid.
To reduce complexity, type inference ignores the directionality of assignments,
so the type relationship between <code>Slice</code> and <code>S</code> can be
expressed via the (symmetric) type equation <code>Slice ≡<sub>A</sub> S</code>
(or <code>S ≡<sub>A</sub> Slice</code> for that matter),
where the <code><sub>A</sub></code> in <code>≡<sub>A</sub></code>
indicates that the LHS and RHS types must match per assignability rules
(see the section on <a href="#Type_unification">type unification</a> for
details).
Similarly, the type parameter <code>S</code> must satisfy its constraint
<code>~[]E</code>. This can be expressed as <code>S ≡<sub>C</sub> ~[]E</code>
where <code>X ≡<sub>C</sub> Y</code> stands for
"<code>X</code> satisfies constraint <code>Y</code>".
These observations lead to a set of two equations
</p>

<pre>
	Slice ≡<sub>A</sub> S      (1)
	S     ≡<sub>C</sub> ~[]E   (2)
</pre>

<p>
which now can be solved for the type parameters <code>S</code> and <code>E</code>.
From (1) a compiler can infer that the type argument for <code>S</code> is <code>Slice</code>.
Similarly, because the underlying type of <code>Slice</code> is <code>[]int</code>
and <code>[]int</code> must match <code>[]E</code> of the constraint,
a compiler can infer that <code>E</code> must be <code>int</code>.
Thus, for these two equations, type inference infers
</p>

<pre>
	S ➞ Slice
	E ➞ int
</pre>

<p>
Given a set of type equations, the type parameters to solve for are
the type parameters of the functions that need to be instantiated
and for which no explicit type arguments is provided.
These type parameters are called <i>bound</i> type parameters.
For instance, in the <code>dedup</code> example above, the type parameters
<code>S</code> and <code>E</code> are bound to <code>dedup</code>.
An argument to a generic function call may be a generic function itself.
The type parameters of that function are included in the set of bound
type parameters.
The types of function arguments may contain type parameters from other
functions (such as a generic function enclosing a function call).
Those type parameters may also appear in type equations but they are
not bound in that context.
Type equations are always solved for the bound type parameters only.
</p>

<p>
Type inference supports calls of generic functions and assignments
of generic functions to (explicitly function-typed) variables.
This includes passing generic functions as arguments to other
(possibly also generic) functions, and returning generic functions
as results.
Type inference operates on a set of equations specific to each of
these cases.
The equations are as follows (type argument lists are omitted for clarity):
</p>

<ul>
<li>
	<p>
	For a function call <code>f(a<sub>0</sub>, a<sub>1</sub>, …)</code> where
	<code>f</code> or a function argument <code>a<sub>i</sub></code> is
	a generic function:
	<br>
	Each pair <code>(a<sub>i</sub>, p<sub>i</sub>)</code> of corresponding
	function arguments and parameters where <code>a<sub>i</sub></code> is not an
	<a href="#Constants">untyped constant</a> yields an equation
	<code>typeof(p<sub>i</sub>) ≡<sub>A</sub> typeof(a<sub>i</sub>)</code>.
	<br>
	If <code>a<sub>i</sub></code> is an untyped constant <code>c<sub>j</sub></code>,
	and <code>typeof(p<sub>i</sub>)</code> is a bound type parameter <code>P<sub>k</sub></code>,
	the pair <code>(c<sub>j</sub>, P<sub>k</sub>)</code> is collected separately from
	the type equations.
	</p>
</li>
<li>
	<p>
	For an assignment <code>v = f</code> of a generic function <code>f</code> to a
	(non-generic) variable <code>v</code> of function type:
	<br>
	<code>typeof(v) ≡<sub>A</sub> typeof(f)</code>.
	</p>
</li>
<li>
	<p>
	For a return statement <code>return …, f, … </code> where <code>f</code> is a
	generic function returned as a result to a (non-generic) result variable
	<code>r</code> of function type:
	<br>
	<code>typeof(r) ≡<sub>A</sub> typeof(f)</code>.
	</p>
</li>
</ul>

<p>
Additionally, each type parameter <code>P<sub>k</sub></code> and corresponding type constraint
<code>C<sub>k</sub></code> yields the type equation
<code>P<sub>k</sub> ≡<sub>C</sub> C<sub>k</sub></code>.
</p>

<p>
Type inference gives precedence to type information obtained from typed operands
before considering untyped constants.
Therefore, inference proceeds in two phases:
</p>

<ol>
<li>
	<p>
	The type equations are solved for the bound
	type parameters using <a href="#Type_unification">type unification</a>.
	If unification fails, type inference fails.
	</p>
</li>
<li>
	<p>
	For each bound type parameter <code>P<sub>k</sub></code> for which no type argument
	has been inferred yet and for which one or more pairs
	<code>(c<sub>j</sub>, P<sub>k</sub>)</code> with that same type parameter
	were collected, determine the <a href="#Constant_expressions">constant kind</a>
	of the constants <code>c<sub>j</sub></code> in all those pairs the same way as for
	<a href="#Constant_expressions">constant expressions</a>.
	The type argument for <code>P<sub>k</sub></code> is the
	<a href="#Constants">default type</a> for the determined constant kind.
	If a constant kind cannot be determined due to conflicting constant kinds,
	type inference fails.
	</p>
</li>
</ol>

<p>
If not all type arguments have been found after these two phases, type inference fails.
</p>

<p>
If the two phases are successful, type inference determined a type argument for each
bound type parameter:
</p>

<pre>
	P<sub>k</sub> ➞ A<sub>k</sub>
</pre>

<p>
A type argument <code>A<sub>k</sub></code> may be a composite type,
containing other bound type parameters <code>P<sub>k</sub></code> as element types
(or even be just another bound type parameter).
In a process of repeated simplification, the bound type parameters in each type
argument are substituted with the respective type arguments for those type
parameters until each type argument is free of bound type parameters.
</p>

<p>
If type arguments contain cyclic references to themselves
through bound type parameters, simplification and thus type
inference fails.
Otherwise, type inference succeeds.
</p>

<h4 id="Type_unification">Type unification</h4>

<p>
Type inference solves type equations through <i>type unification</i>.
Type unification recursively compares the LHS and RHS types of an
equation, where either or both types may be or contain bound type parameters,
and looks for type arguments for those type parameters such that the LHS
and RHS match (become identical or assignment-compatible, depending on
context).
To that effect, type inference maintains a map of bound type parameters
to inferred type arguments; this map is consulted and updated during type unification.
Initially, the bound type parameters are known but the map is empty.
During type unification, if a new type argument <code>A</code> is inferred,
the respective mapping <code>P ➞ A</code> from type parameter to argument
is added to the map.
Conversely, when comparing types, a known type argument
(a type argument for which a map entry already exists)
takes the place of its corresponding type parameter.
As type inference progresses, the map is populated more and more
until all equations have been considered, or until unification fails.
Type inference succeeds if no unification step fails and the map has
an entry for each type parameter.
</p>

<p>
For example, given the type equation with the bound type parameter
<code>P</code>
</p>

<pre>
	[10]struct{ elem P, list []P } ≡<sub>A</sub> [10]struct{ elem string; list []string }
</pre>

<p>
type inference starts with an empty map.
Unification first compares the top-level structure of the LHS and RHS
types.
Both are arrays of the same length; they unify if the element types unify.
Both element types are structs; they unify if they have
the same number of fields with the same names and if the
field types unify.
The type argument for <code>P</code> is not known yet (there is no map entry),
so unifying <code>P</code> with <code>string</code> adds
the mapping <code>P ➞ string</code> to the map.
Unifying the types of the <code>list</code> field requires
unifying <code>[]P</code> and <code>[]string</code> and
thus <code>P</code> and <code>string</code>.
Since the type argument for <code>P</code> is known at this point
(there is a map entry for <code>P</code>), its type argument
<code>string</code> takes the place of <code>P</code>.
And since <code>string</code> is identical to <code>string</code>,
this unification step succeeds as well.
Unification of the LHS and RHS of the equation is now finished.
Type inference succeeds because there is only one type equation,
no unification step failed, and the map is fully populated.
</p>

<p>
Unification uses a combination of <i>exact</i> and <i>loose</i>
unification depending on whether two types have to be
<a href="#Type_identity">identical</a>,
<a href="#Assignability">assignment-compatible</a>, or
only structurally equal.
The respective <a href="#Type_unification_rules">type unification rules</a>
are spelled out in detail in the <a href="#Appendix">Appendix</a>.
</p>

<p>
For an equation of the form <code>X ≡<sub>A</sub> Y</code>,
where <code>X</code> and <code>Y</code> are types involved
in an assignment (including parameter passing and return statements),
the top-level type structures may unify loosely but element types
must unify exactly, matching the rules for assignments.
</p>

<p>
For an equation of the form <code>P ≡<sub>C</sub> C</code>,
where <code>P</code> is a type parameter and <code>C</code>
its corresponding constraint, the unification rules are bit
more complicated:
</p>

<ul>
<li>
	If <code>C</code> has a <a href="#Core_types">core type</a>
	<code>core(C)</code>
	and <code>P</code> has a known type argument <code>A</code>,
	<code>core(C)</code> and <code>A</code> must unify loosely.
	If <code>P</code> does not have a known type argument
	and <code>C</code> contains exactly one type term <code>T</code>
	that is not an underlying (tilde) type, unification adds the
	mapping <code>P ➞ T</code> to the map.
</li>
<li>
	If <code>C</code> does not have a core type
	and <code>P</code> has a known type argument <code>A</code>,
	<code>A</code> must have all methods of <code>C</code>, if any,
	and corresponding method types must unify exactly.
</li>
</ul>

<p>
When solving type equations from type constraints,
solving one equation may infer additional type arguments,
which in turn may enable solving other equations that depend
on those type arguments.
Type inference repeats type unification as long as new type
arguments are inferred.
</p>

<h3 id="Operators">Operators</h3>

<p>
Operators combine operands into expressions.
</p>

<pre class="ebnf">
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&amp;&amp;" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "&lt;" | "&lt;=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "&lt;&lt;" | "&gt;&gt;" | "&amp;" | "&amp;^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&amp;" | "&lt;-" .
</pre>

<p>
Comparisons are discussed <a href="#Comparison_operators">elsewhere</a>.
For other binary operators, the operand types must be <a href="#Type_identity">identical</a>
unless the operation involves shifts or untyped <a href="#Constants">constants</a>.
For operations involving constants only, see the section on
<a href="#Constant_expressions">constant expressions</a>.
</p>

<p>
Except for shift operations, if one operand is an untyped <a href="#Constants">constant</a>
and the other operand is not, the constant is implicitly <a href="#Conversions">converted</a>
to the type of the other operand.
</p>

<p>
The right operand in a shift expression must have <a href="#Numeric_types">integer type</a>
[<a href="#Go_1.13">Go 1.13</a>]
or be an untyped constant <a href="#Representability">representable</a> by a
value of type <code>uint</code>.
If the left operand of a non-constant shift expression is an untyped constant,
it is first implicitly converted to the type it would assume if the shift expression were
replaced by its left operand alone.
</p>

<pre>
var a [1024]byte
var s uint = 33

// The results of the following examples are given for 64-bit ints.
var i = 1&lt;&lt;s                   // 1 has type int
var j int32 = 1&lt;&lt;s             // 1 has type int32; j == 0
var k = uint64(1&lt;&lt;s)           // 1 has type uint64; k == 1&lt;&lt;33
var m int = 1.0&lt;&lt;s             // 1.0 has type int; m == 1&lt;&lt;33
var n = 1.0&lt;&lt;s == j            // 1.0 has type int32; n == true
var o = 1&lt;&lt;s == 2&lt;&lt;s           // 1 and 2 have type int; o == false
var p = 1&lt;&lt;s == 1&lt;&lt;33          // 1 has type int; p == true
var u = 1.0&lt;&lt;s                 // illegal: 1.0 has type float64, cannot shift
var u1 = 1.0&lt;&lt;s != 0           // illegal: 1.0 has type float64, cannot shift
var u2 = 1&lt;&lt;s != 1.0           // illegal: 1 has type float64, cannot shift
var v1 float32 = 1&lt;&lt;s          // illegal: 1 has type float32, cannot shift
var v2 = string(1&lt;&lt;s)          // illegal: 1 is converted to a string, cannot shift
var w int64 = 1.0&lt;&lt;33          // 1.0&lt;&lt;33 is a constant shift expression; w == 1&lt;&lt;33
var x = a[1.0&lt;&lt;s]              // panics: 1.0 has type int, but 1&lt;&lt;33 overflows array bounds
var b = make([]byte, 1.0&lt;&lt;s)   // 1.0 has type int; len(b) == 1&lt;&lt;33

// The results of the following examples are given for 32-bit ints,
// which means the shifts will overflow.
var mm int = 1.0&lt;&lt;s            // 1.0 has type int; mm == 0
var oo = 1&lt;&lt;s == 2&lt;&lt;s          // 1 and 2 have type int; oo == true
var pp = 1&lt;&lt;s == 1&lt;&lt;33         // illegal: 1 has type int, but 1&lt;&lt;33 overflows int
var xx = a[1.0&lt;&lt;s]             // 1.0 has type int; xx == a[0]
var bb = make([]byte, 1.0&lt;&lt;s)  // 1.0 has type int; len(bb) == 0
</pre>

<h4 id="Operator_precedence">Operator precedence</h4>
<p>
Unary operators have the highest precedence.
As the  <code>++</code> and <code>--</code> operators form
statements, not expressions, they fall
outside the operator hierarchy.
As a consequence, statement <code>*p++</code> is the same as <code>(*p)++</code>.
</p>
<p>
There are five precedence levels for binary operators.
Multiplication operators bind strongest, followed by addition
operators, comparison operators, <code>&amp;&amp;</code> (logical AND),
and finally <code>||</code> (logical OR):
</p>

<pre class="grammar">
Precedence    Operator
    5             *  /  %  &lt;&lt;  &gt;&gt;  &amp;  &amp;^
    4             +  -  |  ^
    3             ==  !=  &lt;  &lt;=  &gt;  &gt;=
    2             &amp;&amp;
    1             ||
</pre>

<p>
Binary operators of the same precedence associate from left to right.
For instance, <code>x / y * z</code> is the same as <code>(x / y) * z</code>.
</p>

<pre>
+x                         // x
42 + a - b                 // (42 + a) - b
23 + 3*x[i]                // 23 + (3 * x[i])
x &lt;= f()                   // x &lt;= f()
^a &gt;&gt; b                    // (^a) >> b
f() || g()                 // f() || g()
x == y+1 &amp;&amp; &lt;-chanInt &gt; 0  // (x == (y+1)) && ((<-chanInt) > 0)
</pre>


<h3 id="Arithmetic_operators">Arithmetic operators</h3>
<p>
Arithmetic operators apply to numeric values and yield a result of the same
type as the first operand. The four standard arithmetic operators (<code>+</code>,
<code>-</code>, <code>*</code>, <code>/</code>) apply to
<a href="#Numeric_types">integer</a>, <a href="#Numeric_types">floating-point</a>, and
<a href="#Numeric_types">complex</a> types; <code>+</code> also applies to <a href="#String_types">strings</a>.
The bitwise logical and shift operators apply to integers only.
</p>

<pre class="grammar">
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&amp;    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&amp;^   bit clear (AND NOT)    integers

&lt;&lt;   left shift             integer &lt;&lt; integer &gt;= 0
&gt;&gt;   right shift            integer &gt;&gt; integer &gt;= 0
</pre>

<p>
If the operand type is a <a href="#Type_parameter_declarations">type parameter</a>,
the operator must apply to each type in that type set.
The operands are represented as values of the type argument that the type parameter
is <a href="#Instantiations">instantiated</a> with, and the operation is computed
with the precision of that type argument. For example, given the function:
</p>

<pre>
func dotProduct[F ~float32|~float64](v1, v2 []F) F {
	var s F
	for i, x := range v1 {
		y := v2[i]
		s += x * y
	}
	return s
}
</pre>

<p>
the product <code>x * y</code> and the addition <code>s += x * y</code>
are computed with <code>float32</code> or <code>float64</code> precision,
respectively, depending on the type argument for <code>F</code>.
</p>

<h4 id="Integer_operators">Integer operators</h4>

<p>
For two integer values <code>x</code> and <code>y</code>, the integer quotient
<code>q = x / y</code> and remainder <code>r = x % y</code> satisfy the following
relationships:
</p>

<pre>
x = q*y + r  and  |r| &lt; |y|
</pre>

<p>
with <code>x / y</code> truncated towards zero
(<a href="https://en.wikipedia.org/wiki/Modulo_operation">"truncated division"</a>).
</p>

<pre>
 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2
</pre>

<p>
The one exception to this rule is that if the dividend <code>x</code> is
the most negative value for the int type of <code>x</code>, the quotient
<code>q = x / -1</code> is equal to <code>x</code> (and <code>r = 0</code>)
due to two's-complement <a href="#Integer_overflow">integer overflow</a>:
</p>

<pre>
                         x, q
int8                     -128
int16                  -32768
int32             -2147483648
int64    -9223372036854775808
</pre>

<p>
If the divisor is a <a href="#Constants">constant</a>, it must not be zero.
If the divisor is zero at run time, a <a href="#Run_time_panics">run-time panic</a> occurs.
If the dividend is non-negative and the divisor is a constant power of 2,
the division may be replaced by a right shift, and computing the remainder may
be replaced by a bitwise AND operation:
</p>

<pre>
 x     x / 4     x % 4     x &gt;&gt; 2     x &amp; 3
 11      2         3         2          3
-11     -2        -3        -3          1
</pre>

<p>
The shift operators shift the left operand by the shift count specified by the
right operand, which must be non-negative. If the shift count is negative at run time,
a <a href="#Run_time_panics">run-time panic</a> occurs.
The shift operators implement arithmetic shifts if the left operand is a signed
integer and logical shifts if it is an unsigned integer.
There is no upper limit on the shift count. Shifts behave
as if the left operand is shifted <code>n</code> times by 1 for a shift
count of <code>n</code>.
As a result, <code>x &lt;&lt; 1</code> is the same as <code>x*2</code>
and <code>x &gt;&gt; 1</code> is the same as
<code>x/2</code> but truncated towards negative infinity.
</p>

<p>
For integer operands, the unary operators
<code>+</code>, <code>-</code>, and <code>^</code> are defined as
follows:
</p>

<pre class="grammar">
+x                          is 0 + x
-x    negation              is 0 - x
^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
                                      and  m = -1 for signed x
</pre>


<h4 id="Integer_overflow">Integer overflow</h4>

<p>
For <a href="#Numeric_types">unsigned integer</a> values, the operations <code>+</code>,
<code>-</code>, <code>*</code>, and <code>&lt;&lt;</code> are
computed modulo 2<sup><i>n</i></sup>, where <i>n</i> is the bit width of
the unsigned integer's type.
Loosely speaking, these unsigned integer operations
discard high bits upon overflow, and programs may rely on "wrap around".
</p>

<p>
For signed integers, the operations <code>+</code>,
<code>-</code>, <code>*</code>, <code>/</code>, and <code>&lt;&lt;</code> may legally
overflow and the resulting value exists and is deterministically defined
by the signed integer representation, the operation, and its operands.
Overflow does not cause a <a href="#Run_time_panics">run-time panic</a>.
A compiler may not optimize code under the assumption that overflow does
not occur. For instance, it may not assume that <code>x &lt; x + 1</code> is always true.
</p>

<h4 id="Floating_point_operators">Floating-point operators</h4>

<p>
For floating-point and complex numbers,
<code>+x</code> is the same as <code>x</code>,
while <code>-x</code> is the negation of <code>x</code>.
The result of a floating-point or complex division by zero is not specified beyond the
IEEE-754 standard; whether a <a href="#Run_time_panics">run-time panic</a>
occurs is implementation-specific.
</p>

<p>
An implementation may combine multiple floating-point operations into a single
fused operation, possibly across statements, and produce a result that differs
from the value obtained by executing and rounding the instructions individually.
An explicit <a href="#Numeric_types">floating-point type</a> <a href="#Conversions">conversion</a> rounds to
the precision of the target type, preventing fusion that would discard that rounding.
</p>

<p>
For instance, some architectures provide a "fused multiply and add" (FMA) instruction
that computes <code>x*y + z</code> without rounding the intermediate result <code>x*y</code>.
These examples show when a Go implementation can use that instruction:
</p>

<pre>
// FMA allowed for computing r, because x*y is not explicitly rounded:
r  = x*y + z
r  = z;   r += x*y
t  = x*y; r = t + z
*p = x*y; r = *p + z
r  = x*y + float64(z)

// FMA disallowed for computing r, because it would omit rounding of x*y:
r  = float64(x*y) + z
r  = z; r += float64(x*y)
t  = float64(x*y); r = t + z
</pre>

<h4 id="String_concatenation">String concatenation</h4>

<p>
Strings can be concatenated using the <code>+</code> operator
or the <code>+=</code> assignment operator:
</p>

<pre>
s := "hi" + string(c)
s += " and good bye"
</pre>

<p>
String addition creates a new string by concatenating the operands.
</p>

<h3 id="Comparison_operators">Comparison operators</h3>

<p>
Comparison operators compare two operands and yield an untyped boolean value.
</p>

<pre class="grammar">
==    equal
!=    not equal
&lt;     less
&lt;=    less or equal
&gt;     greater
&gt;=    greater or equal
</pre>

<p>
In any comparison, the first operand
must be <a href="#Assignability">assignable</a>
to the type of the second operand, or vice versa.
</p>
<p>
The equality operators <code>==</code> and <code>!=</code> apply
to operands of <i>comparable</i> types.
The ordering operators <code>&lt;</code>, <code>&lt;=</code>, <code>&gt;</code>, and <code>&gt;=</code>
apply to operands of <i>ordered</i> types.
These terms and the result of the comparisons are defined as follows:
</p>

<ul>
	<li>
	Boolean types are comparable.
	Two boolean values are equal if they are either both
	<code>true</code> or both <code>false</code>.
	</li>

	<li>
	Integer types are comparable and ordered.
	Two integer values are compared in the usual way.
	</li>

	<li>
	Floating-point types are comparable and ordered.
	Two floating-point values are compared as defined by the IEEE-754 standard.
	</li>

	<li>
	Complex types are comparable.
	Two complex values <code>u</code> and <code>v</code> are
	equal if both <code>real(u) == real(v)</code> and
	<code>imag(u) == imag(v)</code>.
	</li>

	<li>
	String types are comparable and ordered.
	Two string values are compared lexically byte-wise.
	</li>

	<li>
	Pointer types are comparable.
	Two pointer values are equal if they point to the same variable or if both have value <code>nil</code>.
	Pointers to distinct <a href="#Size_and_alignment_guarantees">zero-size</a> variables may or may not be equal.
	</li>

	<li>
	Channel types are comparable.
	Two channel values are equal if they were created by the same call to
	<a href="#Making_slices_maps_and_channels"><code>make</code></a>
	or if both have value <code>nil</code>.
	</li>

	<li>
	Interface types that are not type parameters are comparable.
	Two interface values are equal if they have <a href="#Type_identity">identical</a> dynamic types
	and equal dynamic values or if both have value <code>nil</code>.
	</li>

	<li>
	A value <code>x</code> of non-interface type <code>X</code> and
	a value <code>t</code> of interface type <code>T</code> can be compared
	if type <code>X</code> is comparable and
	<code>X</code> <a href="#Implementing_an_interface">implements</a> <code>T</code>.
	They are equal if <code>t</code>'s dynamic type is identical to <code>X</code>
	and <code>t</code>'s dynamic value is equal to <code>x</code>.
	</li>

	<li>
	Struct types are comparable if all their field types are comparable.
	Two struct values are equal if their corresponding
	non-<a href="#Blank_identifier">blank</a> field values are equal.
	The fields are compared in source order, and comparison stops as
	soon as two field values differ (or all fields have been compared).
	</li>

	<li>
	Array types are comparable if their array element types are comparable.
	Two array values are equal if their corresponding element values are equal.
	The elements are compared in ascending index order, and comparison stops
	as soon as two element values differ (or all elements have been compared).
	</li>

	<li>
	Type parameters are comparable if they are strictly comparable (see below).
	</li>
</ul>

<p>
A comparison of two interface values with identical dynamic types
causes a <a href="#Run_time_panics">run-time panic</a> if that type
is not comparable.  This behavior applies not only to direct interface
value comparisons but also when comparing arrays of interface values
or structs with interface-valued fields.
</p>

<p>
Slice, map, and function types are not comparable.
However, as a special case, a slice, map, or function value may
be compared to the predeclared identifier <code>nil</code>.
Comparison of pointer, channel, and interface values to <code>nil</code>
is also allowed and follows from the general rules above.
</p>

<pre>
const c = 3 &lt; 4            // c is the untyped boolean constant true

type MyBool bool
var x, y int
var (
	// The result of a comparison is an untyped boolean.
	// The usual assignment rules apply.
	b3        = x == y // b3 has type bool
	b4 bool   = x == y // b4 has type bool
	b5 MyBool = x == y // b5 has type MyBool
)
</pre>

<p>
A type is <i>strictly comparable</i> if it is comparable and not an interface
type nor composed of interface types.
Specifically:
</p>

<ul>
	<li>
	Boolean, numeric, string, pointer, and channel types are strictly comparable.
	</li>

	<li>
	Struct types are strictly comparable if all their field types are strictly comparable.
	</li>

	<li>
	Array types are strictly comparable if their array element types are strictly comparable.
	</li>

	<li>
	Type parameters are strictly comparable if all types in their type set are strictly comparable.
	</li>
</ul>

<h3 id="Logical_operators">Logical operators</h3>

<p>
Logical operators apply to <a href="#Boolean_types">boolean</a> values
and yield a result of the same type as the operands.
The left operand is evaluated, and then the right if the condition requires it.
</p>

<pre class="grammar">
&amp;&amp;    conditional AND    p &amp;&amp; q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
</pre>


## Address operators

* let's be `x` == operand / type `T`
  * `&x` == address operation
    * -> generates a pointer -- to -- `x` / 
      * type `*T`
  * requirements
    * ⚠️operand MUST be 
      * addressable⚠️
        * == ALLOWED
          * variable,
          * pointer indirection,
          * slice indexing operation,
          * addressable struct operand's field selector
          * addressable array's array indexing operation
      * composite literal⚠️
  * ⚠️if | evaluate `x`, run-time panic -> evaluation of `&x` run-time panic⚠️

* let's be `x` == operand of pointer type `*T`
  * pointer indirection (`*x`) == variable /
    * type `T` 
    * pointed to, -- by -- `x`
  * ⚠️if `x == nil` & you try to evaluate `*x` -> run-time panic⚠️

```go
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic
```

## Receive operator

<p>
For an operand <code>ch</code> whose <a href="#Core_types">core type</a> is a
<a href="#Channel_types">channel</a>,
the value of the receive operation <code>&lt;-ch</code> is the value received
from the channel <code>ch</code>. The channel direction must permit receive operations,
and the type of the receive operation is the element type of the channel.
The expression blocks until a value is available.
Receiving from a <code>nil</code> channel blocks forever.
A receive operation on a <a href="#Close">closed</a> channel can always proceed
immediately, yielding the element type's <a href="#The_zero_value">zero value</a>
after any previously sent values have been received.
</p>

<pre>
v1 := &lt;-ch
v2 = &lt;-ch
f(&lt;-ch)
&lt;-strobe  // wait until clock pulse and discard received value
</pre>

<p>
A receive expression used in an <a href="#Assignment_statements">assignment statement</a> or initialization of the special form
</p>

<pre>
x, ok = &lt;-ch
x, ok := &lt;-ch
var x, ok = &lt;-ch
var x, ok T = &lt;-ch
</pre>

<p>
yields an additional untyped boolean result reporting whether the
communication succeeded. The value of <code>ok</code> is <code>true</code>
if the value received was delivered by a successful send operation to the
channel, or <code>false</code> if it is a zero value generated because the
channel is closed and empty.
</p>


<h3 id="Conversions">Conversions</h3>

<p>
A conversion changes the <a href="#Types">type</a> of an expression
to the type specified by the conversion.
A conversion may appear literally in the source, or it may be <i>implied</i>
by the context in which an expression appears.
</p>

<p>
An <i>explicit</i> conversion is an expression of the form <code>T(x)</code>
where <code>T</code> is a type and <code>x</code> is an expression
that can be converted to type <code>T</code>.
</p>

<pre class="ebnf">
Conversion = Type "(" Expression [ "," ] ")" .
</pre>

<p>
If the type starts with the operator <code>*</code> or <code>&lt;-</code>,
or if the type starts with the keyword <code>func</code>
and has no result list, it must be parenthesized when
necessary to avoid ambiguity:
</p>

<pre>
*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
&lt;-chan int(c)    // same as &lt;-(chan int(c))
(&lt;-chan int)(c)  // c is converted to &lt;-chan int
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)
</pre>

<p>
A <a href="#Constants">constant</a> value <code>x</code> can be converted to
type <code>T</code> if <code>x</code> is <a href="#Representability">representable</a>
by a value of <code>T</code>.
As a special case, an integer constant <code>x</code> can be explicitly converted to a
<a href="#String_types">string type</a> using the
<a href="#Conversions_to_and_from_a_string_type">same rule</a>
as for non-constant <code>x</code>.
</p>

<p>
Converting a constant to a type that is not a <a href="#Type_parameter_declarations">type parameter</a>
yields a typed constant.
</p>

<pre>
uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
float64(-1e-1000)        // 0.0 of type float64
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
myString("foo" + "bar")  // "foobar" of type myString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant
</pre>

<p>
Converting a constant to a type parameter yields a <i>non-constant</i> value of that type,
with the value represented as a value of the type argument that the type parameter
is <a href="#Instantiations">instantiated</a> with.
For example, given the function:
</p>

<pre>
func f[P ~float32|~float64]() {
	… P(1.1) …
}
</pre>

<p>
the conversion <code>P(1.1)</code> results in a non-constant value of type <code>P</code>
and the value <code>1.1</code> is represented as a <code>float32</code> or a <code>float64</code>
depending on the type argument for <code>f</code>.
Accordingly, if <code>f</code> is instantiated with a <code>float32</code> type,
the numeric value of the expression <code>P(1.1) + 1.2</code> will be computed
with the same precision as the corresponding non-constant <code>float32</code>
addition.
</p>

<p>
A non-constant value <code>x</code> can be converted to type <code>T</code>
in any of these cases:
</p>

<ul>
	<li>
	<code>x</code> is <a href="#Assignability">assignable</a>
	to <code>T</code>.
	</li>
	<li>
	ignoring struct tags (see below),
	<code>x</code>'s type and <code>T</code> are not
	<a href="#Type_parameter_declarations">type parameters</a> but have
	<a href="#Type_identity">identical</a> <a href="#Underlying_types">underlying types</a>.
	</li>
	<li>
	ignoring struct tags (see below),
	<code>x</code>'s type and <code>T</code> are pointer types
	that are not <a href="#Types">named types</a>,
	and their pointer base types are not type parameters but
	have identical underlying types.
	</li>
	<li>
	<code>x</code>'s type and <code>T</code> are both integer or floating
	point types.
	</li>
	<li>
	<code>x</code>'s type and <code>T</code> are both complex types.
	</li>
	<li>
	<code>x</code> is an integer or a slice of bytes or runes
	and <code>T</code> is a string type.
	</li>
	<li>
	<code>x</code> is a string and <code>T</code> is a slice of bytes or runes.
	</li>
	<li>
	<code>x</code> is a slice, <code>T</code> is an array [<a href="#Go_1.20">Go 1.20</a>]
	or a pointer to an array [<a href="#Go_1.17">Go 1.17</a>],
	and the slice and array types have <a href="#Type_identity">identical</a> element types.
	</li>
</ul>

<p>
Additionally, if <code>T</code> or <code>x</code>'s type <code>V</code> are type
parameters, <code>x</code>
can also be converted to type <code>T</code> if one of the following conditions applies:
</p>

<ul>
<li>
Both <code>V</code> and <code>T</code> are type parameters and a value of each
type in <code>V</code>'s type set can be converted to each type in <code>T</code>'s
type set.
</li>
<li>
Only <code>V</code> is a type parameter and a value of each
type in <code>V</code>'s type set can be converted to <code>T</code>.
</li>
<li>
Only <code>T</code> is a type parameter and <code>x</code> can be converted to each
type in <code>T</code>'s type set.
</li>
</ul>

<p>
<a href="#Struct_types">Struct tags</a> are ignored when comparing struct types
for identity for the purpose of conversion:
</p>

<pre>
type Person struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var data *struct {
	Name    string `json:"name"`
	Address *struct {
		Street string `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
}

var person = (*Person)(data)  // ignoring tags, the underlying types are identical
</pre>

<p>
Specific rules apply to (non-constant) conversions between numeric types or
to and from a string type.
These conversions may change the representation of <code>x</code>
and incur a run-time cost.
All other conversions only change the type but not the representation
of <code>x</code>.
</p>

<p>
There is no linguistic mechanism to convert between pointers and integers.
The package <a href="#Package_unsafe"><code>unsafe</code></a>
implements this functionality under restricted circumstances.
</p>

<h4>Conversions between numeric types</h4>

<p>
For the conversion of non-constant numeric values, the following rules apply:
</p>

<ol>
<li>
When converting between <a href="#Numeric_types">integer types</a>, if the value is a signed integer, it is
sign extended to implicit infinite precision; otherwise it is zero extended.
It is then truncated to fit in the result type's size.
For example, if <code>v := uint16(0x10F0)</code>, then <code>uint32(int8(v)) == 0xFFFFFFF0</code>.
The conversion always yields a valid value; there is no indication of overflow.
</li>
<li>
When converting a <a href="#Numeric_types">floating-point number</a> to an integer, the fraction is discarded
(truncation towards zero).
</li>
<li>
When converting an integer or floating-point number to a floating-point type,
or a <a href="#Numeric_types">complex number</a> to another complex type, the result value is rounded
to the precision specified by the destination type.
For instance, the value of a variable <code>x</code> of type <code>float32</code>
may be stored using additional precision beyond that of an IEEE-754 32-bit number,
but float32(x) represents the result of rounding <code>x</code>'s value to
32-bit precision. Similarly, <code>x + 0.1</code> may use more than 32 bits
of precision, but <code>float32(x + 0.1)</code> does not.
</li>
</ol>

<p>
In all non-constant conversions involving floating-point or complex values,
if the result type cannot represent the value the conversion
succeeds but the result value is implementation-dependent.
</p>

<h4 id="Conversions_to_and_from_a_string_type">Conversions to and from a string type</h4>

<ol>
<li>
Converting a slice of bytes to a string type yields
a string whose successive bytes are the elements of the slice.

<pre>
string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})   // "hellø"
string([]byte{})                                     // ""
string([]byte(nil))                                  // ""

type bytes []byte
string(bytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})    // "hellø"

type myByte byte
string([]myByte{'w', 'o', 'r', 'l', 'd', '!'})       // "world!"
myString([]myByte{'\xf0', '\x9f', '\x8c', '\x8d'})   // "🌍"
</pre>
</li>

<li>
Converting a slice of runes to a string type yields
a string that is the concatenation of the individual rune values
converted to strings.

<pre>
string([]rune{0x767d, 0x9d6c, 0x7fd4})   // "\u767d\u9d6c\u7fd4" == "白鵬翔"
string([]rune{})                         // ""
string([]rune(nil))                      // ""

type runes []rune
string(runes{0x767d, 0x9d6c, 0x7fd4})    // "\u767d\u9d6c\u7fd4" == "白鵬翔"

type myRune rune
string([]myRune{0x266b, 0x266c})         // "\u266b\u266c" == "♫♬"
myString([]myRune{0x1f30e})              // "\U0001f30e" == "🌎"
</pre>
</li>

<li>
Converting a value of a string type to a slice of bytes type
yields a non-nil slice whose successive elements are the bytes of the string.

<pre>
[]byte("hellø")             // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
[]byte("")                  // []byte{}

bytes("hellø")              // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}

[]myByte("world!")          // []myByte{'w', 'o', 'r', 'l', 'd', '!'}
[]myByte(myString("🌏"))    // []myByte{'\xf0', '\x9f', '\x8c', '\x8f'}
</pre>
</li>

<li>
Converting a value of a string type to a slice of runes type
yields a slice containing the individual Unicode code points of the string.

<pre>
[]rune(myString("白鵬翔"))   // []rune{0x767d, 0x9d6c, 0x7fd4}
[]rune("")                  // []rune{}

runes("白鵬翔")              // []rune{0x767d, 0x9d6c, 0x7fd4}

[]myRune("♫♬")              // []myRune{0x266b, 0x266c}
[]myRune(myString("🌐"))    // []myRune{0x1f310}
</pre>
</li>

<li>
Finally, for historical reasons, an integer value may be converted to a string type.
This form of conversion yields a string containing the (possibly multi-byte) UTF-8
representation of the Unicode code point with the given integer value.
Values outside the range of valid Unicode code points are converted to <code>"\uFFFD"</code>.

<pre>
string('a')          // "a"
string(65)           // "A"
string('\xf8')       // "\u00f8" == "ø" == "\xc3\xb8"
string(-1)           // "\ufffd" == "\xef\xbf\xbd"

type myString string
myString('\u65e5')   // "\u65e5" == "日" == "\xe6\x97\xa5"
</pre>

Note: This form of conversion may eventually be removed from the language.
The <a href="/pkg/cmd/vet"><code>go vet</code></a> tool flags certain
integer-to-string conversions as potential errors.
Library functions such as
<a href="/pkg/unicode/utf8#AppendRune"><code>utf8.AppendRune</code></a> or
<a href="/pkg/unicode/utf8#EncodeRune"><code>utf8.EncodeRune</code></a>
should be used instead.
</li>
</ol>

<h4 id="Conversions_from_slice_to_array_or_array_pointer">Conversions from slice to array or array pointer</h4>

<p>
Converting a slice to an array yields an array containing the elements of the underlying array of the slice.
Similarly, converting a slice to an array pointer yields a pointer to the underlying array of the slice.
In both cases, if the <a href="#Length_and_capacity">length</a> of the slice is less than the length of the array,
a <a href="#Run_time_panics">run-time panic</a> occurs.
</p>

<pre>
s := make([]byte, 2, 4)

a0 := [0]byte(s)
a1 := [1]byte(s[1:])     // a1[0] == s[1]
a2 := [2]byte(s)         // a2[0] == s[0]
a4 := [4]byte(s)         // panics: len([4]byte) > len(s)

s0 := (*[0]byte)(s)      // s0 != nil
s1 := (*[1]byte)(s[1:])  // &amp;s1[0] == &amp;s[1]
s2 := (*[2]byte)(s)      // &amp;s2[0] == &amp;s[0]
s4 := (*[4]byte)(s)      // panics: len([4]byte) > len(s)

var t []string
t0 := [0]string(t)       // ok for nil slice t
t1 := (*[0]string)(t)    // t1 == nil
t2 := (*[1]string)(t)    // panics: len([1]string) > len(t)

u := make([]byte, 0)
u0 := (*[0]byte)(u)      // u0 != nil
</pre>

<h3 id="Constant_expressions">Constant expressions</h3>

<p>
Constant expressions may contain only <a href="#Constants">constant</a>
operands and are evaluated at compile time.
</p>

<p>
Untyped boolean, numeric, and string constants may be used as operands
wherever it is legal to use an operand of boolean, numeric, or string type,
respectively.
</p>

<p>
A constant <a href="#Comparison_operators">comparison</a> always yields
an untyped boolean constant.  If the left operand of a constant
<a href="#Operators">shift expression</a> is an untyped constant, the
result is an integer constant; otherwise it is a constant of the same
type as the left operand, which must be of
<a href="#Numeric_types">integer type</a>.
</p>

<p>
Any other operation on untyped constants results in an untyped constant of the
same kind; that is, a boolean, integer, floating-point, complex, or string
constant.
If the untyped operands of a binary operation (other than a shift) are of
different kinds, the result is of the operand's kind that appears later in this
list: integer, rune, floating-point, complex.
For example, an untyped integer constant divided by an
untyped complex constant yields an untyped complex constant.
</p>

<pre>
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 &lt;&lt; 3.0         // d == 8     (untyped integer constant)
const e = 1.0 &lt;&lt; 3         // e == 8     (untyped integer constant)
const f = int32(1) &lt;&lt; 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) &gt;&gt; 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" &gt; "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)
</pre>

<p>
Applying the built-in function <code>complex</code> to untyped
integer, rune, or floating-point constants yields
an untyped complex constant.
</p>

<pre>
const ic = complex(0, c)   // ic == 3.75i  (untyped complex constant)
const iΘ = complex(0, Θ)   // iΘ == 1i     (type complex128)
</pre>

<p>
Constant expressions are always evaluated exactly; intermediate values and the
constants themselves may require precision significantly larger than supported
by any predeclared type in the language. The following are legal declarations:
</p>

<pre>
const Huge = 1 &lt;&lt; 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
const Four int8 = Huge &gt;&gt; 98  // Four == 4                                (type int8)
</pre>

<p>
The divisor of a constant division or remainder operation must not be zero:
</p>

<pre>
3.14 / 0.0   // illegal: division by zero
</pre>

<p>
The values of <i>typed</i> constants must always be accurately
<a href="#Representability">representable</a> by values
of the constant type. The following constant expressions are illegal:
</p>

<pre>
uint(-1)     // -1 cannot be represented as a uint
int(3.14)    // 3.14 cannot be represented as an int
int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
Four * 100   // product 400 cannot be represented as an int8 (type of Four)
</pre>

<p>
The mask used by the unary bitwise complement operator <code>^</code> matches
the rule for non-constants: the mask is all 1s for unsigned constants
and -1 for signed and untyped constants.
</p>

<pre>
^1         // untyped integer constant, equal to -2
uint8(^1)  // illegal: same as uint8(-2), -2 cannot be represented as a uint8
^uint8(1)  // typed uint8 constant, same as 0xFF ^ uint8(1) = uint8(0xFE)
int8(^1)   // same as int8(-2)
^int8(1)   // same as -1 ^ int8(1) = -2
</pre>

<p>
Implementation restriction: A compiler may use rounding while
computing untyped floating-point or complex constant expressions; see
the implementation restriction in the section
on <a href="#Constants">constants</a>.  This rounding may cause a
floating-point constant expression to be invalid in an integer
context, even if it would be integral when calculated using infinite
precision, and vice versa.
</p>


<h3 id="Order_of_evaluation">Order of evaluation</h3>

<p>
At package level, <a href="#Package_initialization">initialization dependencies</a>
determine the evaluation order of individual initialization expressions in
<a href="#Variable_declarations">variable declarations</a>.
Otherwise, when evaluating the <a href="#Operands">operands</a> of an
expression, assignment, or
<a href="#Return_statements">return statement</a>,
all function calls, method calls,
<a href="#Receive operator">receive operations</a>,
and <a href="#Logical_operators">binary logical operations</a>
are evaluated in lexical left-to-right order.
</p>

<p>
For example, in the (function-local) assignment
</p>
<pre>
y[f()], ok = g(z || h(), i()+x[j()], &lt;-c), k()
</pre>
<p>
the function calls and communication happen in the order
<code>f()</code>, <code>h()</code> (if <code>z</code>
evaluates to false), <code>i()</code>, <code>j()</code>,
<code>&lt;-c</code>, <code>g()</code>, and <code>k()</code>.
However, the order of those events compared to the evaluation
and indexing of <code>x</code> and the evaluation
of <code>y</code> and <code>z</code> is not specified,
except as required lexically. For instance, <code>g</code>
cannot be called before its arguments are evaluated.
</p>

<pre>
a := 1
f := func() int { a++; return a }
x := []int{a, f()}            // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified
m := map[int]int{a: 1, a: 2}  // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified
n := map[int]int{a: f()}      // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified
</pre>

<p>
At package level, initialization dependencies override the left-to-right rule
for individual initialization expressions, but not for operands within each
expression:
</p>

<pre>
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }

// functions u and v are independent of all other variables and functions
</pre>

<p>
The function calls happen in the order
<code>u()</code>, <code>sqr()</code>, <code>v()</code>,
<code>f()</code>, <code>v()</code>, and <code>g()</code>.
</p>

<p>
Floating-point operations within a single expression are evaluated according to
the associativity of the operators.  Explicit parentheses affect the evaluation
by overriding the default associativity.
In the expression <code>x + (y + z)</code> the addition <code>y + z</code>
is performed before adding <code>x</code>.
</p>

<h2 id="Statements">Statements</h2>

<p>
Statements control execution.
</p>

<pre class="ebnf">
Statement =
	Declaration | LabeledStmt | SimpleStmt |
	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
	DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
</pre>

## Terminating statements

<p>
A <i>terminating statement</i> interrupts the regular flow of control in
a <a href="#Blocks">block</a>. The following statements are terminating:
</p>

<ol>
<li>
	A <a href="#Return_statements">"return"</a> or
    	<a href="#Goto_statements">"goto"</a> statement.
	<!-- ul below only for regular layout -->
	<ul> </ul>
</li>

<li>
	A call to the built-in function
	<a href="#Handling_panics"><code>panic</code></a>.
	<!-- ul below only for regular layout -->
	<ul> </ul>
</li>

<li>
	A <a href="#Blocks">block</a> in which the statement list ends in a terminating statement.
	<!-- ul below only for regular layout -->
	<ul> </ul>
</li>

<li>
	An <a href="#If_statements">"if" statement</a> in which:
	<ul>
	<li>the "else" branch is present, and</li>
	<li>both branches are terminating statements.</li>
	</ul>
</li>

<li>
	A <a href="#For_statements">"for" statement</a> in which:
	<ul>
	<li>there are no "break" statements referring to the "for" statement, and</li>
	<li>the loop condition is absent, and</li>
	<li>the "for" statement does not use a range clause.</li>
	</ul>
</li>

<li>
	A <a href="#Switch_statements">"switch" statement</a> in which:
	<ul>
	<li>there are no "break" statements referring to the "switch" statement,</li>
	<li>there is a default case, and</li>
	<li>the statement lists in each case, including the default, end in a terminating
	    statement, or a possibly labeled <a href="#Fallthrough_statements">"fallthrough"
	    statement</a>.</li>
	</ul>
</li>

<li>
	A <a href="#Select_statements">"select" statement</a> in which:
	<ul>
	<li>there are no "break" statements referring to the "select" statement, and</li>
	<li>the statement lists in each case, including the default if present,
	    end in a terminating statement.</li>
	</ul>
</li>

<li>
	A <a href="#Labeled_statements">labeled statement</a> labeling
	a terminating statement.
</li>
</ol>

<p>
All other statements are not terminating.
</p>

<p>
A <a href="#Blocks">statement list</a> ends in a terminating statement if the list
is not empty and its final non-empty statement is terminating.
</p>


<h3 id="Empty_statements">Empty statements</h3>

<p>
The empty statement does nothing.
</p>

<pre class="ebnf">
EmptyStmt = .
</pre>


<h3 id="Labeled_statements">Labeled statements</h3>

<p>
A labeled statement may be the target of a <code>goto</code>,
<code>break</code> or <code>continue</code> statement.
</p>

<pre class="ebnf">
LabeledStmt = Label ":" Statement .
Label       = identifier .
</pre>

<pre>
Error: log.Panic("error encountered")
</pre>


## Expression statements

function and method <a href="#Calls">calls</a> and
<a href="#Receive_operator">receive operations</a>
can appear in statement context. Such statements may be parenthesized.

```go
ExpressionStmt = Expression .
```

* `Expression`
  * ALLOWED 
    * function & method calls
    * receive operations
    * wrap -- with -- `()`
  * ❌NOT ALLOWED❌
    * specific built-in functions

<p>
The following built-in functions are not permitted in statement context:
</p>

<pre>
append cap complex imag len make new real
unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice unsafe.SliceData unsafe.String unsafe.StringData
</pre>

<pre>
h(x+y)
f.Close()
&lt;-ch
(&lt;-ch)
len("foo")  // illegal if len is the built-in function
</pre>


## Send statements

```go
SendStmt = Channel "<-" Expression .
Channel  = Expression .
```

* allows
  * sends a value | channel

* TODO: The channel expression's <a href="#Core_types">core type</a>
must be a <a href="#Channel_types">channel</a>,
the channel direction must permit send operations,
and the type of the value to be sent must be <a href="#Assignability">assignable</a>
to the channel's element type.



<p>
Both the channel and the value expression are evaluated before communication
begins. Communication blocks until the send can proceed.
A send on an unbuffered channel can proceed if a receiver is ready.
A send on a buffered channel can proceed if there is room in the buffer.
A send on a closed channel proceeds by causing a <a href="#Run_time_panics">run-time panic</a>.
A send on a <code>nil</code> channel blocks forever.
</p>

<pre>
ch &lt;- 3  // send value 3 to channel ch
</pre>


## IncDec statements

<p>
The "++" and "--" statements increment or decrement their operands
by the untyped <a href="#Constants">constant</a> <code>1</code>.
As with an assignment, the operand must be <a href="#Address_operators">addressable</a>
or a map index expression.
</p>

<pre class="ebnf">
IncDecStmt = Expression ( "++" | "--" ) .
</pre>

<p>
The following <a href="#Assignment_statements">assignment statements</a> are semantically
equivalent:
</p>

<pre class="grammar">
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
</pre>


## Assignment statements

<p>
An <i>assignment</i> replaces the current value stored in a <a href="#Variables">variable</a>
with a new value specified by an <a href="#Expressions">expression</a>.
An assignment statement may assign a single value to a single variable, or multiple values to a
matching number of variables.
</p>

<pre class="ebnf">
Assignment = ExpressionList assign_op ExpressionList .

assign_op = [ add_op | mul_op ] "=" .
</pre>

<p>
Each left-hand side operand must be <a href="#Address_operators">addressable</a>,
a map index expression, or (for <code>=</code> assignments only) the
<a href="#Blank_identifier">blank identifier</a>.
Operands may be parenthesized.
</p>

<pre>
x = 1
*p = f()
a[i] = 23
(k) = &lt;-ch  // same as: k = &lt;-ch
</pre>

<p>
An <i>assignment operation</i> <code>x</code> <i>op</i><code>=</code>
<code>y</code> where <i>op</i> is a binary <a href="#Arithmetic_operators">arithmetic operator</a>
is equivalent to <code>x</code> <code>=</code> <code>x</code> <i>op</i>
<code>(y)</code> but evaluates <code>x</code>
only once.  The <i>op</i><code>=</code> construct is a single token.
In assignment operations, both the left- and right-hand expression lists
must contain exactly one single-valued expression, and the left-hand
expression must not be the blank identifier.
</p>

<pre>
a[i] &lt;&lt;= 2
i &amp;^= 1&lt;&lt;n
</pre>

<p>
A tuple assignment assigns the individual elements of a multi-valued
operation to a list of variables.  There are two forms.  In the
first, the right hand operand is a single multi-valued expression
such as a function call, a <a href="#Channel_types">channel</a> or
<a href="#Map_types">map</a> operation, or a <a href="#Type_assertions">type assertion</a>.
The number of operands on the left
hand side must match the number of values.  For instance, if
<code>f</code> is a function returning two values,
</p>

<pre>
x, y = f()
</pre>

<p>
assigns the first value to <code>x</code> and the second to <code>y</code>.
In the second form, the number of operands on the left must equal the number
of expressions on the right, each of which must be single-valued, and the
<i>n</i>th expression on the right is assigned to the <i>n</i>th
operand on the left:
</p>

<pre>
one, two, three = '一', '二', '三'
</pre>

<p>
The <a href="#Blank_identifier">blank identifier</a> provides a way to
ignore right-hand side values in an assignment:
</p>

<pre>
_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value
</pre>

<p>
The assignment proceeds in two phases.
First, the operands of <a href="#Index_expressions">index expressions</a>
and <a href="#Address_operators">pointer indirections</a>
(including implicit pointer indirections in <a href="#Selectors">selectors</a>)
on the left and the expressions on the right are all
<a href="#Order_of_evaluation">evaluated in the usual order</a>.
Second, the assignments are carried out in left-to-right order.
</p>

<pre>
a, b = b, a  // exchange a and b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x is []int{3, 5, 3}
</pre>

<p>
In assignments, each value must be <a href="#Assignability">assignable</a>
to the type of the operand to which it is assigned, with the following special cases:
</p>

<ol>
<li>
	Any typed value may be assigned to the blank identifier.
</li>

<li>
	If an untyped constant
	is assigned to a variable of interface type or the blank identifier,
	the constant is first implicitly <a href="#Conversions">converted</a> to its
	 <a href="#Constants">default type</a>.
</li>

<li>
	If an untyped boolean value is assigned to a variable of interface type or
	the blank identifier, it is first implicitly converted to type <code>bool</code>.
</li>
</ol>

## If statements

<p>
"If" statements specify the conditional execution of two branches
according to the value of a boolean expression.  If the expression
evaluates to true, the "if" branch is executed, otherwise, if
present, the "else" branch is executed.
</p>

<pre class="ebnf">
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
</pre>

<pre>
if x &gt; max {
	x = max
}
</pre>

<p>
The expression may be preceded by a simple statement, which
executes before the expression is evaluated.
</p>

<pre>
if x := f(); x &lt; y {
	return x
} else if x &gt; z {
	return z
} else {
	return y
}
</pre>


## Switch statements

<p>
"Switch" statements provide multi-way execution.
An expression or type is compared to the "cases"
inside the "switch" to determine which branch
to execute.
</p>

<pre class="ebnf">
SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .
</pre>

<p>
There are two forms: expression switches and type switches.
In an expression switch, the cases contain expressions that are compared
against the value of the switch expression.
In a type switch, the cases contain types that are compared against the
type of a specially annotated switch expression.
The switch expression is evaluated exactly once in a switch statement.
</p>

<h4 id="Expression_switches">Expression switches</h4>

<p>
In an expression switch,
the switch expression is evaluated and
the case expressions, which need not be constants,
are evaluated left-to-right and top-to-bottom; the first one that equals the
switch expression
triggers execution of the statements of the associated case;
the other cases are skipped.
If no case matches and there is a "default" case,
its statements are executed.
There can be at most one default case and it may appear anywhere in the
"switch" statement.
A missing switch expression is equivalent to the boolean value
<code>true</code>.
</p>

<pre class="ebnf">
ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .
</pre>

<p>
If the switch expression evaluates to an untyped constant, it is first implicitly
<a href="#Conversions">converted</a> to its <a href="#Constants">default type</a>.
The predeclared untyped value <code>nil</code> cannot be used as a switch expression.
The switch expression type must be <a href="#Comparison_operators">comparable</a>.
</p>

<p>
If a case expression is untyped, it is first implicitly <a href="#Conversions">converted</a>
to the type of the switch expression.
For each (possibly converted) case expression <code>x</code> and the value <code>t</code>
of the switch expression, <code>x == t</code> must be a valid <a href="#Comparison_operators">comparison</a>.
</p>

<p>
In other words, the switch expression is treated as if it were used to declare and
initialize a temporary variable <code>t</code> without explicit type; it is that
value of <code>t</code> against which each case expression <code>x</code> is tested
for equality.
</p>

<p>
In a case or default clause, the last non-empty statement
may be a (possibly <a href="#Labeled_statements">labeled</a>)
<a href="#Fallthrough_statements">"fallthrough" statement</a> to
indicate that control should flow from the end of this clause to
the first statement of the next clause.
Otherwise control flows to the end of the "switch" statement.
A "fallthrough" statement may appear as the last statement of all
but the last clause of an expression switch.
</p>

<p>
The switch expression may be preceded by a simple statement, which
executes before the expression is evaluated.
</p>

<pre>
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true"
case x &lt; 0: return -x
default: return x
}

switch {
case x &lt; y: f1()
case x &lt; z: f2()
case x == 4: f3()
}
</pre>

<p>
Implementation restriction: A compiler may disallow multiple case
expressions evaluating to the same constant.
For instance, the current compilers disallow duplicate integer,
floating point, or string constants in case expressions.
</p>

<h4 id="Type_switches">Type switches</h4>

<p>
A type switch compares types rather than values. It is otherwise similar
to an expression switch. It is marked by a special switch expression that
has the form of a <a href="#Type_assertions">type assertion</a>
using the keyword <code>type</code> rather than an actual type:
</p>

<pre>
switch x.(type) {
// cases
}
</pre>

<p>
Cases then match actual types <code>T</code> against the dynamic type of the
expression <code>x</code>. As with type assertions, <code>x</code> must be of
<a href="#Interface_types">interface type</a>, but not a
<a href="#Type_parameter_declarations">type parameter</a>, and each non-interface type
<code>T</code> listed in a case must implement the type of <code>x</code>.
The types listed in the cases of a type switch must all be
<a href="#Type_identity">different</a>.
</p>

<pre class="ebnf">
TypeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
TypeCaseClause  = TypeSwitchCase ":" StatementList .
TypeSwitchCase  = "case" TypeList | "default" .
</pre>

<p>
The TypeSwitchGuard may include a
<a href="#Short_variable_declarations">short variable declaration</a>.
When that form is used, the variable is declared at the end of the
TypeSwitchCase in the <a href="#Blocks">implicit block</a> of each clause.
In clauses with a case listing exactly one type, the variable
has that type; otherwise, the variable has the type of the expression
in the TypeSwitchGuard.
</p>

<p>
Instead of a type, a case may use the predeclared identifier
<a href="#Predeclared_identifiers"><code>nil</code></a>;
that case is selected when the expression in the TypeSwitchGuard
is a <code>nil</code> interface value.
There may be at most one <code>nil</code> case.
</p>

<p>
Given an expression <code>x</code> of type <code>interface{}</code>,
the following type switch:
</p>

<pre>
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
</pre>

<p>
could be rewritten:
</p>

<pre>
v := x  // x is evaluated exactly once
if v == nil {
	i := v                                 // type of i is type of x (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (interface{})
		printString("don't know the type")
	}
}
</pre>

<p>
A <a href="#Type_parameter_declarations">type parameter</a> or a <a href="#Type_declarations">generic type</a>
may be used as a type in a case. If upon <a href="#Instantiations">instantiation</a> that type turns
out to duplicate another entry in the switch, the first matching case is chosen.
</p>

<pre>
func f[P any](x any) int {
	switch x.(type) {
	case P:
		return 0
	case string:
		return 1
	case []P:
		return 2
	case []byte:
		return 3
	default:
		return 4
	}
}

var v1 = f[string]("foo")   // v1 == 0
var v2 = f[byte]([]byte{})  // v2 == 2
</pre>

<p>
The type switch guard may be preceded by a simple statement, which
executes before the guard is evaluated.
</p>

<p>
The "fallthrough" statement is not permitted in a type switch.
</p>

## For statements

<p>
A "for" statement specifies repeated execution of a block. There are three forms:
The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.
</p>

<pre class="ebnf">
ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression .
</pre>

<h4 id="For_condition">For statements with single condition</h4>

<p>
In its simplest form, a "for" statement specifies the repeated execution of
a block as long as a boolean condition evaluates to true.
The condition is evaluated before each iteration.
If the condition is absent, it is equivalent to the boolean value
<code>true</code>.
</p>

<pre>
for a &lt; b {
	a *= 2
}
</pre>

<h4 id="For_clause">For statements with <code>for</code> clause</h4>

<p>
A "for" statement with a ForClause is also controlled by its condition, but
additionally it may specify an <i>init</i>
and a <i>post</i> statement, such as an assignment,
an increment or decrement statement. The init statement may be a
<a href="#Short_variable_declarations">short variable declaration</a>, but the post statement must not.
</p>

<pre class="ebnf">
ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt = SimpleStmt .
PostStmt = SimpleStmt .
</pre>

<pre>
for i := 0; i &lt; 10; i++ {
	f(i)
}
</pre>

<p>
If non-empty, the init statement is executed once before evaluating the
condition for the first iteration;
the post statement is executed after each execution of the block (and
only if the block was executed).
Any element of the ForClause may be empty but the
<a href="#Semicolons">semicolons</a> are
required unless there is only a condition.
If the condition is absent, it is equivalent to the boolean value
<code>true</code>.
</p>

<pre>
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
</pre>

<p>
Each iteration has its own separate declared variable (or variables)
[<a href="#Go_1.22">Go 1.22</a>].
The variable used by the first iteration is declared by the init statement.
The variable used by each subsequent iteration is declared implicitly before
executing the post statement and initialized to the value of the previous
iteration's variable at that moment.
</p>

<pre>
var prints []func()
for i := 0; i < 5; i++ {
	prints = append(prints, func() { println(i) })
	i++
}
for _, p := range prints {
	p()
}
</pre>

<p>
prints
</p>

<pre>
1
3
5
</pre>

<p>
Prior to [<a href="#Go_1.22">Go 1.22</a>], iterations share one set of variables
instead of having their own separate variables.
In that case, the example above prints
</p>

<pre>
6
6
6
</pre>

<h4 id="For_range">For statements with <code>range</code> clause</h4>

<p>
A "for" statement with a "range" clause
iterates through all entries of an array, slice, string or map, values received on
a channel, or integer values from zero to an upper limit [<a href="#Go_1.22">Go 1.22</a>].
For each entry it assigns <i>iteration values</i>
to corresponding <i>iteration variables</i> if present and then executes the block.
</p>

<pre class="ebnf">
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
</pre>

<p>
The expression on the right in the "range" clause is called the <i>range expression</i>,
its <a href="#Core_types">core type</a> must be
an array, pointer to an array, slice, string, map, channel permitting
<a href="#Receive_operator">receive operations</a>, or an integer.
As with an assignment, if present the operands on the left must be
<a href="#Address_operators">addressable</a> or map index expressions; they
denote the iteration variables. If the range expression is a channel or integer,
at most one iteration variable is permitted, otherwise there may be up to two.
If the last iteration variable is the <a href="#Blank_identifier">blank identifier</a>,
the range clause is equivalent to the same clause without that identifier.
</p>

<p>
The range expression <code>x</code> is evaluated once before beginning the loop,
with one exception: if at most one iteration variable is present and
<code>len(x)</code> is <a href="#Length_and_capacity">constant</a>,
the range expression is not evaluated.
</p>

<p>
Function calls on the left are evaluated once per iteration.
For each iteration, iteration values are produced as follows
if the respective iteration variables are present:
</p>

<pre class="grammar">
Range expression                          1st value          2nd value

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, &lt;-chan E       element  e  E
integer         n  integer type           value    i  see below
</pre>

<ol>
<li>
For an array, pointer to array, or slice value <code>a</code>, the index iteration
values are produced in increasing order, starting at element index 0.
If at most one iteration variable is present, the range loop produces
iteration values from 0 up to <code>len(a)-1</code> and does not index into the array
or slice itself. For a <code>nil</code> slice, the number of iterations is 0.
</li>

<li>
For a string value, the "range" clause iterates over the Unicode code points
in the string starting at byte index 0.  On successive iterations, the index value will be the
index of the first byte of successive UTF-8-encoded code points in the string,
and the second value, of type <code>rune</code>, will be the value of
the corresponding code point. If the iteration encounters an invalid
UTF-8 sequence, the second value will be <code>0xFFFD</code>,
the Unicode replacement character, and the next iteration will advance
a single byte in the string.
</li>

<li>
The iteration order over maps is not specified
and is not guaranteed to be the same from one iteration to the next.
If a map entry that has not yet been reached is removed during iteration,
the corresponding iteration value will not be produced. If a map entry is
created during iteration, that entry may be produced during the iteration or
may be skipped. The choice may vary for each entry created and from one
iteration to the next.
If the map is <code>nil</code>, the number of iterations is 0.
</li>

<li>
For channels, the iteration values produced are the successive values sent on
the channel until the channel is <a href="#Close">closed</a>. If the channel
is <code>nil</code>, the range expression blocks forever.
</li>

<li>
For an integer value <code>n</code>, the iteration values 0 through <code>n-1</code>
are produced in increasing order.
If <code>n</code> &lt= 0, the loop does not run any iterations.
</li>
</ol>

<p>
The iteration variables may be declared by the "range" clause using a form of
<a href="#Short_variable_declarations">short variable declaration</a>
(<code>:=</code>).
In this case their <a href="#Declarations_and_scope">scope</a> is the block of the "for" statement
and each iteration has its own new variables [<a href="#Go_1.22">Go 1.22</a>]
(see also <a href="#For_clause">"for" statements with a ForClause</a>).
If the range expression is a (possibly untyped) integer expression <code>n</code>,
the variable has the same type as if it was
<a href="#Variable_declarations">declared</a> with initialization
expression <code>n</code>.
Otherwise, the variables have the types of their respective iteration values.
</p>

<p>
If the iteration variables are not explicitly declared by the "range" clause,
they must be preexisting.
In this case, the iteration values are assigned to the respective variables
as in an <a href="#Assignment_statements">assignment statement</a>.
If the range expression is a (possibly untyped) integer expression <code>n</code>,
<code>n</code> too must be <a href="#Assignability">assignable</a> to the iteration variable;
if there is no iteration variable, <code>n</code> must be assignable to <code>int</code>.
</p>

<pre>
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a is never evaluated; len(testdata.a) is constant
	// i ranges from 0 to 6
	f(i)
}

var a [10]string
for i, s := range a {
	// type of i is int
	// type of s is string
	// s == a[i]
	g(i, s)
}

var key string
var val interface{}  // element type of m is assignable to val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// empty a channel
for range ch {}

// call f(0), f(1), ... f(9)
for i := range 10 {
	// type of i is int (default type for untyped constant 10)
	f(i)
}

// invalid: 256 cannot be assigned to uint8
var u uint8
for u = range 256 {
}
</pre>


## Go statements

```go
GoStmt = "go" Expression .
```
* `Expression`
  * requirements
    * function call OR
    * method call
    * NOT wrap -- with -- `()`
  * ❌NOT ALLOWED❌
    * specific built-in functions
  * ⚠️ALTHOUGH Expression returns something -> discarded⚠️

* execution |
  * 💡independent concurrent thread OR💡
  * ⚠️SAME address space⚠️

* program execution
  * ❌NOT wait / goroutine is completed❌
  * 👀function's value & parameters are evaluated | main goroutine👀
    * NOT | NEW goroutine

## Select statements

<p>
A "select" statement chooses which of a set of possible
<a href="#Send_statements">send</a> or
<a href="#Receive_operator">receive</a>
operations will proceed.
It looks similar to a
<a href="#Switch_statements">"switch"</a> statement but with the
cases all referring to communication operations.
</p>

<pre class="ebnf">
SelectStmt = "select" "{" { CommClause } "}" .
CommClause = CommCase ":" StatementList .
CommCase   = "case" ( SendStmt | RecvStmt ) | "default" .
RecvStmt   = [ ExpressionList "=" | IdentifierList ":=" ] RecvExpr .
RecvExpr   = Expression .
</pre>

<p>
A case with a RecvStmt may assign the result of a RecvExpr to one or
two variables, which may be declared using a
<a href="#Short_variable_declarations">short variable declaration</a>.
The RecvExpr must be a (possibly parenthesized) receive operation.
There can be at most one default case and it may appear anywhere
in the list of cases.
</p>

<p>
Execution of a "select" statement proceeds in several steps:
</p>

<ol>
<li>
For all the cases in the statement, the channel operands of receive operations
and the channel and right-hand-side expressions of send statements are
evaluated exactly once, in source order, upon entering the "select" statement.
The result is a set of channels to receive from or send to,
and the corresponding values to send.
Any side effects in that evaluation will occur irrespective of which (if any)
communication operation is selected to proceed.
Expressions on the left-hand side of a RecvStmt with a short variable declaration
or assignment are not yet evaluated.
</li>

<li>
If one or more of the communications can proceed,
a single one that can proceed is chosen via a uniform pseudo-random selection.
Otherwise, if there is a default case, that case is chosen.
If there is no default case, the "select" statement blocks until
at least one of the communications can proceed.
</li>

<li>
Unless the selected case is the default case, the respective communication
operation is executed.
</li>

<li>
If the selected case is a RecvStmt with a short variable declaration or
an assignment, the left-hand side expressions are evaluated and the
received value (or values) are assigned.
</li>

<li>
The statement list of the selected case is executed.
</li>
</ol>

<p>
Since communication on <code>nil</code> channels can never proceed,
a select with only <code>nil</code> channels and no default case blocks forever.
</p>

<pre>
var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = &lt;-c1:
	print("received ", i1, " from c1\n")
case c2 &lt;- i2:
	print("sent ", i2, " to c2\n")
case i3, ok := (&lt;-c3):  // same as: i3, ok := &lt;-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}
case a[f()] = &lt;-c4:
	// same as:
	// case t := &lt;-c4
	//	a[f()] = t
default:
	print("no communication\n")
}

for {  // send random sequence of bits to c
	select {
	case c &lt;- 0:  // note: no statement, no fallthrough, no folding of cases
	case c &lt;- 1:
	}
}

select {}  // block forever
</pre>


## Return statements

<p>
A "return" statement in a function <code>F</code> terminates the execution
of <code>F</code>, and optionally provides one or more result values.
Any functions <a href="#Defer_statements">deferred</a> by <code>F</code>
are executed before <code>F</code> returns to its caller.
</p>

<pre class="ebnf">
ReturnStmt = "return" [ ExpressionList ] .
</pre>

<p>
In a function without a result type, a "return" statement must not
specify any result values.
</p>
<pre>
func noResult() {
	return
}
</pre>

<p>
There are three ways to return values from a function with a result
type:
</p>

<ol>
	<li>The return value or values may be explicitly listed
		in the "return" statement. Each expression must be single-valued
		and <a href="#Assignability">assignable</a>
		to the corresponding element of the function's result type.
<pre>
func simpleF() int {
	return 2
}

func complexF1() (re float64, im float64) {
	return -7.0, -4.0
}
</pre>
	</li>
	<li>The expression list in the "return" statement may be a single
		call to a multi-valued function. The effect is as if each value
		returned from that function were assigned to a temporary
		variable with the type of the respective value, followed by a
		"return" statement listing these variables, at which point the
		rules of the previous case apply.
<pre>
func complexF2() (re float64, im float64) {
	return complexF1()
}
</pre>
	</li>
	<li>The expression list may be empty if the function's result
		type specifies names for its <a href="#Function_types">result parameters</a>.
		The result parameters act as ordinary local variables
		and the function may assign values to them as necessary.
		The "return" statement returns the values of these variables.
<pre>
func complexF3() (re float64, im float64) {
	re = 7.0
	im = 4.0
	return
}

func (devnull) Write(p []byte) (n int, _ error) {
	n = len(p)
	return
}
</pre>
	</li>
</ol>

<p>
Regardless of how they are declared, all the result values are initialized to
the <a href="#The_zero_value">zero values</a> for their type upon entry to the
function. A "return" statement that specifies results sets the result parameters before
any deferred functions are executed.
</p>

<p>
Implementation restriction: A compiler may disallow an empty expression list
in a "return" statement if a different entity (constant, type, or variable)
with the same name as a result parameter is in
<a href="#Declarations_and_scope">scope</a> at the place of the return.
</p>

<pre>
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // invalid return statement: err is shadowed
	}
	return
}
</pre>

<h3 id="Break_statements">Break statements</h3>

<p>
A "break" statement terminates execution of the innermost
<a href="#For_statements">"for"</a>,
<a href="#Switch_statements">"switch"</a>, or
<a href="#Select_statements">"select"</a> statement
within the same function.
</p>

<pre class="ebnf">
BreakStmt = "break" [ Label ] .
</pre>

<p>
If there is a label, it must be that of an enclosing
"for", "switch", or "select" statement,
and that is the one whose execution terminates.
</p>

<pre>
OuterLoop:
	for i = 0; i &lt; n; i++ {
		for j = 0; j &lt; m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
</pre>

<h3 id="Continue_statements">Continue statements</h3>

<p>
A "continue" statement begins the next iteration of the
innermost enclosing <a href="#For_statements">"for" loop</a>
by advancing control to the end of the loop block.
The "for" loop must be within the same function.
</p>

<pre class="ebnf">
ContinueStmt = "continue" [ Label ] .
</pre>

<p>
If there is a label, it must be that of an enclosing
"for" statement, and that is the one whose execution
advances.
</p>

<pre>
RowLoop:
	for y, row := range rows {
		for x, data := range row {
			if data == endOfRow {
				continue RowLoop
			}
			row[x] = data + bias(x, y)
		}
	}
</pre>

<h3 id="Goto_statements">Goto statements</h3>

<p>
A "goto" statement transfers control to the statement with the corresponding label
within the same function.
</p>

<pre class="ebnf">
GotoStmt = "goto" Label .
</pre>

<pre>
goto Error
</pre>

<p>
Executing the "goto" statement must not cause any variables to come into
<a href="#Declarations_and_scope">scope</a> that were not already in scope at the point of the goto.
For instance, this example:
</p>

<pre>
	goto L  // BAD
	v := 3
L:
</pre>

<p>
is erroneous because the jump to label <code>L</code> skips
the creation of <code>v</code>.
</p>

<p>
A "goto" statement outside a <a href="#Blocks">block</a> cannot jump to a label inside that block.
For instance, this example:
</p>

<pre>
if n%2 == 1 {
	goto L1
}
for n &gt; 0 {
	f()
	n--
L1:
	f()
	n--
}
</pre>

<p>
is erroneous because the label <code>L1</code> is inside
the "for" statement's block but the <code>goto</code> is not.
</p>

<h3 id="Fallthrough_statements">Fallthrough statements</h3>

<p>
A "fallthrough" statement transfers control to the first statement of the
next case clause in an <a href="#Expression_switches">expression "switch" statement</a>.
It may be used only as the final non-empty statement in such a clause.
</p>

<pre class="ebnf">
FallthroughStmt = "fallthrough" .
</pre>


## Defer statements

```go
DeferStmt = "defer" Expression .
```
* `Expression` 
  * ALLOWED
    * function call
    * method call
  * ❌NOT ALLOWED❌
    * wrap with `()`

* "defer" statement 
  * allows
    * | invoke a function,
      * 👀function's execution is deferred -- TILL -- the ⚠️surrounding function returns⚠️👀

* 👀ways / surrounding function returns👀
  * surrounding function
    * executes a [return statement](#return-statements) OR
    * reached its [function body](#function-declarations) end OR
  * corresponding goroutine [panicking](#handling-panics)

* TODO: 
<p>
Calls of built-in functions are restricted as for
<a href="#Expression_statements">expression statements</a>.
</p>

<p>
Each time a "defer" statement
executes, the function value and parameters to the call are
<a href="#Calls">evaluated as usual</a>
and saved anew but the actual function is not invoked.
Instead, deferred functions are invoked immediately before
the surrounding function returns, in the reverse order
they were deferred. That is, if the surrounding function
returns through an explicit <a href="#Return_statements">return statement</a>,
deferred functions are executed <i>after</i> any result parameters are set
by that return statement but <i>before</i> the function returns to its caller.
If a deferred function value evaluates
to <code>nil</code>, execution <a href="#Handling_panics">panics</a>
when the function is invoked, not when the "defer" statement is executed.
</p>

<p>
For instance, if the deferred function is
a <a href="#Function_literals">function literal</a> and the surrounding
function has <a href="#Function_types">named result parameters</a> that
are in scope within the literal, the deferred function may access and modify
the result parameters before they are returned.
If the deferred function has any return values, they are discarded when
the function completes.
(See also the section on <a href="#Handling_panics">handling panics</a>.)
</p>

<pre>
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns

// prints 3 2 1 0 before surrounding function returns
for i := 0; i &lt;= 3; i++ {
	defer fmt.Print(i)
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
</pre>

# Built-in functions

<p>
Built-in functions are
<a href="#Predeclared_identifiers">predeclared</a>.
They are called like any other function but some of them
accept a type instead of an expression as the first argument.
</p>

<p>
The built-in functions do not have standard Go types,
so they can only appear in <a href="#Calls">call expressions</a>;
they cannot be used as function values.
</p>


<h3 id="Appending_and_copying_slices">Appending to and copying slices</h3>

<p>
The built-in functions <code>append</code> and <code>copy</code> assist in
common slice operations.
For both functions, the result is independent of whether the memory referenced
by the arguments overlaps.
</p>

<p>
The <a href="#Function_types">variadic</a> function <code>append</code>
appends zero or more values <code>x</code> to a slice <code>s</code>
and returns the resulting slice of the same type as <code>s</code>.
The <a href="#Core_types">core type</a> of <code>s</code> must be a slice
of type <code>[]E</code>.
The values <code>x</code> are passed to a parameter of type <code>...E</code>
and the respective <a href="#Passing_arguments_to_..._parameters">parameter
passing rules</a> apply.
As a special case, if the core type of <code>s</code> is <code>[]byte</code>,
<code>append</code> also accepts a second argument with core type
<a href="#Core_types"><code>bytestring</code></a> followed by <code>...</code>.
This form appends the bytes of the byte slice or string.
</p>

<pre class="grammar">
append(s S, x ...E) S  // core type of S is []E
</pre>

<p>
If the capacity of <code>s</code> is not large enough to fit the additional
values, <code>append</code> <a href="#Allocation">allocates</a> a new, sufficiently large underlying
array that fits both the existing slice elements and the additional values.
Otherwise, <code>append</code> re-uses the underlying array.
</p>

<pre>
s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 is []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 is []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 is []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 is []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")   //                             t is []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // append string contents      b is []byte{'b', 'a', 'r' }
</pre>

<p>
The function <code>copy</code> copies slice elements from
a source <code>src</code> to a destination <code>dst</code> and returns the
number of elements copied.
The <a href="#Core_types">core types</a> of both arguments must be slices
with <a href="#Type_identity">identical</a> element type.
The number of elements copied is the minimum of
<code>len(src)</code> and <code>len(dst)</code>.
As a special case, if the destination's core type is <code>[]byte</code>,
<code>copy</code> also accepts a source argument with core type
<a href="#Core_types"><code>bytestring</code></a>.
This form copies the bytes from the byte slice or string into the byte slice.
</p>

<pre class="grammar">
copy(dst, src []T) int
copy(dst []byte, src string) int
</pre>

<p>
Examples:
</p>

<pre>
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s is []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s is []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b is []byte("Hello")
</pre>


<h3 id="Clear">Clear</h3>

<p>
The built-in function <code>clear</code> takes an argument of <a href="#Map_types">map</a>,
<a href="#Slice_types">slice</a>, or <a href="#Type_parameter_declarations">type parameter</a> type,
and deletes or zeroes out all elements
[<a href="#Go_1.21">Go 1.21</a>].
</p>

<pre class="grammar">
Call        Argument type     Result

clear(m)    map[K]T           deletes all entries, resulting in an
                              empty map (len(m) == 0)

clear(s)    []T               sets all elements up to the length of
                              <code>s</code> to the zero value of T

clear(t)    type parameter    see below
</pre>

<p>
If the type of the argument to <code>clear</code> is a
<a href="#Type_parameter_declarations">type parameter</a>,
all types in its type set must be maps or slices, and <code>clear</code>
performs the operation corresponding to the actual type argument.
</p>

<p>
If the map or slice is <code>nil</code>, <code>clear</code> is a no-op.
</p>


<h3 id="Close">Close</h3>

<p>
For an argument <code>ch</code> with a <a href="#Core_types">core type</a>
that is a <a href="#Channel_types">channel</a>, the built-in function <code>close</code>
records that no more values will be sent on the channel.
It is an error if <code>ch</code> is a receive-only channel.
Sending to or closing a closed channel causes a <a href="#Run_time_panics">run-time panic</a>.
Closing the nil channel also causes a <a href="#Run_time_panics">run-time panic</a>.
After calling <code>close</code>, and after any previously
sent values have been received, receive operations will return
the zero value for the channel's type without blocking.
The multi-valued <a href="#Receive_operator">receive operation</a>
returns a received value along with an indication of whether the channel is closed.
</p>


<h3 id="Complex_numbers">Manipulating complex numbers</h3>

<p>
Three functions assemble and disassemble complex numbers.
The built-in function <code>complex</code> constructs a complex
value from a floating-point real and imaginary part, while
<code>real</code> and <code>imag</code>
extract the real and imaginary parts of a complex value.
</p>

<pre class="grammar">
complex(realPart, imaginaryPart floatT) complexT
real(complexT) floatT
imag(complexT) floatT
</pre>

<p>
The type of the arguments and return value correspond.
For <code>complex</code>, the two arguments must be of the same
<a href="#Numeric_types">floating-point type</a> and the return type is the
<a href="#Numeric_types">complex type</a>
with the corresponding floating-point constituents:
<code>complex64</code> for <code>float32</code> arguments, and
<code>complex128</code> for <code>float64</code> arguments.
If one of the arguments evaluates to an untyped constant, it is first implicitly
<a href="#Conversions">converted</a> to the type of the other argument.
If both arguments evaluate to untyped constants, they must be non-complex
numbers or their imaginary parts must be zero, and the return value of
the function is an untyped complex constant.
</p>

<p>
For <code>real</code> and <code>imag</code>, the argument must be
of complex type, and the return type is the corresponding floating-point
type: <code>float32</code> for a <code>complex64</code> argument, and
<code>float64</code> for a <code>complex128</code> argument.
If the argument evaluates to an untyped constant, it must be a number,
and the return value of the function is an untyped floating-point constant.
</p>

<p>
The <code>real</code> and <code>imag</code> functions together form the inverse of
<code>complex</code>, so for a value <code>z</code> of a complex type <code>Z</code>,
<code>z&nbsp;==&nbsp;Z(complex(real(z),&nbsp;imag(z)))</code>.
</p>

<p>
If the operands of these functions are all constants, the return
value is a constant.
</p>

<pre>
var a = complex(2, -2)             // complex128
const b = complex(1.0, -1.4)       // untyped complex constant 1 - 1.4i
x := float32(math.Cos(math.Pi/2))  // float32
var c64 = complex(5, -x)           // complex64
var s int = complex(1, 0)          // untyped complex constant 1 + 0i can be converted to int
_ = complex(1, 2&lt;&lt;s)               // illegal: 2 assumes floating-point type, cannot shift
var rl = real(c64)                 // float32
var im = imag(a)                   // float64
const c = imag(b)                  // untyped constant -1.4
_ = imag(3 &lt;&lt; s)                   // illegal: 3 assumes complex type, cannot shift
</pre>

<p>
Arguments of type parameter type are not permitted.
</p>


<h3 id="Deletion_of_map_elements">Deletion of map elements</h3>

<p>
The built-in function <code>delete</code> removes the element with key
<code>k</code> from a <a href="#Map_types">map</a> <code>m</code>. The
value <code>k</code> must be <a href="#Assignability">assignable</a>
to the key type of <code>m</code>.
</p>

<pre class="grammar">
delete(m, k)  // remove element m[k] from map m
</pre>

<p>
If the type of <code>m</code> is a <a href="#Type_parameter_declarations">type parameter</a>,
all types in that type set must be maps, and they must all have identical key types.
</p>

<p>
If the map <code>m</code> is <code>nil</code> or the element <code>m[k]</code>
does not exist, <code>delete</code> is a no-op.
</p>


<h3 id="Length_and_capacity">Length and capacity</h3>

<p>
The built-in functions <code>len</code> and <code>cap</code> take arguments
of various types and return a result of type <code>int</code>.
The implementation guarantees that the result always fits into an <code>int</code>.
</p>

<pre class="grammar">
Call      Argument type    Result

len(s)    string type      string length in bytes
          [n]T, *[n]T      array length (== n)
          []T              slice length
          map[K]T          map length (number of defined keys)
          chan T           number of elements queued in channel buffer
          type parameter   see below

cap(s)    [n]T, *[n]T      array length (== n)
          []T              slice capacity
          chan T           channel buffer capacity
          type parameter   see below
</pre>

<p>
If the argument type is a <a href="#Type_parameter_declarations">type parameter</a> <code>P</code>,
the call <code>len(e)</code> (or <code>cap(e)</code> respectively) must be valid for
each type in <code>P</code>'s type set.
The result is the length (or capacity, respectively) of the argument whose type
corresponds to the type argument with which <code>P</code> was
<a href="#Instantiations">instantiated</a>.
</p>

<p>
The capacity of a slice is the number of elements for which there is
space allocated in the underlying array.
At any time the following relationship holds:
</p>

<pre>
0 &lt;= len(s) &lt;= cap(s)
</pre>

<p>
The length of a <code>nil</code> slice, map or channel is 0.
The capacity of a <code>nil</code> slice or channel is 0.
</p>

<p>
The expression <code>len(s)</code> is <a href="#Constants">constant</a> if
<code>s</code> is a string constant. The expressions <code>len(s)</code> and
<code>cap(s)</code> are constants if the type of <code>s</code> is an array
or pointer to an array and the expression <code>s</code> does not contain
<a href="#Receive_operator">channel receives</a> or (non-constant)
<a href="#Calls">function calls</a>; in this case <code>s</code> is not evaluated.
Otherwise, invocations of <code>len</code> and <code>cap</code> are not
constant and <code>s</code> is evaluated.
</p>

<pre>
const (
	c1 = imag(2i)                    // imag(2i) = 2.0 is a constant
	c2 = len([10]float64{2})         // [10]float64{2} contains no function calls
	c3 = len([10]float64{c1})        // [10]float64{c1} contains no function calls
	c4 = len([10]float64{imag(2i)})  // imag(2i) is a constant and no function call is issued
	c5 = len([10]float64{imag(z)})   // invalid: imag(z) is a (non-constant) function call
)
var z complex128
</pre>


<h3 id="Making_slices_maps_and_channels">Making slices, maps and channels</h3>

<p>
The built-in function <code>make</code> takes a type <code>T</code>,
optionally followed by a type-specific list of expressions.
The <a href="#Core_types">core type</a> of <code>T</code> must
be a slice, map or channel.
It returns a value of type <code>T</code> (not <code>*T</code>).
The memory is initialized as described in the section on
<a href="#The_zero_value">initial values</a>.
</p>

<pre class="grammar">
Call             Core type    Result

make(T, n)       slice        slice of type T with length n and capacity n
make(T, n, m)    slice        slice of type T with length n and capacity m

make(T)          map          map of type T
make(T, n)       map          map of type T with initial space for approximately n elements

make(T)          channel      unbuffered channel of type T
make(T, n)       channel      buffered channel of type T, buffer size n
</pre>

<p>
Each of the size arguments <code>n</code> and <code>m</code> must be of <a href="#Numeric_types">integer type</a>,
have a <a href="#Interface_types">type set</a> containing only integer types,
or be an untyped <a href="#Constants">constant</a>.
A constant size argument must be non-negative and <a href="#Representability">representable</a>
by a value of type <code>int</code>; if it is an untyped constant it is given type <code>int</code>.
If both <code>n</code> and <code>m</code> are provided and are constant, then
<code>n</code> must be no larger than <code>m</code>.
For slices and channels, if <code>n</code> is negative or larger than <code>m</code> at run time,
a <a href="#Run_time_panics">run-time panic</a> occurs.
</p>

<pre>
s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1&lt;&lt;63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for approximately 100 elements
</pre>

<p>
Calling <code>make</code> with a map type and size hint <code>n</code> will
create a map with initial space to hold <code>n</code> map elements.
The precise behavior is implementation-dependent.
</p>


<h3 id="Min_and_max">Min and max</h3>

<p>
The built-in functions <code>min</code> and <code>max</code> compute the
smallest&mdash;or largest, respectively&mdash;value of a fixed number of
arguments of <a href="#Comparison_operators">ordered types</a>.
There must be at least one argument
[<a href="#Go_1.21">Go 1.21</a>].
</p>

<p>
The same type rules as for <a href="#Operators">operators</a> apply:
for <a href="#Comparison_operators">ordered</a> arguments <code>x</code> and
<code>y</code>, <code>min(x, y)</code> is valid if <code>x + y</code> is valid,
and the type of <code>min(x, y)</code> is the type of <code>x + y</code>
(and similarly for <code>max</code>).
If all arguments are constant, the result is constant.
</p>

<pre>
var x, y int
m := min(x)                 // m == x
m := min(x, y)              // m is the smaller of x and y
m := max(x, y, 10)          // m is the larger of x and y but at least 10
c := max(1, 2.0, 10)        // c == 10.0 (floating-point kind)
f := max(0, float32(x))     // type of f is float32
var s []string
_ = min(s...)               // invalid: slice arguments are not permitted
t := max("", "foo", "bar")  // t == "foo" (string kind)
</pre>

<p>
For numeric arguments, assuming all NaNs are equal, <code>min</code> and <code>max</code> are
commutative and associative:
</p>

<pre>
min(x, y)    == min(y, x)
min(x, y, z) == min(min(x, y), z) == min(x, min(y, z))
</pre>

<p>
For floating-point arguments negative zero, NaN, and infinity the following rules apply:
</p>

<pre>
   x        y    min(x, y)    max(x, y)

  -0.0    0.0         -0.0          0.0    // negative zero is smaller than (non-negative) zero
  -Inf      y         -Inf            y    // negative infinity is smaller than any other number
  +Inf      y            y         +Inf    // positive infinity is larger than any other number
   NaN      y          NaN          NaN    // if any argument is a NaN, the result is a NaN
</pre>

<p>
For string arguments the result for <code>min</code> is the first argument
with the smallest (or for <code>max</code>, largest) value,
compared lexically byte-wise:
</p>

<pre>
min(x, y)    == if x <= y then x else y
min(x, y, z) == min(min(x, y), z)
</pre>

## Allocation

* 💡-- via -- `new` 💡

* `new(someType)`
  * == built-in function /
    * | run time, 
      * allocates storage -- for a -- [variable](#variables) of `someType`
      * initializes it -- by -- [zero-value](#the-zero-value)
      * returns a value / type == `*someType` / [point](#pointer-types) -- to -- it

* see [zero value](#the-zero-value)

## Handling panics

* built-in functions,
  ```
  func panic(interface{})
  func recover() interface{}
  ```
  * allows
    * assist | report & handle 
      * [run-time panics](#run-time-panics)
      * program-defined error conditions

* 👀if panic happens -> terminates the execution of the function👀

* panicking
  * == termination sequence
    * _Example:_ [handling-panic-terminationsequence](examples/handling-panic-terminationsequence.go)

* `recover()`
  * allows 
    * a program can manage panicking goroutine's behaviour
      * == keep on goroutine's execution
  * if the goroutine is NOT panicking -> return `nil`

## Bootstrapping

<p>
Current implementations provide several built-in functions useful during
bootstrapping. These functions are documented for completeness but are not
guaranteed to stay in the language. They do not return a result.
</p>

<pre class="grammar">
Function   Behavior

print      prints all arguments; formatting of arguments is implementation-specific
println    like print but prints spaces between arguments and a newline at the end
</pre>

<p>
Implementation restriction: <code>print</code> and <code>println</code> need not
accept arbitrary argument types, but printing of boolean, numeric, and string
<a href="#Types">types</a> must be supported.
</p>


# Packages

<p>
Go programs are constructed by linking together <i>packages</i>.
A package in turn is constructed from one or more source files
that together declare constants, types, variables and functions
belonging to the package and which are accessible in all files
of the same package. Those elements may be
<a href="#Exported_identifiers">exported</a> and used in another package.
</p>

<h3 id="Source_file_organization">Source file organization</h3>

<p>
Each source file consists of a package clause defining the package
to which it belongs, followed by a possibly empty set of import
declarations that declare packages whose contents it wishes to use,
followed by a possibly empty set of declarations of functions,
types, variables, and constants.
</p>

<pre class="ebnf">
SourceFile       = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" } .
</pre>

<h3 id="Package_clause">Package clause</h3>

<p>
A package clause begins each source file and defines the package
to which the file belongs.
</p>

<pre class="ebnf">
PackageClause  = "package" PackageName .
PackageName    = identifier .
</pre>

<p>
The PackageName must not be the <a href="#Blank_identifier">blank identifier</a>.
</p>

<pre>
package math
</pre>

<p>
A set of files sharing the same PackageName form the implementation of a package.
An implementation may require that all source files for a package inhabit the same directory.
</p>

## Import declarations

<p>
An import declaration states that the source file containing the declaration
depends on functionality of the <i>imported</i> package
(<a href="#Program_initialization_and_execution">§Program initialization and execution</a>)
and enables access to <a href="#Exported_identifiers">exported</a> identifiers
of that package.
The import names an identifier (PackageName) to be used for access and an ImportPath
that specifies the package to be imported.
</p>

<pre class="ebnf">
ImportDecl       = "import" ( ImportSpec | "(" { ImportSpec ";" } ")" ) .
ImportSpec       = [ "." | PackageName ] ImportPath .
ImportPath       = string_lit .
</pre>

<p>
The PackageName is used in <a href="#Qualified_identifiers">qualified identifiers</a>
to access exported identifiers of the package within the importing source file.
It is declared in the <a href="#Blocks">file block</a>.
If the PackageName is omitted, it defaults to the identifier specified in the
<a href="#Package_clause">package clause</a> of the imported package.
If an explicit period (<code>.</code>) appears instead of a name, all the
package's exported identifiers declared in that package's
<a href="#Blocks">package block</a> will be declared in the importing source
file's file block and must be accessed without a qualifier.
</p>

<p>
The interpretation of the ImportPath is implementation-dependent but
it is typically a substring of the full file name of the compiled
package and may be relative to a repository of installed packages.
</p>

<p>
Implementation restriction: A compiler may restrict ImportPaths to
non-empty strings using only characters belonging to
<a href="https://www.unicode.org/versions/Unicode6.3.0/">Unicode's</a>
L, M, N, P, and S general categories (the Graphic characters without
spaces) and may also exclude the characters
<code>!"#$%&amp;'()*,:;&lt;=&gt;?[\]^`{|}</code>
and the Unicode replacement character U+FFFD.
</p>

<p>
Consider a compiled a package containing the package clause
<code>package math</code>, which exports function <code>Sin</code>, and
installed the compiled package in the file identified by
<code>"lib/math"</code>.
This table illustrates how <code>Sin</code> is accessed in files
that import the package after the
various types of import declaration.
</p>

<pre class="grammar">
Import declaration          Local name of Sin

import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin
</pre>

<p>
An import declaration declares a dependency relation between
the importing and imported package.
It is illegal for a package to import itself, directly or indirectly,
or to directly import a package without
referring to any of its exported identifiers. To import a package solely for
its side-effects (initialization), use the <a href="#Blank_identifier">blank</a>
identifier as explicit package name:
</p>

<pre>
import _ "lib/math"
</pre>


<h3 id="An_example_package">An example package</h3>

<p>
Here is a complete Go package that implements a concurrent prime sieve.
</p>

<pre>
package main

import "fmt"

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan&lt;- int) {
	for i := 2; ; i++ {
		ch &lt;- i  // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src &lt;-chan int, dst chan&lt;- int, prime int) {
	for i := range src {  // Loop over values received from 'src'.
		if i%prime != 0 {
			dst &lt;- i  // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int)  // Create a new channel.
	go generate(ch)       // Start generate() as a subprocess.
	for {
		prime := &lt;-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
</pre>

# Program initialization and execution

## The zero value

* == default variable value /  
  * | allocate storage, 
    * ❌NO initialized explicitly❌
  * depend -- on -- its type 
    * booleans -- `false`
    * numeric types -- `0`
    * string -- `""`
    * pointers, functions, interfaces, slices, channels, and maps  -- `nil` 
  * ⚠️apply | recursive types⚠️
  * can ALSO be specified EXPLICITLY

* use cases
  * declare a variable
  * call of `new`
  * create a NEW value -- through --
    * composite literal
    * call of `make`
  
## Package initialization

* package-level variable initialization
  * -- via -- steps /
    * select 1 variable / EACH step /
      * earliest | declaration order
        * == | package, from top -- to -- down 
      * requirements
        * NOT YET initialized 
        * either
          * NO initialization expression, OR 
          * ⚠️NO depend -- on -- uninitialized variables⚠️

* initialization cycle
  * == variable cyclic dependency
  * == variables / IMPOSSIBLE to initialize 

* | variable declaration initialization,
  * if | left-hand side, there are SEVERAL variables & | right-hand side, 1! value  -> ALL are initialized
    * == ALL or NONE are initialized

* | package initialization,
  * [blank identifier variables](#blank-identifier----_---) are treated -- like -- other variables | declarations

* if there are MULTIPLE files -> declaration order of variables is 👀determined -- by the -- order / files are presented | compiler👀
  * == if the file is declared FIRSTLY -> variables declared before
  * order / files are presented | compiler == order / passed | `go run`
  * _Example:_ `go run examples/package-initialization-multiple-z.go examples/package-initialization-multiple-a.go`

* dependency analysis
  * NOT rely on
    * CURRENT variables' values
  * rely on
    * 👀lexical references👀 -- to -- the variables | source / analyzed transitively
      * ⚠️lexical reference != copy by reference ⚠️
        * 👀if you need copy by referency -> use pointers 👀
  * use cases
    * reference to a variable OR function == identifier -- to that -- variable or function
    * TODO: A reference to a method <code>m</code> is a
<a href="#Method_values">method value</a> or
<a href="#Method_expressions">method expression</a> of the form
<code>t.m</code>, where the (static) type of <code>t</code> is
not an interface type, and the method <code>m</code> is in the
<a href="#Method_sets">method set</a> of <code>t</code>.
It is immaterial whether the resulting function value
<code>t.m</code> is invoked.
    *  variable, function, or method <code>x</code> depends on a variable
<code>y</code> if <code>x</code>'s initialization expression or body
(for functions and methods) contains a reference to <code>y</code>
or to a function or method that depends on <code>y</code>.
  * / EACH package

<p>
Dependency analysis is performed per package; only references referring
to variables, functions, and (non-interface) methods declared in the current
package are considered. If other, hidden, data dependencies exists between
variables, the initialization order between those variables is unspecified.
</p>

<p>
For instance, given the declarations
</p>

<pre>
var x = I(T{}).ab()   // x has an undetected, hidden dependency on a and b
var _ = sideEffect()  // unrelated to x, a, or b
var a = b
var b = 42

type I interface      { ab() []int }
type T struct{}
func (T) ab() []int   { return []int{a, b} }
</pre>

<p>
the variable <code>a</code> will be initialized after <code>b</code> but
whether <code>x</code> is initialized before <code>b</code>, between
<code>b</code> and <code>a</code>, or after <code>a</code>, and
thus also the moment at which <code>sideEffect()</code> is called (before
or after <code>x</code> is initialized) is not specified.
</p>

* `func init() { … }`
  * requirements
    * declare it | package block
    * NO arguments 
    * NO result parameters
    * functionName == `init`
  * uses
    * package-level variable initialization 
  * ALLOWED
    * define MULTIPLE | SAME
      * package
      * source file
  * | package block,
    * `init` identifier 
      * ONLY ALLOWED -- for -- `func init()`
      * ❌can NOT be referred ❌
  * package initialization steps
    * assign initial values / declared -- via -- ALL package-level variables
    * 👀call AUTOMATICALLY ALL `init` functions / follow the declaration order👀

## Program initialization

<p>
The packages of a complete program are initialized stepwise, one package at a time.
If a package has imports, the imported packages are initialized
before initializing the package itself. If multiple packages import
a package, the imported package will be initialized only once.
The importing of packages, by construction, guarantees that there
can be no cyclic initialization dependencies.
More precisely:
</p>

<p>
Given the list of all packages, sorted by import path, in each step the first
uninitialized package in the list for which all imported packages (if any) are
already initialized is <a href="#Package_initialization">initialized</a>.
This step is repeated until all packages are initialized.
</p>

<p>
Package initialization&mdash;variable initialization and the invocation of
<code>init</code> functions&mdash;happens in a single goroutine,
sequentially, one package at a time.
An <code>init</code> function may launch other goroutines, which can run
concurrently with the initialization code. However, initialization
always sequences
the <code>init</code> functions: it will not invoke the next one
until the previous one has returned.
</p>

<h3 id="Program_execution">Program execution</h3>
<p>
A complete program is created by linking a single, unimported package
called the <i>main package</i> with all the packages it imports, transitively.
The main package must
have package name <code>main</code> and
declare a function <code>main</code> that takes no
arguments and returns no value.
</p>

<pre>
func main() { … }
</pre>

<p>
Program execution begins by <a href="#Program_initialization">initializing the program</a>
and then invoking the function <code>main</code> in package <code>main</code>.
When that function invocation returns, the program exits.
It does not wait for other (non-<code>main</code>) goroutines to complete.
</p>

<h2 id="Errors">Errors</h2>

<p>
The predeclared type <code>error</code> is defined as
</p>

<pre>
type error interface {
	Error() string
}
</pre>

<p>
It is the conventional interface for representing an error condition,
with the nil value representing no error.
For instance, a function to read data from a file might be defined:
</p>

<pre>
func Read(f *File, b []byte) (n int, err error)
</pre>

# Run-time panics

* use cases / trigger them
  * execution errors
    * _Example:_ array[indexOutOfBounds]
  * call DIRECTLY [`panic(runtime.Error's value)`](#handling-panics)

# System considerations

## Package `unsafe`

* built-in package / 
  * known -- to the -- compiler
  * accessible -- through the -- import path `"unsafe"`
  * provides
    * 👀facilities -- for -- low-level programming👀
      * _Example:_ operations / violate the type system
  * ⚠️cons⚠️
    * | type safety, MUST be reviewed MANUALLY
    * may NOT be portable

<!--
These conversions also apply to type parameters with suitable core types.
Determine if we can simply use core type instead of underlying type here,
of if the general conversion rules take care of this.
-->

* `Pointer`
  * == pointer type
  * 's value
    * ⚠️may NOT be dereferenced ⚠️
    Any pointer or value of <a href="#Core_types">core type</a> <code>uintptr</code> can be
    <a href="#Conversions">converted</a> to a type of core type <code>Pointer</code> and vice versa.
    The effect of converting between <code>Pointer</code> and <code>uintptr</code> is implementation-defined.

<pre>
var f float64
bits = *(*uint64)(unsafe.Pointer(&amp;f))

type ptr unsafe.Pointer
bits = *(*uint64)(ptr(&amp;f))

func f[P ~*B, B any](p P) uintptr {
	return uintptr(unsafe.Pointer(p))
}

var p ptr = nil
</pre>

<p>
The functions <code>Alignof</code> and <code>Sizeof</code> take an expression <code>x</code>
of any type and return the alignment or size, respectively, of a hypothetical variable <code>v</code>
as if <code>v</code> was declared via <code>var v = x</code>.
</p>
<p>
The function <code>Offsetof</code> takes a (possibly parenthesized) <a href="#Selectors">selector</a>
<code>s.f</code>, denoting a field <code>f</code> of the struct denoted by <code>s</code>
or <code>*s</code>, and returns the field offset in bytes relative to the struct's address.
If <code>f</code> is an <a href="#Struct_types">embedded field</a>, it must be reachable
without pointer indirections through fields of the struct.
For a struct <code>s</code> with field <code>f</code>:
</p>

<pre>
uintptr(unsafe.Pointer(&amp;s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&amp;s.f))
</pre>

<p>
Computer architectures may require memory addresses to be <i>aligned</i>;
that is, for addresses of a variable to be a multiple of a factor,
the variable's type's <i>alignment</i>.  The function <code>Alignof</code>
takes an expression denoting a variable of any type and returns the
alignment of the (type of the) variable in bytes.  For a variable
<code>x</code>:
</p>

<pre>
uintptr(unsafe.Pointer(&amp;x)) % unsafe.Alignof(x) == 0
</pre>

<p>
A (variable of) type <code>T</code> has <i>variable size</i> if <code>T</code>
is a <a href="#Type_parameter_declarations">type parameter</a>, or if it is an
array or struct type containing elements
or fields of variable size. Otherwise the size is <i>constant</i>.
Calls to <code>Alignof</code>, <code>Offsetof</code>, and <code>Sizeof</code>
are compile-time <a href="#Constant_expressions">constant expressions</a> of
type <code>uintptr</code> if their arguments (or the struct <code>s</code> in
the selector expression <code>s.f</code> for <code>Offsetof</code>) are types
of constant size.
</p>

<p>
The function <code>Add</code> adds <code>len</code> to <code>ptr</code>
and returns the updated pointer <code>unsafe.Pointer(uintptr(ptr) + uintptr(len))</code>
[<a href="#Go_1.17">Go 1.17</a>].
The <code>len</code> argument must be of <a href="#Numeric_types">integer type</a> or an untyped <a href="#Constants">constant</a>.
A constant <code>len</code> argument must be <a href="#Representability">representable</a> by a value of type <code>int</code>;
if it is an untyped constant it is given type <code>int</code>.
The rules for <a href="/pkg/unsafe#Pointer">valid uses</a> of <code>Pointer</code> still apply.
</p>

<p>
The function <code>Slice</code> returns a slice whose underlying array starts at <code>ptr</code>
and whose length and capacity are <code>len</code>.
<code>Slice(ptr, len)</code> is equivalent to
</p>

<pre>
(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
</pre>

<p>
except that, as a special case, if <code>ptr</code>
is <code>nil</code> and <code>len</code> is zero,
<code>Slice</code> returns <code>nil</code>
[<a href="#Go_1.17">Go 1.17</a>].
</p>

<p>
The <code>len</code> argument must be of <a href="#Numeric_types">integer type</a> or an untyped <a href="#Constants">constant</a>.
A constant <code>len</code> argument must be non-negative and <a href="#Representability">representable</a> by a value of type <code>int</code>;
if it is an untyped constant it is given type <code>int</code>.
At run time, if <code>len</code> is negative,
or if <code>ptr</code> is <code>nil</code> and <code>len</code> is not zero,
a <a href="#Run_time_panics">run-time panic</a> occurs
[<a href="#Go_1.17">Go 1.17</a>].
</p>

<p>
The function <code>SliceData</code> returns a pointer to the underlying array of the <code>slice</code> argument.
If the slice's capacity <code>cap(slice)</code> is not zero, that pointer is <code>&slice[:1][0]</code>.
If <code>slice</code> is <code>nil</code>, the result is <code>nil</code>.
Otherwise it  is a non-<code>nil</code> pointer to an unspecified memory address
[<a href="#Go_1.20">Go 1.20</a>].
</p>

<p>
The function <code>String</code> returns a <code>string</code> value whose underlying bytes start at
<code>ptr</code> and whose length is <code>len</code>.
The same requirements apply to the <code>ptr</code> and <code>len</code> argument as in the function
<code>Slice</code>. If <code>len</code> is zero, the result is the empty string <code>""</code>.
Since Go strings are immutable, the bytes passed to <code>String</code> must not be modified afterwards.
[<a href="#Go_1.20">Go 1.20</a>]
</p>

<p>
The function <code>StringData</code> returns a pointer to the underlying bytes of the <code>str</code> argument.
For an empty string the return value is unspecified, and may be <code>nil</code>.
Since Go strings are immutable, the bytes returned by <code>StringData</code> must not be modified
[<a href="#Go_1.20">Go 1.20</a>].
</p>

<h3 id="Size_and_alignment_guarantees">Size and alignment guarantees</h3>

<p>
For the <a href="#Numeric_types">numeric types</a>, the following sizes are guaranteed:
</p>

<pre class="grammar">
type                                 size in bytes

byte, uint8, int8                     1
uint16, int16                         2
uint32, int32, float32                4
uint64, int64, float64, complex64     8
complex128                           16
</pre>

<p>
The following minimal alignment properties are guaranteed:
</p>
<ol>
<li>For a variable <code>x</code> of any type: <code>unsafe.Alignof(x)</code> is at least 1.
</li>

<li>For a variable <code>x</code> of struct type: <code>unsafe.Alignof(x)</code> is the largest of
   all the values <code>unsafe.Alignof(x.f)</code> for each field <code>f</code> of <code>x</code>, but at least 1.
</li>

<li>For a variable <code>x</code> of array type: <code>unsafe.Alignof(x)</code> is the same as
	the alignment of a variable of the array's element type.
</li>
</ol>

<p>
A struct or array type has size zero if it contains no fields (or elements, respectively) that have a size greater than zero. Two distinct zero-size variables may have the same address in memory.
</p>

# Appendix

## Language versions

<p>
The <a href="/doc/go1compat">Go 1 compatibility guarantee</a> ensures that
programs written to the Go 1 specification will continue to compile and run
correctly, unchanged, over the lifetime of that specification.
More generally, as adjustments are made and features added to the language,
the compatibility guarantee ensures that a Go program that works with a
specific Go language version will continue to work with any subsequent version.
</p>

<p>
For instance, the ability to use the prefix <code>0b</code> for binary
integer literals was introduced with Go 1.13, indicated
by [<a href="#Go_1.13">Go 1.13</a>] in the section on
<a href="#Integer_literals">integer literals</a>.
Source code containing an integer literal such as <code>0b1011</code>
will be rejected if the implied or required language version used by
the compiler is older than Go 1.13.
</p>

<p>
The following table describes the minimum language version required for
features introduced after Go 1.
</p>

### Go 1.9

* [alias declaration](#alias-declarations)
  * uses
    * declare an alias name / type

<h4 id="Go_1.13">Go 1.13</h4>
<ul>
<li>
<a href="#Integer_literals">Integer literals</a> may use the prefixes <code>0b</code>, <code>0B</code>, <code>0o</code>,
and <code>0O</code> for binary, and octal literals, respectively.
</li>
<li>
Hexadecimal <a href="#Floating-point_literals">floating-point literals</a> may be written using the prefixes
<code>0x</code> and <code>0X</code>.
</li>
<li>
The <a href="#Imaginary_literals">imaginary suffix</a> <code>i</code> may be used with any (binary, decimal, hexadecimal)
integer or floating-point literal, not just decimal literals.
</li>
<li>
The digits of any number literal may be <a href="#Integer_literals">separated</a> (grouped)
using underscores <code>_</code>.
</li>
<li>
The shift count in a <a href="#Operators">shift operation</a> may be a signed integer type.
</li>
</ul>

<h4 id="Go_1.14">Go 1.14</h4>
<ul>
<li>
Emdedding a method more than once through different <a href="#Embedded_interfaces">embedded interfaces</a>
is not an error.
</li>
</ul>

<h4 id="Go_1.17">Go 1.17</h4>
<ul>
<li>
A slice may be <a href="#Conversions">converted</a> to an array pointer if the slice and array element
types match, and the array is not longer than the slice.
</li>
<li>
The built-in <a href="#Package_unsafe">package <code>unsafe</code></a> includes the new functions
<code>Add</code> and <code>Slice</code>.
</li>
</ul>

### Go 1.18
<p>
The 1.18 release adds polymorphic functions and types ("generics") to the language.
Specifically:
</p>
<ul>
<li>
The set of <a href="#Operators_and_punctuation">operators and punctuation</a> includes the new token <code>~</code>.
</li>
<li>
Function and type declarations may declare <a href="#Type_parameter_declarations">type parameters</a>.
</li>
<li>
Interface types may <a href="#General_interfaces">embed arbitrary types</a> (not just type names of interfaces)
as well as union and <code>~T</code> type elements.
</li>
<li>
The set of <a href="#Predeclared_identifiers">predeclared</a> types includes the new types
<code>any</code> and <code>comparable</code>.
</li>
</ul>

<h4 id="Go_1.20">Go 1.20</h4>
<ul>
<li>
A slice may be <a href="#Conversions">converted</a> to an array if the slice and array element
types match and the array is not longer than the slice.
</li>
<li>
The built-in <a href="#Package_unsafe">package <code>unsafe</code></a> includes the new functions
<code>SliceData</code>, <code>String</code>, and <code>StringData</code>.
</li>
<li>
<a href="#Comparison_operators">Comparable types</a> (such as ordinary interfaces) may satisfy
<code>comparable</code> constraints, even if the type arguments are not strictly comparable.
</li>
</ul>

<h4 id="Go_1.21">Go 1.21</h4>
<ul>
<li>
The set of <a href="#Predeclared_identifiers">predeclared</a> functions includes the new functions
<code>min</code>, <code>max</code>, and <code>clear</code>.
</li>
<li>
<a href="#Type_inference">Type inference</a> uses the types of interface methods for inference.
It also infers type arguments for generic functions assigned to variables or
passed as arguments to other (possibly generic) functions.
</li>
</ul>

<h4 id="Go_1.22">Go 1.22</h4>
<ul>
<li>
In a <a href="#For_statements">"for" statement</a>, each iteration has its own set of iteration
variables rather than sharing the same variables in each iteration.
</li>
<li>
A "for" statement with <a href="#For_range">"range" clause</a> may iterate over
integer values from zero to an upper limit.
</li>
</ul>

<h3 id="Type_unification_rules">Type unification rules</h3>

<p>
The type unification rules describe if and how two types unify.
The precise details are relevant for Go implementations,
affect the specifics of error messages (such as whether
a compiler reports a type inference or other error),
and may explain why type inference fails in unusual code situations.
But by and large these rules can be ignored when writing Go code:
type inference is designed to mostly "work as expected",
and the unification rules are fine-tuned accordingly.
</p>

<p>
Type unification is controlled by a <i>matching mode</i>, which may
be <i>exact</i> or <i>loose</i>.
As unification recursively descends a composite type structure,
the matching mode used for elements of the type, the <i>element matching mode</i>,
remains the same as the matching mode except when two types are unified for
<a href="#Assignability">assignability</a> (<code>≡<sub>A</sub></code>):
in this case, the matching mode is <i>loose</i> at the top level but
then changes to <i>exact</i> for element types, reflecting the fact
that types don't have to be identical to be assignable.
</p>

<p>
Two types that are not bound type parameters unify exactly if any of
following conditions is true:
</p>

<ul>
<li>
	Both types are <a href="#Type_identity">identical</a>.
</li>
<li>
	Both types have identical structure and their element types
	unify exactly.
</li>
<li>
	Exactly one type is an <a href="#Type_inference">unbound</a>
	type parameter with a <a href="#Core_types">core type</a>,
	and that core type unifies with the other type per the
	unification rules for <code>≡<sub>A</sub></code>
	(loose unification at the top level and exact unification
	for element types).
</li>
</ul>

<p>
If both types are bound type parameters, they unify per the given
matching modes if:
</p>

<ul>
<li>
	Both type parameters are identical.
</li>
<li>
	At most one of the type parameters has a known type argument.
	In this case, the type parameters are <i>joined</i>:
	they both stand for the same type argument.
	If neither type parameter has a known type argument yet,
	a future type argument inferred for one the type parameters
	is simultaneously inferred for both of them.
</li>
<li>
	Both type parameters have a known type argument
	and the type arguments unify per the given matching modes.
</li>
</ul>

<p>
A single bound type parameter <code>P</code> and another type <code>T</code> unify
per the given matching modes if:
</p>

<ul>
<li>
	<code>P</code> doesn't have a known type argument.
	In this case, <code>T</code> is inferred as the type argument for <code>P</code>.
</li>
<li>
	<code>P</code> does have a known type argument <code>A</code>,
	<code>A</code> and <code>T</code> unify per the given matching modes,
	and one of the following conditions is true:
	<ul>
	<li>
		Both <code>A</code> and <code>T</code> are interface types:
		In this case, if both <code>A</code> and <code>T</code> are
		also <a href="#Type_definitions">defined</a> types,
		they must be <a href="#Type_identity">identical</a>.
		Otherwise, if neither of them is a defined type, they must
		have the same number of methods
		(unification of <code>A</code> and <code>T</code> already
		established that the methods match).
	</li>
	<li>
		Neither <code>A</code> nor <code>T</code> are interface types:
		In this case, if <code>T</code> is a defined type, <code>T</code>
		replaces <code>A</code> as the inferred type argument for <code>P</code>.
	</li>
	</ul>
</li>
</ul>

<p>
Finally, two types that are not bound type parameters unify loosely
(and per the element matching mode) if:
</p>

<ul>
<li>
	Both types unify exactly.
</li>
<li>
	One type is a <a href="#Type_definitions">defined type</a>,
	the other type is a type literal, but not an interface,
	and their underlying types unify per the element matching mode.
</li>
<li>
	Both types are interfaces (but not type parameters) with
	identical <a href="#Interface_types">type terms</a>,
	both or neither embed the predeclared type
	<a href="#Predeclared_identifiers">comparable</a>,
	corresponding method types unify exactly,
	and the method set of one of the interfaces is a subset of
	the method set of the other interface.
</li>
<li>
	Only one type is an interface (but not a type parameter),
	corresponding methods of the two types unify per the element matching mode,
	and the method set of the interface is a subset of
	the method set of the other type.
</li>
<li>
	Both types have the same structure and their element types
	unify per the element matching mode.
</li>
</ul>
