package main

import "fmt"

func main() {
	course := []string{"reactjs", "javascript", "swift", "python"}

	index := 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
