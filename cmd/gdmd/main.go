package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

const version = "0.0.0"

func main() {
	vFlag := flag.Bool("v", false, "print version")

	flag.Parse()

	if *vFlag {
		println(version)
		return
	}

	args := os.Args
	root, _ := filepath.Abs(args[1])

	pkg, err := Parse(root, "")
	if err != nil {
		log.Fatal(err)
	}
	Generate(root, &pkg)
}
