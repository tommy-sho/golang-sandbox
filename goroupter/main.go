package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"

	"golang.org/x/tools/container/intsets"

	"golang.org/x/tools/imports"

	"io/ioutil"
	"log"
	"os"
)

func main() {
	// flagのバリデーション
	// https: //github.com/golang/tools/blob/3c07937fe18c27668fd78bbaed3d6b8b39e202ea/cmd/goimports/goimports.go#L280
	path := "./testdata/main.gfo"

	// fileのopen
	data, _ := fileOpen(path)

	// importをかける
	//src := process(data, "test.go")
	src := data
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "sample.go", src, 0)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	//
	//mergeImports(&ProcessEnv{LocalPrefix: ""}, fset, f)
	//
	//imps := astutil.Imports(fset, f)
	//var spacesBefore []string // import paths we need spaces before
	//for _, impSection := range imps {
	//	// Within each block of contiguous imports, see if any
	//	// import lines are in different group numbers. If so,
	//	// we'll need to put a space between them so it's
	//	// compatible with gofmt.
	//	lastGroup := -1
	//	for _, importSpec := range impSection {
	//		importPath, _ := strconv.Unquote(importSpec.Path.Value)
	//		groupNum := importGroup(&ProcessEnv{LocalPrefix: ""}, importPath)
	//		if groupNum != lastGroup && lastGroup != -1 {
	//			spacesBefore = append(spacesBefore, importPath)
	//		}
	//		lastGroup = groupNum
	//	}
	//
	//}
	//fmt.Println(spacesBefore)
	//
	//printerMode := printer.UseSpaces
	//printConfig := &printer.Config{Mode: printerMode, Tabwidth: 0}
	//
	//var buf bytes.Buffer
	//err = printConfig.Fprint(&buf, fset, f)
	//if err != nil {
	//	panic(err)
	//}
	//
	//out := buf.Bytes()
	//
	//if len(spacesBefore) > 0 {
	//	out, err = addImportSpaces(bytes.NewReader(out), spacesBefore)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//
	//out, err = format.Source(out)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(out))
	//sortImports(&ProcessEnv{LocalPrefix: ""}, fset, f)
	//ast.Print(fset, f)
	//pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	//pp.Fprint(os.Stdout, fset, f)

	importMap := make(map[int]string)
	var paths []string
	var posSpans []Positions

	// import文をpathだけ抜き出す
	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.ImportSpec); ok {
			importMap[fset.Position(v.Path.Pos()).Line] = v.Path.Value
			paths = append(paths, v.Path.Value)
			posSpans = append(posSpans, Positions{Start: int(v.Pos()), End: int(v.End())})

		}
		return true
	})

	// 削除
	for _, path := range paths {
		t, _ := strconv.Unquote(path)
		fmt.Println(t)
		astutil.DeleteImport(fset, f, t)
	}

	// 順番を直してimportに再追加
	so := fixImportPath(importMap)
	for _, v := range so {
		astutil.AddImport(fset, f, v)
	}

	//ast.Print(fset, f) //fmt.Println(so)
	//fmt.Println(posSpans)
	//s := replaceDecl(posSpans[0], paths)
	//f.Decls[0].(*ast.GenDecl).Specs = s

	var buf bytes.Buffer
	pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	pp.Fprint(&buf, fset, f)
	d := process(buf.Bytes(), "tges.go")

	fmt.Println(string(d))
}

type Positions struct {
	Start int
	End   int
}

func fixImportPath(paths map[int]string) map[int]string {
	min := min(paths)
	res := make(map[int]string, len(paths))
	for _, path := range paths {
		t, _ := strconv.Unquote(path)
		res[min] = t
		min++
	}

	return res
}

func min(paths map[int]string) int {
	min := intsets.MaxInt
	for k := range paths {
		if k < min {
			min = k
		}
	}

	return min
}

func process(data []byte, name string) []byte {
	b, _ := imports.Process("hoge.go", data, &imports.Options{})
	return b
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
