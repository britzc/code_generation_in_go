//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"go/format"
	"html/template"
	"os"
	"strings"
)

type Field struct {
	Name string
	Type string
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

	tmpl := template.Must(template.ParseFiles(templatePath))

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

		fields = append(fields, Field{Name: parts[0], Type: parts[1]})
	}

	object := Object{
		Name:   objectName,
		Fields: fields,
	}

	var tmplData bytes.Buffer
	tmpl.Execute(&tmplData, object)

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
