package main

import "fmt"

func main() {

	fmt.Println("if else")
	var result string
	loginCount := 11

	if loginCount < 10 {
		result = "regular user"
	} else {
		result = "premium user"
	}
	fmt.Println(result)

	if 100%2 == 0 {
		fmt.Println("number is even")
	}

}
