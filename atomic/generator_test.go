package generator

import (
	"testing"
)

const TestCount = 1000000

func TestClosure(t *testing.T) {
	cl := Closure()

	for i := 1; i < TestCount; i++ {
		actual := cl()
		expected := i
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
}

func TestMutexClosure(t *testing.T) {
	cl := MutexClosure()

	for i := 1; i < TestCount; i++ {
		actual := cl()
		expected := i
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
}

func TestAtomicClosure(t *testing.T) {
	cl := AtomicClosure()

	var i int32
	for i = 1; i < TestCount; i++ {
		actual := cl()
		expected := i
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
}

func TestChannel(t *testing.T) {
	ch := Channel()

	for i := 1; i < TestCount; i++ {
		actual := <-ch
		expected := i
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
		}
	}
}

func BenchmarkClosure(b *testing.B) {
	cl := Closure()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cl()
	}
}

func BenchmarkMutexClosure(b *testing.B) {
	cl := MutexClosure()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cl()
	}
}

func BenchmarkAtomicClosure(b *testing.B) {
	cl := AtomicClosure()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cl()
	}
}

func BenchmarkChannel(b *testing.B) {
	ch := Channel()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		<-ch
	}
}
