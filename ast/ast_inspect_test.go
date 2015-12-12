package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAstInspect(t *testing.T) {

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "example.go", nil, 0)
	if err != nil {
		t.Error(err)
		return
	}

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
		}
		if s != "" {
			t.Logf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})

}
