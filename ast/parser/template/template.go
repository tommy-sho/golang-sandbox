package template

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type Data struct {
	Struct []Struct
}

type Struct struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
}

func New(data Data) error {
	fread, err := ioutil.ReadFile("./template/temp.tmol")
	if err != nil {
		return fmt.Errorf("failed read file: %w", err)
	}

	t := template.Must(template.New("model").Parse(string(fread)))
	if err := t.Execute(os.Stdout, data); err != nil {
		return fmt.Errorf("faield Exec termplate: %w", err)
	}

	return nil
}
