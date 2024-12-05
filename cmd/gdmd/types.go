// The single package in the project, contains data representation, parsing, and generation logic.
package main

import (
	"go/doc"
	"go/doc/comment"
	"go/printer"
	"go/token"
	"strings"
)

var printerConf = printer.Config{Mode: printer.UseSpaces, Tabwidth: 2}

// Package represents a Go package with its contents.
type Package struct {
	Doc       string
	Name      string
	Dir       string
	Constants []Variable
	Variables []Variable
	Functions []Function
	Types     []Type
	Nested    []Package
	Files     []string
}

func NewPackage(fset *token.FileSet, p *doc.Package, dir string, nested []Package, files []string) (Package, error) {
	cp := p.Parser()

	consts := []Variable{}
	for _, c := range p.Consts {
		nc, err := NewVariable(cp, fset, c)
		if err != nil {
			return Package{}, err
		}

		consts = append(consts, nc)
	}

	vars := []Variable{}
	for _, v := range p.Vars {
		nv, err := NewVariable(cp, fset, v)
		if err != nil {
			return Package{}, err
		}

		vars = append(vars, nv)
	}

	funcs := []Function{}
	for _, f := range p.Funcs {
		nf, err := NewFunction(cp, fset, f)
		if err != nil {
			return Package{}, err
		}

		funcs = append(funcs, nf)
	}

	types := []Type{}
	for _, t := range p.Types {
		nt, err := NewType(cp, fset, t)
		if err != nil {
			return Package{}, err
		}
		types = append(types, nt)
	}

	return Package{
		Doc:       p.Doc,
		Name:      p.Name,
		Dir:       dir,
		Constants: consts,
		Variables: vars,
		Functions: funcs,
		Types:     types,
		Nested:    nested,
		Files:     files,
	}, nil
}

// Variable represents constant or variable declarations within () or single one.
type Variable struct {
	Doc   string // doc comment under the block or single declaration
	Names []string
	Src   string // piece of source code with the declaration
}

func NewVariable(cp *comment.Parser, fset *token.FileSet, v *doc.Value) (Variable, error) {
	b := strings.Builder{}
	err := printerConf.Fprint(&b, fset, v.Decl)
	if err != nil {
		return Variable{}, err
	}

	return Variable{
		Names: v.Names,
		Doc:   computeLinks(cp, v.Doc),
		Src:   b.String(),
	}, nil
}

// Position is a file name and line number of a declaration.
type Position struct {
	Filename string
	Line     int
}

// Function represents a function or method declaration.
type Function struct {
	Doc       string
	Name      string
	Pos       Position
	Recv      string // "" for functions, receiver name for methods
	Signature string
}

func NewFunction(cp *comment.Parser, fset *token.FileSet, f *doc.Func) (Function, error) {
	b := strings.Builder{}
	err := printerConf.Fprint(&b, fset, f.Decl)
	if err != nil {
		return Function{}, err
	}
	pos := fset.Position(f.Decl.Pos())

	recv := ""
	if f.Decl.Recv != nil {
		recv = f.Decl.Recv.List[0].Names[0].Name
	}

	return Function{
		Doc:       computeLinks(cp, f.Doc),
		Name:      f.Name,
		Pos:       Position{pos.Filename, pos.Line},
		Recv:      recv,
		Signature: b.String(),
	}, nil
}

// Type is a struct or interface declaration.
type Type struct {
	Doc       string
	Name      string
	Pos       Position
	Src       string // piece of source code with the declaration
	Constants []Variable
	Variables []Variable
	Functions []Function
	Methods   []Function
}

func NewType(cp *comment.Parser, fset *token.FileSet, t *doc.Type) (Type, error) {
	b := strings.Builder{}
	err := printerConf.Fprint(&b, fset, t.Decl)
	if err != nil {
		return Type{}, err
	}
	consts := []Variable{}
	for _, c := range t.Consts {
		nc, err := NewVariable(cp, fset, c)
		if err != nil {
			return Type{}, err
		}
		consts = append(consts, nc)
	}

	vars := []Variable{}
	for _, v := range t.Vars {
		nv, err := NewVariable(cp, fset, v)
		if err != nil {
			return Type{}, err
		}
		vars = append(vars, nv)
	}

	funcs := []Function{}
	for _, f := range t.Funcs {
		nf, err := NewFunction(cp, fset, f)
		if err != nil {
			return Type{}, err
		}
		funcs = append(funcs, nf)
	}

	methods := []Function{}
	for _, m := range t.Methods {
		nm, err := NewFunction(cp, fset, m)
		if err != nil {
			return Type{}, err
		}
		methods = append(methods, nm)
	}

	pos := fset.Position(t.Decl.Pos())

	return Type{
		Doc:       t.Doc,
		Name:      t.Name,
		Pos:       Position{pos.Filename, pos.Line},
		Src:       b.String(),
		Constants: consts,
		Variables: vars,
		Functions: funcs,
		Methods:   methods,
	}, nil
}

// computeLinks adds markdown links to the documentation.
func computeLinks(p *comment.Parser, s string) string {
	docComment := p.Parse(s)
	cp := comment.Printer{
		DocLinkURL: func(link *comment.DocLink) string {
			if link.ImportPath == "" {
				// TODO: add link to current package, for now they will have no links
				return ""
			}

			first, _, _ := strings.Cut(link.ImportPath, "/")
			if strings.Contains(first, ".") {
				// this is supposed to catch github.com, gitlab.com but also all vanity URLs

				// here we are assuming that documentation is public
				// TODO: add support for private documentation by checking if import path is part of the current module

				return link.DefaultURL("https://pkg.go.dev/")
			}

			// TODO: sort what to do with relative import paths, for now they will have no links
			return ""
		},
	}
	return string(cp.Markdown(docComment))
}
