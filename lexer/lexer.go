package lexer

import (
	"log"
	"regexp"
	"strconv"

	token "github.com/StevenK8/Compilateur/token"
)

func Lexer(data []byte) []token.Token {

	var tokenTab []token.Token
	numOfLine := 1

	for charPos := 0; charPos < len(data); charPos++ {

		if checkMatchChar(`\n`, string(data[charPos])) {
			numOfLine++
		}

		dataType, longueur := getOperator(data, charPos)

		if dataType != token.Ignore {
			i, _ := strconv.Atoi(string(data[charPos : charPos+longueur+1]))
			tokenTab = append(tokenTab, token.Token{dataType, string(data[charPos : charPos+longueur+1]), i, numOfLine})
			charPos += longueur
		}

	}
	return tokenTab
}

func checkMatchChar(regex string, char string) bool {
	r, err := regexp.Compile(regex)
	if err != nil {
		log.Fatal(err)
	}

	matched := r.MatchString(char)

	return matched
}

func getOperator(data []byte, charPos int) (token.TokenType, int) {
	var dataType token.TokenType
	longueur := 0
	switch string(data[charPos]) {
	case "+":
		dataType = token.OperatorPlus
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualAdd
			} else if string(data[charPos+1]) == "+" {
				longueur++
				dataType = token.Increment
			}
		}

	case "-":
		dataType = token.OperatorMinus
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualSub
			}
		}

	case "*":
		dataType = token.OperatorMult
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualMult
			} else if string(data[charPos+1]) == "*" {
				longueur++
				dataType = token.Pow
			}
		}

	case "/":
		dataType = token.OperatorDiv
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualDiv
			}
		}

	case "%":
		dataType = token.OperatorMod
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.EqualMod
			}
		}

	case "=":
		dataType = token.Equal
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.Equalequal
			}
		}

	case "!":
		dataType = token.Not
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "=" {
				longueur++
				dataType = token.NotEqual
			}
		}

	case "&":
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "&" {
				longueur++
				dataType = token.And
			}
		}

	case "|":
		if charPos < len(data)-1 {
			if string(data[charPos+1]) == "|" {
				longueur++
				dataType = token.Or
			}
		}

	case "(":
		dataType = token.ParentheseOuvrante

	case ")":
		dataType = token.ParentheseFermante

	case "{":
		dataType = token.LeftBrace

	case "}":
		dataType = token.RightBrace

	case ";":
		dataType = token.PointVirgule

	case ",":
		dataType = token.Virgule

	case "<":
		dataType = token.LessThan
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = token.LessOrEqual
		}

	case ">":
		dataType = token.GreaterThan
		if string(data[charPos+1]) == "=" {
			longueur++
			dataType = token.GreaterOrEqual
		}

	default:
		return getIdent(data, charPos)

	}
	return dataType, longueur
}

func getIdent(data []byte, charPos int) (token.TokenType, int) {
	var dataType token.TokenType
	dataType = token.Ident
	var longueur int

	if !checkMatchChar(`[a-zA-Z]`, string(data[charPos])) {
		return getNumber(data, charPos)
	}

	for longueur = 0; charPos < len(data)-1; charPos++ {
		if condGetWord(data, charPos, "if") {
			dataType = token.KeywordIf
		} else if condGetWord(data, charPos, "while") {
			dataType = token.KeywordWhile
		} else if condGetWord(data, charPos, "true") {
			dataType = token.BooleanTrue
		} else if condGetWord(data, charPos, "false") {
			dataType = token.BooleanFalse
		} else if condGetWord(data, charPos, "else") {
			dataType = token.KeywordElse
		} else if condGetWord(data, charPos, "debug") {
			dataType = token.Debug
		} else if condGetWord(data, charPos, "int") {
			dataType = token.KeywordInt
		} else if condGetWord(data, charPos, "for") {
			dataType = token.KeywordFor
		} else if condGetWord(data, charPos, "break") {
			dataType = token.KeywordBreak
		} else if condGetWord(data, charPos, "continue") {
			dataType = token.KeywordContinue
		} else if condGetWord(data, charPos, "return") {
			dataType = token.Return
		} else if condGetWord(data, charPos, "send") {
			dataType = token.KeywordSend
		}

		if !checkMatchChar(`[a-zA-Z]`, string(data[charPos+1])) {
			break
		}

		longueur++
	}

	return dataType, longueur
}

func getNumber(data []byte, charPos int) (token.TokenType, int) {
	var dataType token.TokenType = token.Constant
	var longueur int

	if !checkMatchChar(`[0-9]`, string(data[charPos])) {
		return token.Ignore, 0
	}

	for longueur = 0; charPos < len(data)-1; charPos++ {
		if !checkMatchChar(`[0-9]`, string(data[charPos+1])) {
			break
		}
		longueur++
	}

	return dataType, longueur
}

func condGetWord(data []byte, charPos int, mot string) bool {
	lenMot := len(mot)
	if charPos < len(data)-(lenMot-1) && string(data[charPos:charPos+lenMot]) == mot && !checkMatchChar(`[a-zA-Z]`, string(data[charPos+lenMot])) && (charPos == 0 || charPos >= 1 && !checkMatchChar(`[a-zA-Z]`, string(data[charPos-1]))) {
		return true
	}
	return false
}
