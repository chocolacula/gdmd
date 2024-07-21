package main

import (
	"bytes"
	"go/doc"
	"go/printer"
	"go/token"
)

type Package struct {
	Doc       string
	Constants []Constant
	Variables []Variable
	Functions []Function
	Types     []Type
	Deps      []Package
	Files     []string
}

// Constant represents constant declarations within () or single constant.
type Constant struct {
	Doc   string // doc comment under the block or single constant
	Names []string
	Src   []byte // piece of source code with the declaration
}

func NewConstant(fset *token.FileSet, v *doc.Value) Constant {
	b := bytes.Buffer{}
	printer.Fprint(&b, fset, v.Decl)

	return Constant{
		Names: v.Names,
		Doc:   v.Doc,
		Src:   b.Bytes(),
	}
}

type Variable struct {
	Name  string
	Type  string
	Value string
}

type Function struct {
	Doc       string
	Name      string
	Signature []byte
}

func NewFunction(fset *token.FileSet, v *doc.Func) Function {
	b := bytes.Buffer{}
	printer.Fprint(&b, fset, v.Decl)

	return Function{
		Doc:       v.Doc,
		Name:      v.Name,
		Signature: b.Bytes(),
	}
}

type Type struct {
	Doc       string
	Name      string
	Constants []Constant
	Variables []Variable
	Functions []Function
	Methods   []Function
}
