package main

import (
	"go/types"
	"strings"
	"testing"

	"gopkg.in/src-d/go-parse-utils.v1"
)

func TestImporter(t *testing.T) {
	p := "."
	pkg, err := scanPackage(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("pkg.Name:%v", pkg.Name())
	t.Logf("pkg.Scope.Names:%v", pkg.Scope().Names())
	objs := objectsInScope(pkg.Scope())
	for _, obj := range objs {
		/*
			t.Logf("obj.ID:%v", obj.Id())
			t.Logf("obj.Name:%v", obj.Name())
			t.Logf("obj.Type:%T", obj.Type())
		*/

		s := scanStruct(obj)
		if s == nil {
			continue
		}

		fields := scanStructFields(s)
		for _, f := range fields {
			t.Logf("field %v %v", f, f.Type())
		}
	}
}

func objectsInScope(scope *types.Scope) (objs []types.Object) {
	for _, n := range scope.Names() {
		obj := scope.Lookup(n)
		objs = append(objs, obj)

		typ := obj.Type()
		if _, ok := typ.Underlying().(*types.Struct); ok {
			objs = append(objs, methodsForType(types.NewPointer(typ))...)
		}
	}
	return
}

func methodsForType(typ types.Type) (objs []types.Object) {
	methods := types.NewMethodSet(typ)

	for i := 0; i < methods.Len(); i++ {
		objs = append(objs, methods.At(i).Obj())
	}
	return
}

func scanPackage(p string) (*types.Package, error) {
	importer := parseutil.NewImporter()
	pkg, err := importer.ImportWithFilters(
		p,
		parseutil.FileFilters{
			func(pkg, file string, typ parseutil.FileType) bool {
				return !strings.HasSuffix(file, ".sone.go")
			},
		},
	)
	return pkg, err
}

func scanStruct(o types.Object) *types.Struct {
	switch o.Type().(type) {
	case *types.Named:
		if s, ok := o.Type().Underlying().(*types.Struct); ok {
			return s
		}
	}
	return nil
}

func scanStructFields(s *types.Struct) (fields []*types.Var) {
	for i := 0; i < s.NumFields(); i++ {
		v := s.Field(i)
		fields = append(fields, v)
	}
	return fields
}
