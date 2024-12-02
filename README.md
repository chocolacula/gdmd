# Go Doc Markdown

godoc alternative for static Markdown documentation.

## Motivation

Standard [pkg.go.dev](https://pkg.go.dev/) is awesome but not suitable for private repos. You can use [pkgsite](https://cs.opensource.google/go/x/pkgsite) as local alternative but it's not perfect. Both occasionally skip certain files and generate incomplete documentation.

As stated in the [comment](https://github.com/golang/go/issues/2381#issuecomment-2183224009) of an engineer from Amazon sometimes it's a nightmare. You have to checkout a repo, run local server and open it in a browser only to read documentation.

With `gdmd` you don't even have to host your documentation, you can keep it in your repo along with the code. The generator creates a `README.md` with a package documentation in a package folder.
You can navigate through the documentation right in GitHub UI, any open directory with source code will render its documentation.

## How to use

First, you have to

```sh
go install github.com/chocolacula/gdmd/cmd/gdmd@latest
```

Then generate documentation for a package in a directory

```sh
gdmd ./directory
```

## Example

package `main`

The single package in the project, contains data representation, parsing, and generation logic.

## Index

- [Variables](#variables)
- [Functions](#functions)
  - [func Generate(root string, pkg *Package) error](#func-generate)
- [Types](#types)
  - [type Function](#type-function)
    - [func NewFunction(fset *token.FileSet, f *doc.Func) (Function, error)](#func-newfunction)
  - [type Package](#type-package)
    - [func NewPackage(fset *token.FileSet, p *doc.Package, dir string, nested []Package, files []string) (Package, error)](#func-newpackage)
    - [func Parse(root, path string, recursive bool) (Package, error)](#func-parse)
  - [type Position](#type-position)
  - [type Type](#type-type)
    - [func NewType(fset *token.FileSet, t *doc.Type) (Type, error)](#func-newtype)
  - [type Variable](#type-variable)
    - [func NewVariable(fset *token.FileSet, v *doc.Value) (Variable, error)](#func-newvariable)
- [Source files](#source-files)

## Variables

ErrEmpty sentinel indicating empty folder

```go
var ErrEmpty = errors.New("empty folder")
```

## Functions

### func [Generate](cmd/gdmd/generate.go#L33)

```go
func Generate(root string, pkg *Package) error
```

Generate creates Markdown files for the given [Package] and its nested packages.

## Types

### type [Function](cmd/gdmd/types.go#L104)

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

### func [NewFunction](cmd/gdmd/types.go#L112)

```go
func NewFunction(fset *token.FileSet, f *doc.Func) (Function, error)
```

### type [Package](cmd/gdmd/types.go#L14)

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

### func [NewPackage](cmd/gdmd/types.go#L26)

```go
func NewPackage(fset *token.FileSet, p *doc.Package, dir string, nested []Package, files []string) (Package, error)
```

### func [Parse](cmd/gdmd/parse.go#L28)

```go
func Parse(root, path string, recursive bool) (Package, error)
```

Parse walks the directory tree rooted at root and parses all .go files
it returns a [Package] for each directory containing .go files
or empty [Package] and [ErrEmpty]

### type [Position](cmd/gdmd/types.go#L98)

```go
type Position struct {
  Filename string
  Line     int
}
```

Position is a file name and line number of a declaration.

### type [Type](cmd/gdmd/types.go#L135)

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

### func [NewType](cmd/gdmd/types.go#L146)

```go
func NewType(fset *token.FileSet, t *doc.Type) (Type, error)
```

### type [Variable](cmd/gdmd/types.go#L77)

```go
type Variable struct {
  Doc   string // doc comment under the block or single declaration
  Names []string
  Src   string // piece of source code with the declaration
}
```

Variable represents constant or variable declarations within () or single one.

### func [NewVariable](cmd/gdmd/types.go#L83)

```go
func NewVariable(fset *token.FileSet, v *doc.Value) (Variable, error)
```

> You can compare [self-generated](cmd/gdmd/README.md) documentation with [pkg.dev](https://pkg.go.dev/github.com/chocolacula/gdmd).
