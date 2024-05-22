package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("SwitchCase in go")
	rand.Seed(time.Now().UnixNano())
	dise := rand.Intn(6) + 1
	fmt.Println(dise)

	switch dise {
	case 1:
		fmt.Println(" this is 1")
	case 5:
		fmt.Println("this is 5")
	default:
		fmt.Println("love")
	}

}
