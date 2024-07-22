package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func mustParse(fset *token.FileSet, filename string, src []byte) *ast.File {
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}

func parse(root, path string) Package {
	ent, _ := os.ReadDir(filepath.Join(root, path))

	fset := token.NewFileSet()

	files := []*ast.File{}

	pkgs := []Package{}
	fnames := []string{}

	for _, e := range ent {
		next := filepath.Join(path, e.Name())

		if e.IsDir() {
			pkgs = append(pkgs, parse(root, next))
		} else {
			if !strings.HasSuffix(e.Name(), ".go") {
				continue
			}
			fnames = append(fnames, e.Name())

			src, _ := os.ReadFile(filepath.Join(root, next))
			files = append(files, mustParse(fset, e.Name(), src))
		}
	}

	p, err := doc.NewFromFiles(fset, files, "example.com")
	if err != nil {
		panic(err)
	}

	return NewPackage(fset, p, path, pkgs, fnames)
}
