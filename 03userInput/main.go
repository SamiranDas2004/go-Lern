package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome my love"
	println(welcome)

	reader := bufio.NewReader(os.Stdin)
	println("enter the ratings for my skills")

	bal, _ := reader.ReadString('\n')
	fmt.Printf("ratings %T\n", bal)
}
