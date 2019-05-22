package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {
			// ノードを直で書き換える
			v.Name = &ast.Ident{
				Name: "plus",
			}
		}
		return true
	})

	// 指定したFileがあったら開く、なかったら作る。
	file, err := os.OpenFile("example/result.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// go/printer パッケージの機能でASTからソースコードを作る。
	pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	pp.Fprint(file, fset, f)
}

// 対象のノードがどのラインにあるのかを表示
//func main() {
//	fset := token.NewFileSet()
//	f, err := parser.ParseFile(fset, "samples/simple.go", nil, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	ast.Inspect(f, func(n ast.Node) bool {
//		if v, ok := n.(*ast.FuncDecl); ok {
//			fmt.Println(v.Name)
//			fmt.Println(fset.Position(v.Pos()).Line)
//		}
//		return true
//	})
//}

// importしているパッケージを表示
//func main() {
//	fset := token.NewFileSet()
//	f, err := parser.ParseFile(fset, "samples/simple.go", nil, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, d := range f.Imports {
//		ast.Print(fset, d)
//	}
//}
