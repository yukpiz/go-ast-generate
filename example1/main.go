package main

import (
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", f)

	file, err := os.Create("gen_example.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := format.Node(file, fset, f); err != nil {
		log.Fatal(err)
	}
}
