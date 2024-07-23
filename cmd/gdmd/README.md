# Overview

package `main`

The single package in the project, contains data representation, parsing and generation logic.

## Index

- [Variables](#variables)
- [Functions](#functions)
  - [func Generate](#func-generate)
- [Types](#types)
  - [type A](#type-a)
  - [type Function](#type-function)
  - [type Package](#type-package)
  - [type Position](#type-position)
  - [type Type](#type-type)
  - [type Variable](#type-variable)
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

### type [A](./parse.go#L25)

```go
type A struct {
}
```

### type [Function](./types.go#L85)

```go
type Function struct {
  Doc       string
  Name      string
  Pos       Position
  Signature string
}
```

Function represents a function or method declaration.

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

### type [Position](./types.go#L79)

```go
type Position struct {
  Filename string
  Line     int
}
```

Position is a file name and line number of a declaration.

### type [Type](./types.go#L107)

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

### type [Variable](./types.go#L61)

```go
type Variable struct {
  Doc   string // doc comment under the block or single declaration
  Names []string
  Src   string // piece of source code with the declaration
}
```

Variable represents constant or variable declarations within () or single one.


## Source files

[generate.go](./generate.go)
[main.go](./main.go)
[parse.go](./parse.go)
[template.go](./template.go)
[types.go](./types.go)
