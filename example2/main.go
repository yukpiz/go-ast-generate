package main

import (
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	file, err := os.Create("gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f := &ast.File{
		Name: ast.NewIdent("main"),
		Decls: []ast.Decl{
			&ast.FuncDecl{
				Name: ast.NewIdent("Hello"),
				Type: &ast.FuncType{},
			},
		},
	}

	if err := format.Node(file, fset, f); err != nil {
		log.Fatal(err)
	}
}
