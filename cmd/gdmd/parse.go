package main

import (
	"errors"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Simple error to indicate empty folder
var EmptyErr = errors.New("empty folder")

func mustParse(fset *token.FileSet, filename string, src []byte) *ast.File {
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}

// Parse walks the directory tree rooted at root and parses all .go files
// it returns a [Package] for each directory containing .go files
// or empty [Package] and [EmptyErr]
func Parse(root, path string, recursive bool) (Package, error) {
	ent, _ := os.ReadDir(filepath.Join(root, path))

	fset := token.NewFileSet()

	files := []*ast.File{}

	pkgs := []Package{}
	fnames := []string{}

	for _, e := range ent {
		next := filepath.Join(path, e.Name())

		if e.IsDir() && recursive {
			pkg, err := Parse(root, next, recursive)
			if err == nil {
				pkgs = append(pkgs, pkg)
			} // else ignore error
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
		return Package{}, err
	}
	if len(fnames) == 0 && len(pkgs) == 0 {
		return Package{}, EmptyErr
	}
	return NewPackage(fset, p, path, pkgs, fnames), nil
}
