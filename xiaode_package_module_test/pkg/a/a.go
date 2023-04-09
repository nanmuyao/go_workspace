package main

var  A1 = 1

func Add(a int) int {
	return a
}

func a(){
	println("func a run")
}

func init() {
	a()
	println("init run")
}