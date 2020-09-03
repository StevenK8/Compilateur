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

		tokenTab, isLetter, currentWord = checkMatch(`[a-zA-Z]`, isLetter, currentChar, currentWord, charPos, tokenTab)
		tokenTab, isNumber, currentWord = checkMatch(`[0-9]`, isNumber, currentChar, currentWord, charPos, tokenTab)
		tokenTab, isOperator, currentWord = checkMatch(`[(+*;=)]`, isOperator, currentChar, currentWord, charPos, tokenTab)

	}
	return tokenTab
}

func checkMatch(regex string, isType bool, currentChar string, currentWord string, charPos int, tokenTab []token) ([]token, bool, string) {
	matched, err := regexp.MatchString(regex, currentChar)
	if err != nil {
		println(err)
	}
	if matched && (isType || len(currentWord) == 0) {
		currentWord += currentChar
		isType = true
	} else if isType && !matched {
		tokenTab = append(tokenTab, token{"operateur", currentWord, 0, charPos})
		currentWord = ""
		isType = false
	}
	return tokenTab, isType, currentWord
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
