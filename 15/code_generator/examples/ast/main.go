package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	fset = token.NewFileSet()
)

func main() {

	f, _ := parser.ParseFile(fset, "example.go", nil, 0)

	for _, node2 := range f.Decls {
		switch node1 := node2.(type) {
		case *ast.FuncDecl:
			ast.Inspect(node1, func(node ast.Node) bool {
				switch node := node.(type) {
				case *ast.BlockStmt:
					writeBlock(node, node1)
				}
				return true
			})
		}

	}

	writeToFile("example.go", f)
}

func writeBlock(node *ast.BlockStmt, node1 *ast.FuncDecl) {
	for i, v := range node.List {
		switch v := v.(type) {
		case *ast.AssignStmt:
			writeAssign(v, node, node1, i)
		}
	}
}

func writeAssign(v *ast.AssignStmt, node *ast.BlockStmt, node1 *ast.FuncDecl, i int) {
	for _, k := range v.Lhs {
		if l, ok := k.(*ast.Ident); ok {
			if l.Name == "err" {
				var returnRet []ast.Expr

				for _, r := range node1.Type.Results.List {
					switch r := r.Type.(type) {
					case *ast.Ident:
						switch r.Name {
						case "int":
							returnRet = append(returnRet, ast.NewIdent("0"))
						case "error":
							returnRet = append(returnRet, ast.NewIdent("err"))
						}

					}
				}
				body := &ast.BlockStmt{
					Lbrace: l.Pos() + 1,
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Return:  l.Pos() + 2,
							Results: returnRet,
						},
					},
					Rbrace: l.Pos() + 3,
				}

				ifs := &ast.IfStmt{
					Cond: ast.NewIdent("err != nil"),
					Body: body,
				}

				// allocate a new statement list big enough for our new statement
				j := i + 1
				list := make([]ast.Stmt, len(node.List)+1)
				copy(list, node.List[:j])
				list[j] = ifs
				copy(list[j+1:], node.List[j:])
				node.List = list

			}
		}
	}
}

func writeToFile(fname string, f *ast.File) {
	buf := &bytes.Buffer{}
	err := printer.Fprint(buf, fset, f)
	if err != nil {
		logrus.Fatal(err)
	}

	err = ioutil.WriteFile(fname, buf.Bytes(), os.ModePerm)
	if err != nil {
		logrus.Fatal("Unable to write to file", err)
	}

	/*  out, err := exec.Command("gofmt", "-w", fname).Output() */
	// if err != nil {
	// logrus.Fatal("Unable to run gofmt", fname, err, string(out))
	/* } */

	fmt.Println("–– Write to file", fname)
}
