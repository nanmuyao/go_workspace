package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(reflect.TypeOf(b))
	os.Stdout.Write(b)
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s", b))
	fmt.Println(string(b))
	fmt.Println(reflect.TypeOf(string(b)))

}
