package main

import "fmt"

func main() {
	rahul := User{"Rahul", "Sam@gmai.com", 20, "male"}
	fmt.Println(rahul)
	fmt.Printf(" %+v", rahul)
}

type User struct {
	Name  string
	Email string
	Age   int
	Sex   string
}
