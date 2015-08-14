package main

import (
	"testing"
)

func afunc(a int) int {
	return a + 1
}

func InterfaceTest(t *testing.T) {
	var a interface{} = 1

	//cannot use a (type interface {}) as type int in argument to afunc: need type assertion
	b := afunc(a.(int))
	if b != 2 {
		t.Error(b)
	}

}
