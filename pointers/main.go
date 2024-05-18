package main

import "fmt"

func main() {
	maypointer := 33
	var ptr = &maypointer
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
