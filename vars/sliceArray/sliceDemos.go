package main

import (
	"fmt"
	"reflect"
)

func test(v interface{}) string {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Slice:
		fmt.Println("is a slice with element type", rt.Elem())
		return "slice"
	case reflect.Array:
		fmt.Println("is an array with element type", rt.Elem())
		return "array"
	default:
		fmt.Println("is something else entirely")
		return ""
	}
}

func main() {
	fmt.Println("hello world")
	s0 := [6]int{1, 2, 3, 4, 5}
	fmt.Println(s0, reflect.TypeOf(s0), test(s0))
	//s0 = append(s0, 6) first argument to append must be a slice; have s0 (variable of type [6]int)
	s2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(s2, reflect.TypeOf(s2), test(s2))
	//s2 = append(s2, 6)

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s, reflect.TypeOf(s), test(s))

	s1 := make([]int, 5)
	fmt.Println(s1, reflect.TypeOf(s1), test(s))

	s3 := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Println(s3, reflect.TypeOf(s3), test(s3))

	s4 := []byte{1, 2, 3}
	fmt.Println("s4==", s4, reflect.TypeOf(s4), test(s4))
	fmt.Println()

	s5 := "abc"
	s5Slice := []byte(s5)
	fmt.Println("s5==", s5, reflect.TypeOf(s5), test(s5))
	fmt.Println("s5Slice==", s5Slice, reflect.TypeOf(s5Slice), test(s5Slice))
	for i := 0; i < len(s5Slice)/2; i++ {
		fmt.Printf("%c ", s5Slice[i])
		s5Slice[i], s5Slice[len(s5Slice)-i-1] = s5Slice[len(s5Slice)-i-1], s5Slice[i]
	}
	fmt.Println("reverse s5 slice", s5Slice)

	// slice 不会自动扩容
	s6 := make([]byte, 3, 3)
	fmt.Println("s6===", s6)
	s6 = append(s6, []byte{1, 2, 3, 4}...)
	fmt.Println("s6===", s6)
	//panic: runtime error: index out of range [7] with length 7
	for i := uint8(0); i < 10; i++ {
		s6[i] = i
	}
	fmt.Println("s6===", s6)
}
