# Go Doc Markdown

godoc alternative for static markdown documentation.

## Motivation

Standard [pkg.go.dev](https://pkg.go.dev/) is awesome but not suitable for private repos. You can use [pkgsite](https://cs.opensource.google/go/x/pkgsite) as local alternative but it's not perfect. Both occasionally skip certain files and generate incomplete documentation.

As stated in the [comment](https://github.com/golang/go/issues/2381#issuecomment-2183224009) of an engineer from Amazon sometimes it's a nightmare. You have to checkout a repo, run local server and open it in a browser only to read documentation.

With `gdmd` you don't even have to host your documentation, you can keep it in your repo along with the code. The generator creates a `README.md` with a package documentation in a package folder.  
You can navigate through the documentation right in GitHub UI, any open directory with source code will render its documentation.

## How to use

First, you have to

```sh
go install github.com/chocolacula/gdmd/cmd/gdmd
```

Then generate documentation for a package in a directory

```sh
gdmd ./directory
```

## Example

package `main`

The single package in the project, contains data representation, parsing and generation logic.

## Index

- [Variables](#variables)
- [Functions](#functions)
  - [func Generate(root string, pkg *Package)](#func-generate)
- [Types](#types)
  - [type Function](#type-function)
    - [func NewFunction(fset *token.FileSet, f *doc.Func) Function](#func-newfunction)

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

Generate creates markdown files for the given [Package] and its nested packages.

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

> You can compare [self-generated](cmd/gdmd/README.md) documentation with [pkg.dev](https://pkg.go.dev/github.com/chocolacula/gdmd).
