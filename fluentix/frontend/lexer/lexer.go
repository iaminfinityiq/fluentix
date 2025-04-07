package lexer

import (
	"fluentix/helpers"
	"fluentix/runtime"
	"fmt"
)

// Character checkers
func is_digit(r rune) bool {
	return helpers.RuneInString(r, "0123456789.")
}

func is_whitespace(r rune) bool {
	return helpers.RuneInString(r, " ")
}

func is_letter(r rune) bool {
	return helpers.RuneInString(r, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")
}

func is_legal(r rune) bool {
	return is_digit(r) || is_letter(r)
}

func Tokenize(code string, file_extension string) runtime.RuntimeResult {
	characters := helpers.StringToRune(code)
	tokens := []Token{}

	keywords := map[string]int{
		"if":       If,
		"unless":   Unless,
		"elseif":   Unless,
		"elif":     Unless,
		"else":     Else,
		"fn":       Fn,
		"function": Fn,
		"return":   Return,
		"do":       Do,
		"true":     True,
		"false":    False,
	}

	for len(characters) > 0 {
		switch characters[0] {
		case ';':
			tokens = append(tokens, Token{TokenType: Semicolon, Value: ";"})
		case '\n':
			tokens = append(tokens, Token{TokenType: Semicolon, Value: ";"})
		case '\t':
			if file_extension == "flu" {
				tokens = append(tokens, Token{TokenType: Tab, Value: "\\t"})
			}
		case '+':
			tokens = append(tokens, Token{TokenType: Plus, Value: "+"})
		case '-':
			characters = characters[1:]
			if len(characters) == 0 {
				tokens = append(tokens, Token{TokenType: Minus, Value: "-"})
				continue
			}

			if characters[0] == '>' {
				tokens = append(tokens, Token{TokenType: Do, Value: "->"})
				characters = characters[1:]
				continue
			}

			tokens = append(tokens, Token{TokenType: Minus, Value: "-"})
		case '*':
			tokens = append(tokens, Token{TokenType: Multiply, Value: "*"})
		case '%':
			tokens = append(tokens, Token{TokenType: Modulus, Value: "%"})
		case '/':
			tokens = append(tokens, Token{TokenType: Divide, Value: "/"})
		case '(':
			tokens = append(tokens, Token{TokenType: LeftParentheses, Value: "("})
		case ')':
			tokens = append(tokens, Token{TokenType: RightParentheses, Value: ")"})
		case '{':
			tokens = append(tokens, Token{TokenType: LeftBrace, Value: "{"})
		case '}':
			tokens = append(tokens, Token{TokenType: RightBrace, Value: "}"})
		case '|':
			tokens = append(tokens, Token{TokenType: Pipe, Value: "|"})
		case '!':
			characters = characters[1:]
			if len(characters) == 0 {
				tokens = append(tokens, Token{TokenType: Exclamation, Value: "!"})
				continue
			}

			if characters[0] == '=' {
				tokens = append(tokens, Token{TokenType: NotEquals, Value: "!="})
				characters = characters[1:]
				continue
			}

			tokens = append(tokens, Token{TokenType: Exclamation, Value: "!"})
		case '=':
			characters = characters[1:]
			if len(characters) == 0 {
				tokens = append(tokens, Token{TokenType: Equals, Value: "="})
				continue
			}

			if characters[0] == '=' {
				tokens = append(tokens, Token{TokenType: DoubleEquals, Value: "="})
				characters = characters[1:]
				continue
			}

			tokens = append(tokens, Token{TokenType: Equals, Value: "="})
		case '>':
			characters = characters[1:]
			if len(characters) == 0 {
				tokens = append(tokens, Token{TokenType: GreaterThan, Value: ">"})
				continue
			}

			if characters[0] == '=' {
				tokens = append(tokens, Token{TokenType: GreaterThanOrEquals, Value: ">="})
				characters = characters[1:]
				continue
			}

			tokens = append(tokens, Token{TokenType: GreaterThan, Value: ">"})
		case '<':
			characters = characters[1:]
			if len(characters) == 0 {
				tokens = append(tokens, Token{TokenType: SmallerThan, Value: "<"})
				continue
			}

			if characters[0] == '=' {
				tokens = append(tokens, Token{TokenType: SmallerThanOrEquals, Value: "<="})
				characters = characters[1:]
				continue
			}

			tokens = append(tokens, Token{TokenType: SmallerThan, Value: "<"})
		case ',':
			tokens = append(tokens, Token{TokenType: Comma, Value: ","})
		default:
			if characters[0] == ' ' && file_extension == "flu" {
				if len(characters) >= 4 {
					if characters[1] == ' ' && characters[2] == ' ' && characters[3] == ' ' {
						tokens = append(tokens, Token{TokenType: Tab, Value: "\t"})
						characters = characters[1:]
						characters = characters[1:]
						characters = characters[1:]
						characters = characters[1:]
						continue
					}
				}
			}

			if is_whitespace(characters[0]) {
				characters = characters[1:]
				continue
			}

			// Handle numbers
			if is_digit(characters[0]) {
				number := ""
				dotCount := uint8(0)

				for len(characters) > 0 && is_digit(characters[0]) {
					if characters[0] == '.' {
						dotCount++
					}
					number += string(characters[0])
					characters = characters[1:]
				}

				switch dotCount {
				case 0:
					tokens = append(tokens, Token{TokenType: Int, Value: number})
				case 1:
					tokens = append(tokens, Token{TokenType: Double, Value: number})
				default:
					return runtime.Failure(runtime.SyntaxError{
						Reason: "Too many '.' in number: " + number,
					})
				}

				continue
			}

			if is_letter(characters[0]) {
				var identifier string = ""

				for len(characters) > 0 && is_legal(characters[0]) {
					identifier += string(characters[0])
					characters = characters[1:]
				}

				if tokenType, found := keywords[identifier]; found {
					tokens = append(tokens, Token{TokenType: tokenType, Value: identifier})
				} else {
					tokens = append(tokens, Token{TokenType: Identifier, Value: identifier})
				}

				continue
			}

			return runtime.Failure(runtime.SyntaxError{
				Reason: fmt.Sprintf("Unexpected character: '%s'", string(characters[0])),
			})
		}

		characters = characters[1:]
	}

	tokens = append(tokens, Token{TokenType: EOF, Value: "EOF"})
	return runtime.Success(tokens)
}
