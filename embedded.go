package enum

import (
	"embed"
	"text/template"
)

//go:embed templates/*
var embeddedTemplates embed.FS

// Parse parses declared templates
func ParseTemplates(t *template.Template) (*template.Template, error) {
	directory, err := embeddedTemplates.ReadDir("templates")
	if err != nil {
		return nil, err
	}
	for _, entry := range directory {
		b, err := embeddedTemplates.ReadFile("templates/" + entry.Name())
		if err != nil {
			return nil, err
		}
		var tmpl *template.Template
		if t == nil {
			t = template.New(entry.Name())
		}
		if entry.Name() == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(entry.Name())
		}
		if _, err := tmpl.Parse(string(b)); err != nil {
			return nil, err
		}
	}
	return t, nil
}
