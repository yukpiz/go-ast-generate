package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	pos := fset.Position(f.Package)
	fmt.Printf("%+v\n", pos.Offset)

	ast.Fprint(os.Stdout, fset, f, nil)

	//file, err := os.Create("gen_example.go")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()

	//if err := format.Node(file, fset, f); err != nil {
	//	log.Fatal(err)
	//}
}
