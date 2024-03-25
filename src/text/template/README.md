# Goal
* data-driven templates — generate → textual output
  * == html/template but generating there HTML

# How does it work?
* templates — are applied to a — data structure
* execution of the template
  * walks through the data structure
    * parallel executions could be safely done -- if parallel executions share Writer → output may be interleaved --
  * current location of the cursor, represented by “dot” `.`

# Template’s Input text
* UTF-8 encoded
* any format

# Actions
* := data evaluations OR control structures which
  * `{{ ContainedHere }}`
    * all text outside here → copied directly as output

# Text & spaces
* if template with text between actions is executed → text is copied exactly as it’s 
* for removing  white spaces
  * `“{{-`  — trims the — preceding tailing whitespaces
  * `-}}` — trims the — next leading whitespaces

# Note
* Check 'exampledoc_test.go' & 'golang-examples' repo to comprehend the library
## How to run it?
* Problems:
  * Problem1: Any `go` command runs
    * Reason: 'go.mod', specifying `go 1.23` which NOT downloaded - `go list -m` to check the location - 
    * Solution: Modify to a downloaded go version
* TODO:
  * Attempt1: `go run exampledoc_test.go` -- `go run` is JUST valid for compiling and running GO source files -
  * Attempt2: `go test exampledoc_test.go` 
