package main

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

const hello = `package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
`

func TestPkgInfo(t *testing.T) {
	fset := token.NewFileSet()

	// Parse the input string,[]byte,io.Reader
	// ParseFile returns an *ast.File,a syntax tree
	f, err := parser.ParseFile(fset, "hello.go", hello, 0)
	if err != nil {
		t.Fatal(err)
	}

	// A Config controls various options of the type checker.
	// The defaults work fine except for one settings:
	// we must specify how to deal with imports
	conf := types.Config{Importer: importer.Default()}

	// Type-check the package containing only file f
	// Check returns a *types.Package
	pkg, err := conf.Check("aaa", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Package  %q\n", pkg.Path())
	t.Logf("Name:    %s\n", pkg.Name())
	t.Logf("Imports: %s\n", pkg.Imports())
	t.Logf("Scope:   %s\n", pkg.Scope())
}
