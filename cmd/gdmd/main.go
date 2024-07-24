package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

const (
	version = "0.0.0"
	usage   = `usage: gdmd <directory>

go doc markdown

options:`
)

func main() {
	flag.Usage = func() {
		println(usage)
		flag.PrintDefaults()
	}
	vFlag := flag.Bool("v", false, "print version")

	flag.Parse()

	if *vFlag {
		println(version)
		return
	}

	args := os.Args
	root, _ := filepath.Abs(args[1])

	_, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("directory %s does not exist", root)
		} else {
			log.Fatal(err)
		}
	}

	pkg, err := Parse(root, "")
	if err != nil {
		log.Fatal(err)
	}
	Generate(root, &pkg)
}
