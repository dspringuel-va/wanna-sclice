package wanna_slice

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ArraySuite struct {
	suite.Suite
}

func (s *ArraySuite) SetupTest() {
	fmt.Printf("------\n%s\n\n", s.T().Name())
}
func (s *ArraySuite) TearDownTest() {
	fmt.Printf("------\n")
}

func (s *ArraySuite) Test_Array_Uninitialized() {
	var a [2]string
	fmt.Printf("a: %#v\n", a)
}

func (s *ArraySuite) Test_Array_Empty() {
	a := [2]string{}
	fmt.Printf("a: %#v\n", a)
}

func (s *ArraySuite) Test_Array_HalfInitialized() {
	a := [2]string{"a"}
	fmt.Printf("a: %#v\n", a)
}

func (s *ArraySuite) Test_Array_Len() {
	a := [2]string{"a", "b"}
	fmt.Printf("len(a): %#v\n", len(a))
}

func (s *ArraySuite) Test_Array_Cap() {
	a := [2]string{"a", "b"}
	fmt.Printf("len(a): %#v\n", cap(a))
}

func (s *ArraySuite) Test_Array_LenEmpty() {
	a := [2]string{}
	fmt.Printf("len(a): %#v\n", len(a))
}

func (s *ArraySuite) Test_Array_make() {
	// a := make([2]string, 2) doesn't compile
	// fmt.Printf("a: %#v\n", a)
}

func (s *ArraySuite) Test_Array_append() {
	a := [2]string{"a", "b"}
	// a = append(a, "c") doesn't compile
	fmt.Printf("a: %#v\n", a)
}

func (s *ArraySuite) Test_Array_Equality() {
	a := [2]string{"a", "b"}
	b := a
	fmt.Printf("a == b: %v\n", a == b)
}

func (s *ArraySuite) Test_Array_Equality_DifferentObject() {
	a := [2]string{"a", "b"}
	b := [2]string{"a", "b"}
	fmt.Printf("a == b: %v\n", a == b)
}

func (s *ArraySuite) Test_Array_Equality_AfterModification() {
	a := [2]string{"a", "b"}
	b := a
	b[0] = "c"
	fmt.Printf("a == b: %v\n", a == b)
}

func (s *ArraySuite) Test_Array_PointerEquality() {
	a := [2]string{"a", "b"}
	b := a

	fmt.Printf("&a == &b: %v\n", &a == &b)
}

func (s *ArraySuite) Test_Array_PointerInArray() {
	a := [2]string{"a", "b"}

	fmt.Printf("&a == &a[0]: %v\n", fmt.Sprintf("%p", &a) == fmt.Sprintf("%p", &a[0]))
	fmt.Printf("&a == &(a[0]): %v\n", fmt.Sprintf("%p", &a) == fmt.Sprintf("%p", &(a[0])))
}

func (s *ArraySuite) Test_Array_PointerDereferencedEquality() {
	a := [2]string{"a", "b"}
	p := &a
	b := *p

	fmt.Printf("a == b: %v\n", a == b)
	fmt.Printf("&a == &b: %v\n", &a == &b)
}


func (s *ArraySuite) Test_Array_PassArrayToFunction() {
	a := [2]string{"a", "b"}

	f := func(b [2]string) {
		b[0] = "c"
	}
	f(a)

	fmt.Printf("a: %#v\n", a)
}

func (s *ArraySuite) Test_Array_PassPointerArrayToFunction() {
	a := [2]string{"a", "b"}

	f := func(b *[2]string) {
		b[0] = "c"
	}
	f(&a)

	fmt.Printf("a: %#v\n", a)
}

func (s *ArraySuite) Test_Array_InStruct() {
	type Struct struct {
		array [2]string
	}

	a := Struct{}

	fmt.Printf("a: %#v\n", a)
}

func (s *ArraySuite) Test_Array_InStruct_Assignment() {
	type Struct struct {
		array [2]string
	}
	a := [2]string {"a", "b"}

	b := Struct{array: a}
	c := b

	a[0] = "c"
	fmt.Printf("b: %#v\n", b)
	fmt.Printf("c: %#v\n", c)
}



func Test_Arrays(t *testing.T) {
	suite.Run(t, &ArraySuite{})
}
