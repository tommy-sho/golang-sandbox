package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// flagのバリデーション
	// https: //github.com/golang/tools/blob/3c07937fe18c27668fd78bbaed3d6b8b39e202ea/cmd/goimports/goimports.go#L280
	path := "./goroupter/testdata/main.gfo"

	data, _ := fileOpen(path)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "sample.go", data, 0)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	//	b, _ := pr.Process("hoge.go", data, &pr.Options{})
	ast.Print(fset, f)
	//
	//ast.SortImports(fset, f)
	//
	//pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	//pp.Fprint(os.Stdout, fset, f)
}

func fileOpen(path string) ([]byte, error) {
	switch dir, err := os.Stat(path); {
	case err != nil:
		return nil, fmt.Errorf(": %w", err)
	case dir.IsDir():
		return nil, fmt.Errorf("is dir")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("open")
	}

	return data, nil
}
