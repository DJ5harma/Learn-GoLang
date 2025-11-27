package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("")

	// var m = make(map[string]int)
	// var m = map[string]int{"price": 90}

	// m["name"] = 100
	// m["area"] = 2

	// delete(m, "area")

	// fmt.Println(m["name"])
	// fmt.Println(m["wrong"])
	// fmt.Println("length:", len(m))
	// fmt.Println(m)

	// m := map[int]string{1: "Hello"}

	// temp, ok := m[1]
	// if ok {
	// 	fmt.Println("Found", ok, temp)
	// } else {
	// 	fmt.Println("Not Found", ok, temp)
	// }

	m1 := map[string]int{"Hello": 1}
	m2 := map[string]int{"Hello": 1}

	fmt.Println(maps.Equal(m1, m2))

}
