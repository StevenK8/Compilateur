package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	operators := [6]string{"+", "-", "*", "/", "%", "^"}
	ident := [4]string{"if", "for", "while", "else"}

	data, err := ioutil.ReadFile("test.go")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	for charPos := 0; charPos < len(data); charPos++ {
		fmt.Println(string(data[charPos]))
		for _, op := range operators {
			if op == string(data[charPos]) {
				fmt.Println("operator")
			}
		}

		for _, op := range ident {
			if op == string(data[charPos]) {
				fmt.Println("ident")
			}
		}
	}

}
