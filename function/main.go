package main

import "fmt"

func main() {
	hello()
	pro, message := aniother(2, 3, 4, 5, 6, 6, 7, 7, 8)
	fmt.Println(pro)
	fmt.Println("hello from the function")
	fmt.Println("hello from the function2", message)
}

func aniother(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "yo"
}

func hello() {
	fmt.Println("another hello")
}
