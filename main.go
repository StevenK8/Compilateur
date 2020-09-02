package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	operators := [6]string{"+", "-", "*", "/", "%", "^"}
	ident := [4]string{"if", "for", "while", "else"}

	data, err := ioutil.ReadFile("test.go")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	tab := strings.Fields(string(data))
	fmt.Println("tab:", tab)
	// fmt.Println("Contents of file:", string(data))
}
