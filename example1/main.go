package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	pos := fset.Position(f.Package)
	fmt.Printf("%+v\n", pos.Offset)

	for _, imp := range f.Imports {
		pos = fset.Position(imp.Pos())
		fmt.Printf("%+v\n", pos.Offset)
	}

	//file, err := os.Create("gen_example.go")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()

	//if err := format.Node(file, fset, f); err != nil {
	//	log.Fatal(err)
	//}
}
