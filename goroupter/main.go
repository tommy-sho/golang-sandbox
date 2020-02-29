package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

//
//func main() {
//	// flagのバリデーション
//	// https: //github.com/golang/tools/blob/3c07937fe18c27668fd78bbaed3d6b8b39e202ea/cmd/goimports/goimports.go#L280
//	//path := "./testdata/main.go"
//	//
//	//// fileのopen
//	//data, _ := fileOpen(path)
//	//
//	//// importをかける
//	////src := process(data, "test.go")
//	//src := data
//	//fset := token.NewFileSet()
//	//f, err := parser.ParseFile(fset, "sample.go", src, 0)
//	//if err != nil {
//	//	log.Fatalln("Error:", err)
//	//}
//	paths := []string{"./testdata/main.go"}
//	err := grouperMain(Env{
//		paths: paths,
//		write: true,
//	})
//	if err != nil {
//		panic(err)
//	}
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
	//
	//importMap := make(map[int]string)
	//var paths []string
	//var posSpans []Positions
	//
	//// import文をpathだけ抜き出す
	//ast.Inspect(f, func(n ast.Node) bool {
	//	if v, ok := n.(*ast.ImportSpec); ok {
	//		importMap[fset.Position(v.Path.Pos()).Line] = v.Path.Value
	//		paths = append(paths, v.Path.Value)
	//		posSpans = append(posSpans, Positions{Start: int(v.Pos()), End: int(v.End())})
	//
	//	}
	//	return true
	//})
	//
	//// 削除
	//for _, path := range paths {
	//	t, _ := strconv.Unquote(path)
	//	fmt.Println(t)
	//	astutil.DeleteImport(fset, f, t)
	//}
	//
	//// 順番を直してimportに再追加
	//so := fixImportPath(importMap)
	//for _, v := range so {
	//	astutil.AddImport(fset, f, v)
	//}
	//
	////ast.Print(fset, f) //fmt.Println(so)
	////fmt.Println(posSpans)
	////s := replaceDecl(posSpans[0], paths)
	////f.Decls[0].(*ast.GenDecl).Specs = s
	//
	//var buf bytes.Buffer
	//pp := &printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}
	//pp.Fprint(&buf, fset, f)
	//d := process(buf.Bytes(), "tges.go")
	//
	//fmt.Println(string(d))
//}

type Positions struct {
	Start int
	End   int
}



const version = "0.0.1"

var revision = "HEAD"

func main() {
	app := newApp()
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "grouper"
	app.Usage = "Force grouped import path"
	app.Version = fmt.Sprintf("%s-%s", version, revision)
	app.Authors = []*cli.Author{{
		Name:  "tommy-sho",
		Email: "tomiokasyogo@gmail.com",
	}}
	app.Action = action
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "local",
			Usage: "specify imports prefix beginning with this string after 3rd-party packages. especially your own organization name. comma-separated list",
		},
		&cli.BoolFlag{
			Name:  "write",
			Usage: "write result source to original file instead od stdout",
		},
	}

	return app
}

func action(c *cli.Context) error {
	env := Env{
		Paths:       c.Args().Slice(),
		Write:       c.Bool("write"),
		LocalPrefix: c.String("local"),
	}

	return grouperMain(env)
}
