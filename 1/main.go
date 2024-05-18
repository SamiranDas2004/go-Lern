package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice")

	frutelist := []string{"esha"}

	fmt.Printf("Type of data: %T\n", frutelist)

	frutelist = append(frutelist, "riti", "malobi")
	// frutelist = append(frutelist[:2])
	fmt.Println(frutelist)

	gf := make([]int, 0)

	gf = append(gf, 1, 1, 2, 23, 4)

	fmt.Println(gf)

	sort.Ints(gf)
	fmt.Println(gf)

	var alcohol = []string{"beer", "rum", "gin", "vodka", "whisky", "smhampain", "desi"}

	fmt.Println(alcohol)
	var index int = 2
	alcohol = append(alcohol[:index], alcohol[index+1:]...)
	fmt.Println(alcohol)
}
