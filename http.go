package main

import (
	"embed"
	_ "embed"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

//go:embed templates
var templates embed.FS

func compileTemplates(dir string) (*template.Template, error) {
	// const fun = "compileTemplates"
	tpl := template.New("")
	// Since filepath.Walk only handles filesystem directories, we use the new
	// and optimized fs.WalkDir introduced in Go 1.16, which takes an fs.FS.
	err := fs.WalkDir(templates, dir, func(path string, info fs.DirEntry, err error) error {
		// Skip non-templates.
		if info.IsDir() || !strings.HasSuffix(path, ".tpl") {
			return nil
		}
		// Load file from embed virtual file, or use the shortcut
		// templates.ReadFile(path).
		f, _ := templates.Open(path)
		// Now read it.
		sl, _ := io.ReadAll(f)
		// It can now be parsed as a string.
		tpl.Parse(string(sl))
		return nil
	})
	return tpl, err
}
