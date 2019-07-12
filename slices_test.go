package wanna_slice

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type SliceSuite struct {
	suite.Suite
}

func (s *SliceSuite) SetupTest() {
	fmt.Printf("------\n%s\n\n", s.T().Name())
}
func (s *SliceSuite) TearDownTest() {
	fmt.Printf("------\n")
}

func (s *SliceSuite) Test_Slice_Uninitialized() {
	var a []string
	fmt.Printf("a: %#v\n", a)
}

func (s *SliceSuite) Test_Slice_Empty() {
	a := []string{}
	fmt.Printf("a: %#v\n", a)
}

func (s *SliceSuite) Test_Slice_Initialized() {
	a := []string{"a"}
	fmt.Printf("a: %#v\n", a)
}

func (s *SliceSuite) Test_Slice_Len() {
	a := []string{"a", "b"}
	fmt.Printf("len(a): %#v\n", len(a))
}

func (s *SliceSuite) Test_Slice_Cap() {
	a := []string{"a", "b"}
	fmt.Printf("cap(a): %#v\n", cap(a))
}

func (s *SliceSuite) Test_Slice_LenEmpty() {
	a := []string{}
	fmt.Printf("len(a): %#v\n", len(a))
}

func (s *SliceSuite) Test_Slice_make() {
	a := make([]string, 2)
	fmt.Printf("a: %#v\n", a)
}

func (s *SliceSuite) Test_Slice_append() {
	a := []string{"a", "b"}
	a = append(a, "c")
	fmt.Printf("a: %#v\n", a)
}

func (s *SliceSuite) Test_Slice_Equality() {
	a := []string{"a", "b"}
	b := a
	// fmt.Printf("a == b: %v\n", a == b) doesn't compile
	fmt.Printf("a == b: %v\n", reflect.DeepEqual(a, b))
}

func (s *SliceSuite) Test_Slice_Equality_AfterModification() {
	a := []string{"a", "b"}
	b := a
	b[0] = "c"
	fmt.Printf("a == b: %v\n", reflect.DeepEqual(a, b))
}

func (s *SliceSuite) Test_Slice_PointerEquality() {
	a := []string{"a", "b"}
	b := a

	fmt.Printf("&a == &b: %v\n", &a == &b)
}

func (s *SliceSuite) Test_Slice_PointerInArray() {
	a := []string{"a", "b"}

	fmt.Printf("&a == &a[0]: %v\n", fmt.Sprintf("%p", &a) == fmt.Sprintf("%p", &a[0]))
	fmt.Printf("&a == &(a[0]): %v\n", fmt.Sprintf("%p", &a) == fmt.Sprintf("%p", &(a[0])))
	fmt.Printf("a == &a[0]: %v\n", fmt.Sprintf("%p", a) == fmt.Sprintf("%p", &a[0]))
}

func (s *SliceSuite) Test_Slice_PointerDereferencedEquality() {
	a := []string{"a", "b"}
	p := &a
	b := *p

	fmt.Printf("a == b: %v\n", reflect.DeepEqual(a, b))
	fmt.Printf("&a == &b: %v\n", &a == &b)
}


func (s *SliceSuite) Test_Slice_PassArrayToFunction() {
	a := []string{"a", "b"}

	f := func(b []string) {
		b[0] = "c"
	}
	f(a)

	fmt.Printf("a: %#v\n", a)
}

func (s *SliceSuite) Test_Slice_PassPointerArrayToFunction() {
	a := []string{"a", "b"}

	f := func(b *[]string) {
		// a[0] = "c" doesn't compile
		(*b)[0] = "c"
	}
	f(&a)

	fmt.Printf("a: %#v\n", a)
}

func (s *SliceSuite) Test_Slice_InStruct() {
	type Struct struct {
		array []string
	}

	a := Struct{}

	fmt.Printf("a: %#v\n", a)
}

func (s *SliceSuite) Test_Slice_InStruct_Assignment() {
	type Struct struct {
		array []string
	}
	a := []string {"a", "b"}

	b := Struct{array: a}
	c := b

	a[0] = "c"
	fmt.Printf("b: %#v\n", b)
	fmt.Printf("c: %#v\n", c)
}


func Test_Slices(t *testing.T) {
	suite.Run(t, &SliceSuite{})
}
