package main

import (
	"log"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

const (
	version = "0.1.3"
	usage   = `usage: gdmd [options] <directory>

go doc markdown generator

options:`
)

func main() {
	hFlag := flag.BoolP("help", "h", false, "print this help message")
	vFlag := flag.BoolP("version", "v", false, "print version")
	rFlag := flag.BoolP("recursive", "r", false, "walk directories recursively")

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
