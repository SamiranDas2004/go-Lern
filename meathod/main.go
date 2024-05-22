package main

import "fmt"

func main() {
	rahul := User{"Rahul", "Sam@gmai.com", 20, "male"}

	fmt.Printf(" %+v", rahul)
	rahul.Emails()
	fmt.Printf(" %+v", rahul)
}

type User struct {
	Name  string
	Email string
	Age   int
	Sex   string
}

func (u User) Emails() {
	fmt.Println("is user active: ", u.Email)
	u.Email = "fuck"
	println(u.Email)
}
