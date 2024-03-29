package enum

import (
	"bufio"
	"bytes"
	"text/template"
)

// Generate parses and executes the templates using the data set in Options,
// returning the generated code
func Generate(o *Options) ([]byte, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	w.WriteString("// Code generated by github.com/harry93848bb7/enum-codegen DO NOT EDIT.\n")

	root, err := ParseTemplates(template.New("root"))
	if err != nil {
		return nil, err
	}
	if err := root.ExecuteTemplate(w, "enum_package.tmpl", o); err != nil {
		return nil, err
	}
	// walk over all the types define and build up the buffer
	for _, t := range o.Types {
		if err := root.ExecuteTemplate(w, "enum_type.tmpl", t); err != nil {
			return nil, err
		}
		if err := root.ExecuteTemplate(w, "enum.tmpl", t); err != nil {
			return nil, err
		}
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GenerateTests parses and executes the templates using the data set in Options,
// returning the generated code tests
func GenerateTests(o *Options) ([]byte, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	w.WriteString("// Code generated by github.com/harry93848bb7/enum-codegen DO NOT EDIT.\n")

	root, err := ParseTemplates(template.New("root"))
	if err != nil {
		return nil, err
	}
	if err := root.ExecuteTemplate(w, "test_package.tmpl", o); err != nil {
		return nil, err
	}
	for _, t := range o.Types {
		if err := root.ExecuteTemplate(w, "test.tmpl", t); err != nil {
			return nil, err
		}
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
