package main

import (
	"fmt"
	"io/ioutil"
)

type token struct {
	dataType     string
	valeurString string
	valeurInt    int
	nbLigne      int
}

func main() {
	operators := [6]string{"+", "-", "*", "/", "%", "^"}
	ident := [4]string{"if", "for", "while", "else"}

	data, err := ioutil.ReadFile("test.go")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	for charPos := 0; charPos < len(data); charPos++ {

		letter := string(data[charPos])
		fmt.Println(letter)

		for _, op := range operators {
			if isFirstLetter(letter, op) {
				fmt.Println("operator")
			}

		}

		for _, id := range ident {
			if isFirstLetter(letter, id) {
				fmt.Println("ident")
			}
		}
	}

}

func isFirstLetter(char string, word string) bool {
	if char == string(word[0]) {
		return true
	} else {
		return false
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
