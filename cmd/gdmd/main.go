package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	version = "0.1.2"
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

	fset := flag.NewFlagSet("parse", flag.ExitOnError)
	rFlag := fset.Bool("r", false, "recursive")

	if !strings.HasPrefix(flag.Arg(0), "-") {
		// cut of the directory argument
		_ = fset.Parse(flag.Args()[1:])
	}

	root, _ := filepath.Abs(flag.Arg(0))

	_, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("directory %s does not exist", root)
		} else {
			log.Fatal(err)
		}
	}

	pkg, err := Parse(root, "", *rFlag)
	if err != nil {
		log.Fatal(err)
	}
	Generate(root, &pkg)
}
