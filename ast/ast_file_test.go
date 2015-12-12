package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"

	"testing"
)

func TestAstFile(t *testing.T) {

	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "example.go", nil, 0)
	if err != nil {
		t.Error(err)
	}
	//f is *ast.File

	for _, d := range f.Decls {
		t.Logf("Decl:%v %T", d, d)
		switch x := d.(type) {
		case *ast.GenDecl:
			for _, s := range x.Specs {
				t.Logf(">s:%v", s)
				switch x := s.(type) {
				case *ast.TypeSpec:
					t.Logf(">s:%v Name:%v Type:%v", x, x.Name, x.Type)
					switch x := x.Type.(type) {
					case *ast.StructType:
						t.Logf(">>s:%v", x)
						for _, f := range x.Fields.List {
							switch x := f.Type.(type) {
							case *ast.Ident:
								tag := strings.TrimLeft(f.Tag.Value, "`")
								tag = strings.TrimRight(tag, "`")
								t.Logf("aa:%v bb:%v", f.Tag.Value, tag)
								st := reflect.StructTag(tag)

								t.Logf(">>>field Name:%v Type:%v Tag:%v", f.Names, x.Name, f.Tag.Value)
								t.Logf(">>>field Name:%v StructTag:%v", f.Names, st.Get("tag"))
							}
						}

					}
				}
			}

		}
	}
}
