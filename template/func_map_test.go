package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func TestSimple(t *testing.T) {
	sweaters := Inventory{"wool", 17}

	tmpl, err := template.New("test").Parse("{{ .Count }} ")
	if err != nil {
		t.Error(err)
	}

	result := bytes.NewBufferString("")

	err = tmpl.Execute(result, sweaters)
	if err != nil {
		t.Error(err)
	}

	resultInt, _ := strconv.Atoi(strings.TrimSpace(result.String()))
	if resultInt != int(sweaters.Count) {
		t.Errorf("%v|%v", resultInt, sweaters.Count)
	}

}

func TestFuncWithTmplFile(t *testing.T) {
	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	tmpl, err := template.New("test.tmpl").Funcs(funcMap).ParseFiles("test.tmpl")
	if err != nil {
		t.Error(err)
	}

	result := bytes.NewBufferString("")
	err = tmpl.Execute(result, "the book")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}

}
