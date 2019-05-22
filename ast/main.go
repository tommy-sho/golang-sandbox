package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
)

var s = "rrr"

// sample.go
const src = `package main

func main() {
	msg := "Hello,"
	msg += " World"
	println(msg)
}`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "sample.go", src, 0)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	ast.Print(fset, f)

	conf := &types.Config{
		Importer: importer.Default(),
	}

	// 入れ物を作成している
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}

	_, err = conf.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// 47バイト目 = msg
	const p = 47
	var pos token.Pos
	fset.Iterate(func(f *token.File) bool {
		if f.Name() == "sample.go" {
			pos = f.Pos(p)
			return false
		}
		return true
	})

	//　この部分で目的の識別子を探してる
	var ident *ast.Ident
	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		// 関数の名前を抽出するやつ
		case *ast.FuncDecl:
			fmt.Println("This is Ident", n.(*ast.FuncDecl).Name)
		case *ast.IndexExpr:
			fmt.Println("This is IndexExpr", n.(*ast.IndexExpr).X)
		case *ast.Ident:
			fmt.Println("This is Indent", n.(*ast.Ident).Name)
		//	fmt.Println("   This is parent", n.(*ast.Ident).Obj.Kind.String())
		case *ast.CallExpr:
			fmt.Println("This is Function", n.(*ast.CallExpr).Fun.(*ast.Ident).Name)
		}
		if n == nil || pos < n.Pos() || pos > n.End() {
			return true
		}

		if n, ok := n.(*ast.Ident); ok {
			ident = n
			return false
		}

		return true
	})

	from := ident.Name
	const to = "message"
	fmt.Println(ident, "->", to)

	obj := info.Defs[ident]
	if obj == nil {
		obj = info.Uses[ident]
	}

	// 宣言している部分を探している
	fmt.Println("== Defs ==")
	for i, o := range info.Defs {
		if i.Name == from && o.Parent() == obj.Parent() {
			fmt.Println(fset.Position(i.Pos()))
			i.Name = to
		}
	}

	// 宣言されたものが使用されている部分を探している
	fmt.Println("== Uses ==")
	for i, o := range info.Uses {
		if i.Name == from && o.Parent() == obj.Parent() {
			fmt.Println(fset.Position(i.Pos()))
			i.Name = to
		}
	}

	fmt.Println()
	fmt.Println("==== Before ====")
	fmt.Println(src)

	fmt.Println()
	fmt.Println("==== After ====")
	format.Node(os.Stdout, fset, f)
}
