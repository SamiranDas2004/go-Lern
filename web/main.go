package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:3000/"

func main() {
	fmt.Println(" web req")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response type %T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)
}
