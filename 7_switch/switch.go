package main

import (
	"fmt"
	"go/types"
)

func main() {

	// switch 5 {
	// case 1, 2, 3, 4:
	// 	print("Smallers")
	// case 5, 6, 7, 8, 9:
	// 	print("Biggers")
	// case 10:
	// 	print("Biggest: TEN")
	// default:
	// 	print("None")
	// }
	// print("\n")

	// switch time.Now().Weekday() {
	// case time.Thursday, time.Friday:
	// 	print("Endinggg")
	// default:
	// 	print("NONONONONONO")
	// }

	// whoAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		println("Integer")
	// 	case string:
	// 		println("String")
	// 	default:
	// 		fmt.Println("Donno ", t)
	// 	}
	// }
	// whoAmI("golang")
	// whoAmI(100)
	// whoAmI(43.5)
	// whoAmI(true)

	// x := 10
	// y := 20
	fmt.Print(types.Bool)
}
