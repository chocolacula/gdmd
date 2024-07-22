package main

import (
	"go/doc"
	"go/printer"
	"go/token"
	"strings"
)

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

func NewPackage(fset *token.FileSet, p *doc.Package, dir string, nested []Package, files []string) Package {
	consts := []Variable{}
	for _, c := range p.Consts {
		consts = append(consts, NewVariable(fset, c))
	}

	vars := []Variable{}
	for _, v := range p.Vars {
		vars = append(vars, NewVariable(fset, v))
	}

	funcs := []Function{}
	for _, f := range p.Funcs {
		funcs = append(funcs, NewFunction(fset, f))
	}

	types := []Type{}
	for _, t := range p.Types {
		types = append(types, NewType(fset, t))
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
	}
}

// Variable represents constant or variable declarations within () or single one.
type Variable struct {
	Doc   string // doc comment under the block or single declaration
	Names []string
	Src   string // piece of source code with the declaration
}

func NewVariable(fset *token.FileSet, v *doc.Value) Variable {
	b := strings.Builder{}
	printer.Fprint(&b, fset, v.Decl)

	return Variable{
		Names: v.Names,
		Doc:   v.Doc,
		Src:   b.String(),
	}
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
	Signature string
}

func NewFunction(fset *token.FileSet, f *doc.Func) Function {
	b := strings.Builder{}
	printer.Fprint(&b, fset, f.Decl)

	pos := fset.Position(f.Decl.Pos())

	return Function{
		Doc:       f.Doc,
		Name:      f.Name,
		Pos:       Position{pos.Filename, pos.Line},
		Signature: b.String(),
	}
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

func NewType(fset *token.FileSet, t *doc.Type) Type {
	b := strings.Builder{}
	printer.Fprint(&b, fset, t.Decl)

	consts := []Variable{}
	for _, c := range t.Consts {
		consts = append(consts, NewVariable(fset, c))
	}

	vars := []Variable{}
	for _, v := range t.Vars {
		vars = append(vars, NewVariable(fset, v))
	}

	funcs := []Function{}
	for _, f := range t.Funcs {
		funcs = append(funcs, NewFunction(fset, f))
	}

	methods := []Function{}
	for _, m := range t.Methods {
		methods = append(methods, NewFunction(fset, m))
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
	}
}
