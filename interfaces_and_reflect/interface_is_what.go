//package main
//
//import "fmt"
//
//type stockPosition struct {
//	ticker     string
//	sharePrice float32
//	count      float32
//}
//
///* method to determine the value of a stock position */
//func (s stockPosition) getValue() float32 {
//	return s.sharePrice * s.count
//}
//
//type car struct {
//	make  string
//	model string
//	price float32
//}
//
///* method to determine the value of a car */
//func (c car) getValue() float32 {
//	return c.price
//}
//
///* contract that defines different things that have value */
//type valuable interface {
//	getValue() float32
//}
//
//func showValue(asset valuable) {
//	fmt.Printf("Value of the asset is %f\n", asset.getValue())
//}
//
//func main() {
//	var o valuable = stockPosition{"GOOG", 577.20, 4}
//	showValue(o)
//	o = car{"BMW", "M3", 66500}
//	showValue(o)
//}

package main

import "fmt"

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	i int
}

func (s *Simple) Get() int {
	return s.i
}

func (s *Simple) Set(i int) {
	s.i = i
}

func Foo(simple Simpler) {
	i := simple.Get()
	fmt.Println("111", i)
	simple.Set(1)
	i = simple.Get()
	fmt.Println("222", i)
}

func main() {
	//var s Simpler = Simple{}
	//var interfaceS Simpler
	//interfaceS = &s
	//Foo(interfaceS)

	//报错 cannot use Simple{…} (value of type Simple) as Simpler value in variable
	//declaration: Simple does not implement Simpler (method Get has pointer receiver)
	//var s Simpler = Simple{i: 10}

	// 这里需要取地址，因为方法实现使用的是指针
	var s Simpler = &Simple{i: 10}
	Foo(s)
}
