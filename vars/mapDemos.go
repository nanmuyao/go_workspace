package main

import "fmt"

func getMap() map[string]int {
	return map[string]int{}
}

func initMap() {
	var m map[string]int
	//panic: assignment to entry in nil map
	//m["hello"] = 1
	m = getMap()
	m["hello"] = 1
	fmt.Println(m)
}

func initMap1() {
	var m map[string]int = map[string]int{}
	m["hello"] = 1
	fmt.Println(m)
}

func main() {
	fmt.Println("hello world")
	//initMap()
	initMap1()
}

// 声明变量
// var xxx [identifyTypes]
