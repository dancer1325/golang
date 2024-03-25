package template_test

import (
	"fmt"
	"os"
	"testing"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func TestWhiteSpace(t *testing.T) {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	// TODO: Why NOTHING displayed by console?
	fmt.Print(tmpl.Name())
}
