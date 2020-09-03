package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type LeaveType string

const (
	AnnualLeave LeaveType = "AnnualLeave"
	Sick                  = "Sick"
	BankHoliday           = "BankHoliday"
	Other                 = "Other"
)

type token struct {
	dataType     string
	valeurString string
	valeurInt    int
	nbLigne      int
}

func main() {

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	tokenTab := lexer(data)
	for _, a := range tokenTab {
		println(string(a.valeurString))
	}

}

func lexer(data []byte) []token {

	var isLetter bool
	var isOperator bool
	var isNumber bool
	var currentWord string
	var tokenTab []token

	for charPos := 0; charPos < len(data); charPos++ {

		currentChar := string(data[charPos])

		matched, err := regexp.MatchString(`[a-zA-Z]`, currentChar)
		if err != nil {
			println(err)
		}

		if matched && (isLetter || len(currentWord) == 0) {
			currentWord += currentChar
			isLetter = true
		} else if isLetter && !matched {
			tokenTab = append(tokenTab, token{"lettre", currentWord, 0, charPos})
			currentWord = ""
			isLetter = false
		}

		matched, err = regexp.MatchString(`[0-9]`, currentChar)
		if err != nil {
			println(err)
		}

		if matched && (isNumber || len(currentWord) == 0) {
			currentWord += currentChar
			isLetter = false
			isNumber = true
			isOperator = false
		} else if isNumber && !matched {
			tokenTab = append(tokenTab, token{"nombre", currentWord, 0, charPos})
			currentWord = ""
			isNumber = false
		}

		matched, err = regexp.MatchString(`[(+*;=]`, currentChar)
		if err != nil {
			println(err)
		}

		if matched && (isOperator || len(currentWord) == 0) {
			currentWord += currentChar
			isLetter = false
			isNumber = false
			isOperator = true
		} else if isOperator && !matched {
			tokenTab = append(tokenTab, token{"ope", currentWord, 0, charPos})
			currentWord = ""
			isOperator = false
		}

		// fmt.Println(currentChar)
	}
	return tokenTab
}

func isFirstcurrentChar(char string, word string) bool {
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

func isOperator(op string) bool {
	operators := []string{"+", "-", "*", "/", "%", "^"}
	if contains(operators, op) {
		return true
	}
	return false
}

func isIdent(id string) bool {
	ident := []string{"if", "for", "while", "else"}
	if contains(ident, id) {
		return true
	}
	return false
}
