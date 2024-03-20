package main

import "fmt"

func getMap() map[string]bool {
	return nil
}

func getMap1() (m map[string]bool) {
	return
}

func getMap2() (m *map[string]bool) {
	return nil
}

func getMap4() (m map[string]bool) {
	return map[string]bool{}
}

func main() {
	//fmt.Println(11)
	m := getMap()
	fmt.Println("m===", m)
	fmt.Println("nil===", m == nil)
	fmt.Println(m["1"])

	fmt.Println("m1==========")
	m1 := getMap1()
	fmt.Println("m1===", m1)
	fmt.Println("nil==", m1 == nil)
	fmt.Println(m1["1"])

	fmt.Println("m2==========")
	m2 := getMap2()
	fmt.Println("m2===", m2)
	fmt.Println("nil==", m2 == nil)
	//fmt.Println(m2["1"])
	//fmt.Println((*m2)["1"])

	m4 := getMap4()
	fmt.Println("m4==========")
	fmt.Println("m4===", m4)
	fmt.Println("nil==", m4 == nil)
}
