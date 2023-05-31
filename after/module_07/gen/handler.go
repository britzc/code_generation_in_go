//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"go/format"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type Field struct {
	Name string
	Type string
	DB   string
}

type Object struct {
	Name   string
	Fields []Field
}

func main() {
	args := os.Args[1:]

	objectName := args[0]
	schemaPath := args[1]
	templatePath := args[2]
	outputPath := args[3]

	funcMap := template.FuncMap{
		"toLower": strings.ToLower,
	}

	tmpl := template.Must(template.New("handler").Funcs(funcMap).ParseFiles(templatePath))

	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		panic(err)
	}
	defer schemaFile.Close()

	fields := make([]Field, 0)

	scanner := bufio.NewScanner(schemaFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		fields = append(fields, Field{Name: parts[0], Type: parts[1], DB: parts[2]})
	}

	object := Object{
		Name:   objectName,
		Fields: fields,
	}

	templateName := filepath.Base(templatePath)

	var tmplData bytes.Buffer
	if err := tmpl.ExecuteTemplate(&tmplData, templateName, object); err != nil {
		panic(err)
	}

	fmtData, err := format.Source(tmplData.Bytes())
	if err != nil {
		panic(err)
	}

	codeFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer codeFile.Close()

	codeFile.Write(fmtData)
}
