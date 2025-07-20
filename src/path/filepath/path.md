* TODO:

* `func Join(elem ...string) string {}`
  * 👀ANY NUMBER of path elements are joined -- into -- 1! path 👀/ 
    * ⚠️separated -- with an -- OS specific `Separator`⚠️
    * ignore empty elements 
    * result is Cleaned
      * if there are redundant separators -> removed
      * manage relative paths
    * if `elem ...string` == empty OR ALL elements are empty -> returns ""
    * | Windows,
      * if the FIRST non-empty element == UNC path -> result == UNC path 

* TODO: