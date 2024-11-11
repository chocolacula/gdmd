# Overview

package `main`

The single package in the project, contains data representation, parsing and generation logic.

## Index

- [Variables](#variables)
- [Functions](#functions)
  - [func Generate(root string, pkg *Package)](#func-generate)
- [Types](#types)
  - [type Function](#type-function)
    - [func NewFunction(fset *token.FileSet, f *doc.Func) Function](#func-newfunction)
  - [type Package](#type-package)
    - [func NewPackage(fset *token.FileSet, p *doc.Package, dir string, nested []Package, files []string) Package](#func-newpackage)
    - [func Parse(root, path string, recursive bool) (Package, error)](#func-parse)
  - [type Position](#type-position)
  - [type Type](#type-type)
    - [func NewType(fset *token.FileSet, t *doc.Type) Type](#func-newtype)
  - [type Variable](#type-variable)
    - [func NewVariable(fset *token.FileSet, v *doc.Value) Variable](#func-newvariable)
- [Source files](#source-files)

## Variables

Simple error to indicate empty folder

```go
var EmptyErr = errors.New("empty folder")
```

## Functions

### func [Generate](./generate.go#L30)

```go
func Generate(root string, pkg *Package)
```

Generate creates a markdown files for the given [Package] and its nested packages.

## Types

### type [Function](./types.go#L85)

```go
type Function struct {
  Doc       string
  Name      string
  Pos       Position
  Recv      string // "" for functions, receiver name for methods
  Signature string
}
```

Function represents a function or method declaration.

### func [NewFunction](./types.go#L93)

```go
func NewFunction(fset *token.FileSet, f *doc.Func) Function
```

### type [Package](./types.go#L14)

```go
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
```

Package represents a Go package with its contents.

### func [NewPackage](./types.go#L26)

```go
func NewPackage(fset *token.FileSet, p *doc.Package, dir string, nested []Package, files []string) Package
```

### func [Parse](./parse.go#L28)

```go
func Parse(root, path string, recursive bool) (Package, error)
```

Parse walks the directory tree rooted at root and parses all .go files
it returns a [Package] for each directory containing .go files
or empty [Package] and [EmptyErr]

### type [Position](./types.go#L79)

```go
type Position struct {
  Filename string
  Line     int
}
```

Position is a file name and line number of a declaration.

### type [Type](./types.go#L114)

```go
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
```

Type is a struct or interface declaration.

### func [NewType](./types.go#L125)

```go
func NewType(fset *token.FileSet, t *doc.Type) Type
```

### type [Variable](./types.go#L61)

```go
type Variable struct {
  Doc   string // doc comment under the block or single declaration
  Names []string
  Src   string // piece of source code with the declaration
}
```

Variable represents constant or variable declarations within () or single one.

### func [NewVariable](./types.go#L67)

```go
func NewVariable(fset *token.FileSet, v *doc.Value) Variable
```

## Source files

[generate.go](./generate.go)
[main.go](./main.go)
[parse.go](./parse.go)
[template.go](./template.go)
[types.go](./types.go)
