package main

import (
	"fmt"
	"net/url"
)

const Url = "http://localhost:3000/verify/samiran?param1=value1&param2=value2"

func main() {
	fmt.Println("Welcome to URL parsing")
	fmt.Println(Url)

	result, _ := url.Parse(Url)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)

	quaryP := result.Query()
	fmt.Printf("type of query params%T\n", quaryP)

	fmt.Println(quaryP["param2"])

	partsOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "localhost:3000",
		Path:    "/verify/samiran",
		RawPath: "param1=value1&param2=value2",
	}

	newUrl := partsOfUrl.String()
	fmt.Println(newUrl)
}
