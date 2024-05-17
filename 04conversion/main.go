package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	println("welcome to our pizza app")
	println("please rate our pizza between 1 and 5")

	read := bufio.NewReader(os.Stdin)

	input, _ := read.ReadString('\n')
	fmt.Printf(" %T", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1", numrating+1)
	}
	// fmt.Printf("%T", numrating)
}
