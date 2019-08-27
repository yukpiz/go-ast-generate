package main

import (
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

var GEN_FILE_PATH = "cmd/gen.go"

func main() {
	fset := token.NewFileSet()
	file, err := os.Create(GEN_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := format.Node(file, fset, basicAST); err != nil {
		log.Fatal(err)
	}

	f, err := parser.ParseFile(fset, GEN_FILE_PATH, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", f)
}
