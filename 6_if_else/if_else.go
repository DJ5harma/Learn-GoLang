package main

func main() {
	age := 11

	if age > 18 {
		print("Is Adult")
	} else if age >= 12 {
		print("Is Teen")
	} else {
		print("Is Child")
	}
	print("\n")

	var role = "admin"

	var hasPermissions = false
	if role == "admin" && hasPermissions {
		print("Allowed")
	} else {
		print("Denied")
	}
	print("\n")

	if AGE := 15; AGE > 18 {
		print("AGE ", "Adult")
	} else {
		print("AGE ", "NOT Adult")
	}
	print("\n")

}
