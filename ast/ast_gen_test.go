package main

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"testing"
)

func TestASTGen(t *testing.T) {
	var decls []ast.Decl
	var name = "gen"

	decls = append(decls, declMethod("Encode"))

	file := buildFile(name, decls)
	if file == nil {
		t.Fatal(file)
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, token.NewFileSet(), file)
	t.Logf("output: %v", buf.String())
}

func declMethod(name string) ast.Decl {
	return &ast.FuncDecl{
		Recv: fields(field("s", ptr(ast.NewIdent(name)))),
		Name: ast.NewIdent(name),
		Type: methodType(),
		Body: declMethodBody(),
	}
}

func declMethodBody() *ast.BlockStmt {
	body := declMethodBodyBase()
	return body
}

func declMethodBodyBase() *ast.BlockStmt {
	return &ast.BlockStmt{
		List: []ast.Stmt{},
	}
}

func methodType() *ast.FuncType {
	return &ast.FuncType{
		Params: fields(
			field("r", ast.NewIdent("io.Reader")),
		),
		Results: fields(
			field("", ast.NewIdent("error")),
		),
	}
}

func buildFile(name string, decls []ast.Decl) *ast.File {
	f := &ast.File{
		Name: ast.NewIdent(name),
	}
	f.Decls = append(f.Decls, decls...)
	return f
}

func fields(fields ...*ast.Field) *ast.FieldList {
	return &ast.FieldList{List: fields}
}
func field(name string, t ast.Expr) *ast.Field {
	var names []*ast.Ident
	if name != "" {
		names = append(names, ast.NewIdent(name))
	}
	return &ast.Field{
		Names: names,
		Type:  t,
	}
}

func ptr(expr ast.Expr) ast.Expr {
	return &ast.StarExpr{X: expr}
}
