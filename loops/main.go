package main

import "fmt"

func main() {
	fmt.Println("loops in go lang")

	days := []string{"sunday", "monday", "tuseday", "thusday", "friday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for _, days := range days {
		fmt.Println("index is and value", days)
	}

	index := 1
	// for index < 10 {
	// 	fmt.Println(" me")
	// 	index++
	// }

	for index < 10 {
		if index == 2 {
			goto me
		}
		index++
	}
me:
	fmt.Println("hello")
}
