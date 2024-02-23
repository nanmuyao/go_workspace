package main

import "fmt"

func getMap() map[string]string {
	return nil
}

func getMap1() (m map[string]string) {
	return
}

func getMap2() (m *map[string]string) {
	return nil
}

func main() {
	fmt.Println(11)
	m := getMap()
	fmt.Println("m===", m)
	fmt.Println(m == nil)
	fmt.Println(m["1"])

	m1 := getMap1()
	fmt.Println("m1===", m1)
	fmt.Println(m1 == nil)
	fmt.Println(m1["1"])

	m2 := getMap2()
	fmt.Println("m2===", m2)
	fmt.Println(m2 == nil)
	fmt.Println((*m2)["1"])

}
