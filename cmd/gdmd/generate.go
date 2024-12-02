package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func generateOne(root string, tmpl *template.Template, pkg *Package) error {
	filename := filepath.Join(root, pkg.Dir, "README.md")
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, pkg)
	if err != nil {
		return err
	}

	for _, nstd := range pkg.Nested {
		err = generateOne(root, tmpl, &nstd)
		if err != nil {
			return err
		}
	}
	return nil
}

// Generate creates Markdown files for the given [Package] and its nested packages.
func Generate(root string, pkg *Package) error {
	funcs := template.FuncMap{
		"ToLower": strings.ToLower,
	}
	tmpl, err := template.
		New("markdown").
		Funcs(funcs).
		Parse(templateData)
	if err != nil {
		return err
	}
	return generateOne(root, tmpl, pkg)
}
