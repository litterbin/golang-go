package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

func Benchmark_parseUintBytes(b *testing.B) {
	b.ReportAllocs()

	a := []byte("1")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		n, err := parseUintBytes(a, 0, 16)
		if err != nil {
			b.Fatal(err)
		}
		if n != 1 {
			b.Fatal(fmt.Errorf("n != 1"))
		}

	}
}

func Benchmark_parseUintStrconv(b *testing.B) {
	b.ReportAllocs()

	a := []byte("1")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n, err := strconv.ParseUint(string(a), 0, 16)
		if err != nil {
			b.Fatal(err)
		}
		if n != 1 {
			b.Fatal(fmt.Errorf("n != 1 "))
		}
	}
}

func Benchmark_parseUintStrUnsafe(b *testing.B) {

	b.ReportAllocs()

	a := []byte("1")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n, err := strconv.ParseUint(unsafeString(a), 0, 16)
		if err != nil {
			b.Fatal(err)
		}
		if n != 1 {
			b.Fatal(fmt.Errorf("n != 1 "))
		}
	}
}

func Test_parseUintBytes(t *testing.T) {
	a := []byte("1")
	t.Logf("a:%v", a)
	n, err := parseUintBytes(a, 0, 16)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("n:%v", n)

}

func unsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// parseUintBytes is like strconv.ParseUint, but using a []byte.
func parseUintBytes(s []byte, base int, bitSize int) (n uint64, err error) {
	var cutoff, maxVal uint64

	if bitSize == 0 {
		bitSize = int(strconv.IntSize)
	}

	s0 := s
	switch {
	case len(s) < 1:
		err = strconv.ErrSyntax
		goto Error

	case 2 <= base && base <= 36:
		// valid base; nothing to do

	case base == 0:
		// Look for octal, hex prefix.
		switch {
		case s[0] == '0' && len(s) > 1 && (s[1] == 'x' || s[1] == 'X'):
			base = 16
			s = s[2:]
			if len(s) < 1 {
				err = strconv.ErrSyntax
				goto Error
			}
		case s[0] == '0':
			base = 8
		default:
			base = 10
		}

	default:
		err = errors.New("invalid base " + strconv.Itoa(base))
		goto Error
	}

	n = 0
	cutoff = cutoff64(base)
	maxVal = 1<<uint(bitSize) - 1

	for i := 0; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		default:
			n = 0
			err = strconv.ErrSyntax
			goto Error
		}
		if int(v) >= base {
			n = 0
			err = strconv.ErrSyntax
			goto Error
		}

		if n >= cutoff {
			// n*base overflows
			n = 1<<64 - 1
			err = strconv.ErrRange
			goto Error
		}
		n *= uint64(base)

		n1 := n + uint64(v)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			n = 1<<64 - 1
			err = strconv.ErrRange
			goto Error
		}
		n = n1
	}

	return n, nil

Error:
	return n, &strconv.NumError{Func: "ParseUint", Num: string(s0), Err: err}
}

// Return the first number n such that n*base >= 1<<64.
func cutoff64(base int) uint64 {
	if base < 2 {
		return 0
	}
	return (1<<64-1)/uint64(base) + 1
}
