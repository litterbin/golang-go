package main

type Struct struct {
	Int int
	Str string
}

func (s *Struct) String() string {
	return ""
}

func NewStruct() *Struct {
	s := &Struct{}
	return s
}
