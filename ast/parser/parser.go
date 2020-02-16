package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"github.com/tommy-sho/golang-sandbox/ast/parser/template"
)

// ファイルかstructの情報のみを取り出す
func DetectStruct(node ast.Node) template.Data {
	list := make([]template.Struct, 0)
	ast.Inspect(node, func(n ast.Node) bool {
		if vv, ok := n.(*ast.TypeSpec); ok {
			if v, ok := vv.Type.(*ast.StructType); ok {
				s := template.Struct{
					Name: vv.Name.Name,
				}

				fmt.Println("-----", vv.Name)
				fff := make([]template.Field, 0, len(v.Fields.List))
				for _, v := range v.Fields.List {
					fmt.Printf("%#v\n", v.Type)
					vvv, _ := v.Type.(*ast.Ident)
					f := template.Field{
						Name: v.Names[0].Obj.Name,
						Type: vvv.Name,
					}
					fff = append(fff, f)
				}
				s.Fields = fff
				list = append(list, s)
			}
		}
		return true
	})
	return template.Data{Struct: list}
}

func main() {
	fileName := "./text.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.Mode(0))
	if err != nil {
		log.Fatal(err)
	}
	s := DetectStruct(f)

	if err := template.New(s); err != nil {
		log.Fatal(err)
	}
}
