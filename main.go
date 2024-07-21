package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
)

func mustParse(fset *token.FileSet, filename string, src []byte) *ast.File {
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}

func walk(root string) {
	ent, _ := os.ReadDir(root)

	fset := token.NewFileSet()
	files := []*ast.File{}

	for _, e := range ent {
		if e.IsDir() {
			// walk(e.Name())
		} else {
			src, _ := os.ReadFile(filepath.Join(root, e.Name()))
			files = append(files, mustParse(fset, e.Name(), src))
		}
	}

	// Compute package documentation with examples.
	p, err := doc.NewFromFiles(fset, files, "example.com/p")
	if err != nil {
		panic(err)
	}

	for _, v := range p.Consts {
		c := NewConstant(fset, v)
		println(string(c.Src))
	}
	f := NewFunction(fset, p.Funcs[0])
	println(string(f.Signature))

	for _, v := range p.Types {
		printer.Fprint(os.Stdout, fset, v.Decl)
	}
}

func main() {
	// dir, _ := os.Getwd()

	args := os.Args

	// data, err := os.ReadFile(path)
	// defer file.Close()

	// var tree map[string]any

	// save := func(path string, d os.DirEntry, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}

	// 	println(path)

	// 	if d.IsDir() {
	// 		tree[path] = map[string]any{}
	// 	} else if strings.HasSuffix(path, ".go") {
	// 		name := filepath.Base(path)
	// 		dir := filepath.Dir(path)

	// 		if m, ok := tree[dir].(map[string]any); ok {
	// 			m[name]
	// 		}
	// 		tree[path] = path
	// 	}

	// 	return nil
	// }

	// err := filepath.WalkDir(p, save)
	// if err != nil {
	// 	fmt.Printf("Error walking directory: %v\n", err)
	// }

	p, _ := filepath.Abs(args[1])
	walk(p)
}
