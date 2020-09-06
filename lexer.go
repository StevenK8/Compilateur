package main

import (
	"regexp"
	"strconv"
)

type tokenType string

const (
	operator           tokenType = "Operator"
	equalAdd           tokenType = "+="
	equalSub           tokenType = "-="
	equalMult          tokenType = "*="
	equalIncrement     tokenType = "++"
	equalPow           tokenType = "**"
	equalequal         tokenType = "=="
	equal              tokenType = "Equal"
	operatorPlus       tokenType = "Add"
	operatorMinus      tokenType = "Sub"
	operatorMult       tokenType = "Mult"
	parentheseOuvrante tokenType = "Open_Paren"
	parentheseFermante tokenType = "Close_Paren"
	pointVirgule       tokenType = "Semicolon"
	constant           tokenType = "Number"
	word               tokenType = "Word"
	keywordIf          tokenType = "If"
	keywordWhile       tokenType = "While"
)

type token struct {
	dataType     tokenType
	valeurString string
	valeurInt    int
	nbLigne      int
}

func lexer(data []byte) []token {

	var tokenTab []token
	numOfLine := 1

	for charPos := 0; charPos < len(data); charPos++ {

		currentChar := string(data[charPos])

		changeLine, err := regexp.MatchString(`\n`, currentChar)
		if err != nil {
			println(err)
		}
		if changeLine {
			numOfLine++
		}

		if checkMatchChar("[(+*;=)]", currentChar) {
			dataType, longueur := getOperator(data, charPos)
			tokenTab = append(tokenTab, token{dataType, string(data[charPos : charPos+longueur+1]), 0, numOfLine})
			charPos += longueur
		} else if checkMatchChar("[a-zA-Z]", currentChar) {
			dataType, longueur := getIdent(data, charPos)
			tokenTab = append(tokenTab, token{dataType, string(data[charPos : charPos+longueur+1]), 0, numOfLine})
			charPos += longueur
		} else if checkMatchChar("[0-9]", currentChar) {
			dataType, longueur := getNumber(data, charPos)
			i, err := strconv.Atoi(string(data[charPos : charPos+longueur+1]))
			if err != nil {
				println(err)
			}
			tokenTab = append(tokenTab, token{dataType, string(data[charPos : charPos+longueur+1]), i, numOfLine})
			charPos += longueur
		}

	}
	return tokenTab
}

func checkMatchChar(regex string, char string) bool {
	matched, err := regexp.MatchString(regex, char)
	if err != nil {
		println(err)
	}

	return matched
}

func getOperator(data []byte, charPos int) (tokenType, int) {
	var dataType tokenType
	longueur := 0
	switch string(data[charPos]) {
	case "+":
		dataType = operatorPlus
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = equalAdd
		} else if string(data[charPos+1]) == "+" {
			longueur++
			dataType = equalIncrement
		}
		break
	case "-":
		dataType = operatorMinus
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = equalSub
		}
		break
	case "*":
		dataType = operatorMult
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = equalMult
		} else if string(data[charPos+1]) == "*" {
			longueur++
			dataType = equalPow
		}
		break
	case "=":
		dataType = equal
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = equalequal
		}
		break
	case "(":
		dataType = parentheseOuvrante
		break
	case ")":
		dataType = parentheseFermante
		break
	case ";":
		dataType = pointVirgule
		break
	}
	return dataType, longueur
}

func getIdent(data []byte, charPos int) (tokenType, int) {
	dataType := word
	var longueur int

	for longueur = 0; charPos < len(data); charPos++ {
		if string(data[charPos:charPos+2]) == "if" && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+2])) {
			dataType = keywordIf
		} else if string(data[charPos:charPos+5]) == "while" && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+5])) {
			dataType = keywordWhile
		}

		if !checkMatchChar(`[a-zA-Z]`, string(data[charPos+1])) {
			break
		}

		longueur++
	}

	return dataType, longueur
}

func getNumber(data []byte, charPos int) (tokenType, int) {
	dataType := constant
	var longueur int

	for longueur = 0; charPos < len(data); charPos++ {
		if !checkMatchChar(`[0-9]`, string(data[charPos+1])) {
			break
		}
		longueur++
	}

	return dataType, longueur
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
