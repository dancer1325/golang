package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 1. Join()
	join()
}

func join() {
	// 1. simple basic file path
	path := filepath.Join("dir", "subdir", "file.txt")
	fmt.Println("basic file path: ", path) // separated -- with -- OS specific separator
	// | Windows: dir\subdir\file.txt
	// | Unix: dir/subdir/file.txt

	// 2. redundant separators are removed
	path = filepath.Join("dir/", "/subdir/", "/file.txt")
	fmt.Println("redundant separators: ", path)

	// 3. empty elements are ignored
	path = filepath.Join("dir", "", "subdir", "", "file.txt")
	fmt.Println("empty elements are removed: ", path)

	// 4. manage relative paths
	path = filepath.Join("dir", "..", "otherdir", "file.txt")
	fmt.Println("manage relative paths: ", path)

	// 5. base route + subroute
	basePath := filepath.Join("home", "user")
	configPath := filepath.Join(basePath, "config", "settings.json")
	fmt.Println("complete path: ", configPath)
}
