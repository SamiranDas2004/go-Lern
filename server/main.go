package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb video")
	//GetRequest()
	//PostRequest()
	postFormReq()
}

func GetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)
	fmt.Println(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseStringe strings.Builder

	res, _ := ioutil.ReadAll(response.Body)
	buteCount, _ := responseStringe.Write(res)
	fmt.Println(responseStringe.String())

	fmt.Println("in the body", string(buteCount))
	defer response.Body.Close()
}

func PostRequest() {
	const myUrl = "http://localhost:8000/post"

	reqBody := strings.NewReader(`
	{
		"coursename":"lern go",
		"purpose":"be a better engineer"
	}
	`)
	res, err := http.Post(myUrl, "application/json", reqBody)

	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
	defer res.Body.Close()
}

func postFormReq() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "esha")
	data.Add("lastName", "Das")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	defer response.Body.Close()
}
