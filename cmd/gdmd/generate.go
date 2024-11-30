package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func generateOne(root string, tmpl *template.Template, pkg *Package) {
	filename := filepath.Join(root, pkg.Dir, "README.md")
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, pkg)
	if err != nil {
		panic(err)
	}

	for _, nstd := range pkg.Nested {
		generateOne(root, tmpl, &nstd)
	}
}

// Generate creates Markdown files for the given [Package] and its nested packages.
func Generate(root string, pkg *Package) {
	funcs := template.FuncMap{
		"ToLower": strings.ToLower,
	}
	tmpl, err := template.
		New("markdown").
		Funcs(funcs).
		Parse(templateData)
	if err != nil {
		log.Fatal(err)
	}
	generateOne(root, tmpl, pkg)
}
