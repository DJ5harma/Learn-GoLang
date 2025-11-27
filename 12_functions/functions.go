package main

func add(a int, b int) (int, int) {
	return a + b, 2
}

func sub(a, b int) (int, int) {
	return a - b, 2
}

func languages() (string, string, string) {
	return "javascript", "python", "c++"
}

func processIt(power2 func(a int) int) {
	var x int = power2(10)
	println(x)
}

func main() {

	processIt(func(a int) int {

		a = a * a
		return a
	})
	// l1, l2, _ := languages()
	// println(l1, l2)

}
