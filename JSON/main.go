package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"loveName"`
	Price    int    `json:"tags,omitempty"`
	Password string `json:"-"`
}

func main() {
	fmt.Println("this is  json video")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"rahul", 5000, "123456"},
		{"rahul", 6000, "0977"},
		{"rahul", 9000, "1258667456"},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonData := []byte(
		`
		{
			"courseName":"GO LANG",
			"Price":"00",
			"Resourse":"youtube"
		}
		`,
	)
	var Course course

	res := json.Valid(jsonData)

	if res {
		fmt.Println("json")

	}
	err := json.Unmarshal(jsonData, &Course)
	if err != nil {
		panic(err)
	}
	fmt.Println(Course)
}
