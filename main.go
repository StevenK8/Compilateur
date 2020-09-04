package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type TokenType string

const (
	Operator            TokenType = "Operator"
	Equal_add           TokenType = "+="
	Equal_sub           TokenType = "-="
	Equal_mult          TokenType = "*="
	Equal_increment     TokenType = "++"
	Equal_pow           TokenType = "**"
	Equal_Equal         TokenType = "=="
	Equal               TokenType = "="
	Operator_plus       TokenType = "+"
	Operator_minus      TokenType = "-"
	Operator_mult       TokenType = "*"
	Parenthese_ouvrante TokenType = "("
	Parenthese_fermante TokenType = ")"
	Point_virgule       TokenType = ";"
	Constant            TokenType = "Number"
	Word                TokenType = "Word"
)

type token struct {
	dataType     TokenType
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
		println(strconv.Itoa(a.nbLigne) + " \t[" + string(a.valeurString) + "]")
	}

}

func lexer(data []byte) []token {

	// var isLetter bool
	// var isOperator bool
	// var isNumber bool
	// var currentWord string
	var tokenTab []token
	var numOfLine int = 0

	for charPos := 0; charPos < len(data); charPos++ {

		currentChar := string(data[charPos])

		changeLine, err := regexp.MatchString(`\n`, currentChar)
		if err != nil {
			println(err)
		}
		if changeLine {
			numOfLine++
		}

		//Creation ou continuté du mot

		//Verifier la fin du mot avant de tester la continuité
		// tokenTab, isLetter, currentWord = checkMatch(`[a-zA-Z]`, isLetter, currentChar, currentWord, numOfLine, tokenTab)
		// tokenTab, isNumber, currentWord = checkMatch(`[0-9]`, isNumber, currentChar, currentWord, numOfLine, tokenTab)
		// tokenTab, isOperator, currentWord = checkMatch(`[(+*;=)]`, isOperator, currentChar, currentWord, numOfLine, tokenTab)

		if checkMatchChar("[(+*;=)]", currentChar) {
			dataType, longueur := get_operator(data, charPos)
			tokenTab = append(tokenTab, token{dataType, string(data[charPos : charPos+longueur+1]), 0, charPos})
			charPos += longueur
		} else if checkMatchChar("[a-zA-Z]", currentChar) {
			dataType, longueur := get_ident(data, charPos)
			tokenTab = append(tokenTab, token{dataType, string(data[charPos : charPos+longueur+1]), 0, charPos})
			charPos += longueur
		} else if checkMatchChar("[0-9]", currentChar) {
			dataType, longueur := get_number(data, charPos)
			i, err := strconv.Atoi(string(data[charPos : charPos+longueur+1]))
			if err != nil {
				println(err)
			}
			tokenTab = append(tokenTab, token{dataType, string(data[charPos : charPos+longueur+1]), i, charPos})
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

func checkMatch(regex string, isType bool, currentChar string, currentWord string, charPos int, tokenTab []token) ([]token, bool, string) {
	matched, err := regexp.MatchString(regex, currentChar)
	if err != nil {
		println(err)
	}
	//matchedCurrentWord, err := regexp.MatchString("^([a-zA-Z]|[^ ])*$", currentWord)
	// println("Current word : ", currentWord, " - ", matchedCurrentWord)
	if matched && (isType || len(currentWord) == 0) {
		currentWord += currentChar
		isType = true
	} else if isType && !matched {
		// tokenTab = append(tokenTab, token{"operateur", currentWord, 0, charPos})
		currentWord = ""
		isType = false
	}
	//  else if matched && !isType {
	// 	tokenTab = append(tokenTab, token{"", currentWord, 0, charPos})
	// 	currentWord = currentChar
	// 	isType = true
	// 	//Et l'autre qui etait en cours ?
	// }
	return tokenTab, isType, currentWord
}

func get_operator(data []byte, charPos int) (TokenType, int) {
	var dataType TokenType
	longueur := 0
	switch string(data[charPos]) {
	case "+":
		dataType = Operator_plus
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = Equal_add
		} else if string(data[charPos+1]) == "+" {
			longueur++
			dataType = Equal_increment
		}
		break
	case "-":
		dataType = Operator_minus
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = Equal_sub
		}
		break
	case "*":
		dataType = Operator_mult
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = Equal_mult
		} else if string(data[charPos+1]) == "*" {
			longueur++
			dataType = Equal_pow
		}
		break
	case "=":
		dataType = Equal
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = Equal_Equal
		}
		break
	case "(":
		dataType = Parenthese_ouvrante
		break
	case ")":
		dataType = Parenthese_fermante
		break
	case ";":
		dataType = Point_virgule
		break
	}
	return dataType, longueur
}

func get_ident(data []byte, charPos int) (TokenType, int) {
	dataType := Word
	var longueur int

	for longueur = 0; charPos < len(data); charPos++ {
		if !checkMatchChar(`[a-zA-Z]`, string(data[charPos+1])) {
			break
		}
		longueur++
	}

	return dataType, longueur
}

func get_number(data []byte, charPos int) (TokenType, int) {
	dataType := Constant
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
