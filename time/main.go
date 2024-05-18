package main

import (
	"fmt"
	"time"
)

func main() {
	println("welcome to golang time")

	presetime := time.Now()
	// fmt.Printf("time is ", presetime)

	fmt.Println(presetime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2090, time.August, 7, 23, 33, 45, 0, time.UTC)
	fmt.Println(createDate.Format("01-22-2006 15:04:05 Monday"))
}
