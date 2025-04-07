package lexer

const (
	// EOF
	EOF = iota

	// Whitespaces
	Semicolon
	Tab

	// Comparative Operators
	DoubleEquals
	NotEquals
	GreaterThan
	SmallerThan
	GreaterThanOrEquals
	SmallerThanOrEquals

	// Additive Operators/Signs
	Plus
	Minus

	// Multiplicative Operators
	Multiply
	Divide
	Modulus

	// Parentheses/Brackets
	LeftParentheses
	RightParentheses
	LeftBracket
	RightBracket
	LeftBrace
	RightBrace

	// Other signs
	Equals
	Pipe
	Exclamation
	Comma

	// Basic data types
	Int
	Double

	// Booleans
	True
	False

	// Identifier
	Identifier

	// Keywords
	Fn
	Do
	Return
	If
	Unless
	Else
)