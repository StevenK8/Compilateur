package token

type TokenType string

const (
	Operator           = "Operator"
	EqualAdd           = "+="
	EqualSub           = "-="
	EqualMult          = "*="
	Increment          = "++"
	Pow                = "**"
	Equalequal         = "=="
	Equal              = "Equal"
	OperatorPlus       = "Add"
	OperatorMinus      = "Sub"
	OperatorMult       = "Mult"
	ParentheseOuvrante = "Open_Paren"
	ParentheseFermante = "Close_Paren"
	LeftBrace          = "Open_Brace"
	RightBrace         = "Close_Brace"
	PointVirgule       = "Semicolon"
	Constant           = "Number"
	Word               = "Word"
	KeywordIf          = "If"
	KeywordWhile       = "While"
)

type Token struct {
	DataType     TokenType
	ValeurString string
	ValeurInt    int
	NbLigne      int
}
