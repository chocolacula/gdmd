# Go Doc Markdown

godoc alternative for static markdown documentation.

## Motivation

Standard for Go [pkg.go.dev](https://pkg.go.dev/) is awesome but not suitable for private repos. You can use [pkgsite](https://cs.opensource.google/go/x/pkgsite) as local alternative but it's not perfect. As stated in the [comment](https://github.com/golang/go/issues/2381#issuecomment-2183224009) of an engineer from Amazon sometimes it's a nightmare. You have to checkout a repo, run local server and open it in a browser only to read documentation.

With `gdmd` you don't even have to host your documentation, you can keep it in your repo along with the code. The generator creates a `README.md` with a package documentation in a package folder.  
You can navigate in documentation right in GitHub UI, any open directory with source code will render it's documentation.

## How to use

First, you have to

```sh
go install github.com/chocolacula/gdmd/cmd/gdmd
```

Then generate documentation for packages in a directory and all subdirectories

```sh
gdmd ./directory
```

## Example

package `main`

The single package in the project, contains data representation, parsing and generation logic.

### Index

- [Variables](#variables)
- [Functions](#functions)
  - [func Generate](#func-generate)
- [Types](#types)
  - [type Function](#type-function)
    - [func NewFunction](#func-newfunction)

### Variables

Simple error to indicate empty folder

```go
var EmptyErr = errors.New("empty folder")
```

### Functions

#### func [Generate](./cmd/gdmd/generate.go#L30)

```go
func Generate(root string, pkg *Package)
```

Generate creates a markdown files for the given [Package] and its nested packages.

### Types

#### type [Function](./cmd/gdmd/types.go#L85)

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

#### func [NewFunction](./cmd/gdmd/types.go#L93)

```go
func NewFunction(fset *token.FileSet, f *doc.Func) Function
```

> Full [documentation](cmd/gdmd/README.md)
