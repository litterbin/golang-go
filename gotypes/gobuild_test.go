package main

import (
	"go/build"
	"os"
	"testing"
)

func TestGoBuildImport(t *testing.T) {
	path := "."
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	pkg, err := build.Default.Import(path, wd, build.ImportComment)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("dir:%v name:%v root:%v srcroot:%v", pkg.Dir, pkg.Name, pkg.Root, pkg.SrcRoot)
	t.Logf("GoFiles:%v", pkg.GoFiles)
	t.Logf("imports:%v", pkg.Imports)

}
