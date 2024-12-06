package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

const (
	version = "0.1.4"
	usage   = `usage: gdmd [options] <directory>

go doc Markdown generator

options:`
)

func main() {
	hFlag := flag.Bool("help", false, "print this help message")
	vFlag := flag.Bool("v", false, "print version")
	rFlag := flag.Bool("r", false, "walk directories recursively")

	flag.Parse()

	if *hFlag {
		println(usage)
		flag.PrintDefaults()
		return
	}
	if *vFlag {
		println(version)
		return
	}

	root, _ := filepath.Abs(flag.Arg(0))

	_, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("directory %s does not exist", root)
		}
		log.Fatal(err)
	}

	pkg, err := Parse(root, "", *rFlag)
	if err != nil {
		log.Fatal(err)
	}
	err = Generate(root, &pkg)
	if err != nil {
		log.Fatal(err)
	}
}
