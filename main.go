package main

import (
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	root, _ := filepath.Abs(args[1])

	pkg := parse(root, "")

	generate(root, &pkg)
}
