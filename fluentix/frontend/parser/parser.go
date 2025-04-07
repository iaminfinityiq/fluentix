package parser

import (
	"fluentix/frontend/lexer"
	"fluentix/helpers"
	"fluentix/runtime"
	"fmt"
	"strconv"
)

func at(tokens *[]lexer.Token, file_extension string) lexer.Token {
	return (*tokens)[0]
}

func eat(tokens *[]lexer.Token, file_extension string) lexer.Token {
	var returned lexer.Token = at(tokens, file_extension)
	skip_tabs(tokens, file_extension)
	*tokens = (*tokens)[1:]

	return returned
}

func skip_tabs(tokens *[]lexer.Token, file_extension string) {
	for (*tokens)[0].TokenType == lexer.Tab {
		*tokens = (*tokens)[1:]
	}
}

func expect(tokens *[]lexer.Token, file_extension string, token_type int) runtime.RuntimeResult {
	var returned lexer.Token = eat(tokens, file_extension)
	if returned.TokenType != token_type {
		return runtime.Failure(runtime.SyntaxError{Reason: fmt.Sprintf("Expected %d, got '%s'", token_type, returned.Value)})
	}

	return runtime.Success(returned)
}

func is_eof(tokens *[]lexer.Token, file_extension string) bool {
	return at(tokens, file_extension).TokenType == lexer.EOF
}

func is_semicolon(tokens *[]lexer.Token, file_extension string) bool {
	return at(tokens, file_extension).TokenType == lexer.Semicolon
}

func is_semicolon_or_eof(tokens *[]lexer.Token, file_extension string) bool {
	return is_semicolon(tokens, file_extension) || is_eof(tokens, file_extension)
}

func not_semicolon_or_eof(tokens *[]lexer.Token, file_extension string) bool {
	return !is_semicolon_or_eof(tokens, file_extension)
}

func ProduceBlock(tokens *[]lexer.Token, file_extension string, expect_right_brace bool) runtime.RuntimeResult {
	var block Block = Block{Body: []Statement{}}

	for len(*tokens) > 0 {
		for at(tokens, file_extension).TokenType == lexer.Semicolon {
			eat(tokens, file_extension)
			if len(*tokens) == 1 {
				break
			}
		}

		if len(*tokens) == 1 {
			break
		}

		if at(tokens, file_extension).TokenType == lexer.RightBrace && expect_right_brace {
			break
		}

		var rt runtime.RuntimeResult = parse_statement(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		block.Body = append(block.Body, rt.Result.(Statement))
	}

	if expect_right_brace {
		var rt runtime.RuntimeResult = expect(tokens, file_extension, lexer.RightBrace)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}
	}

	return runtime.Success(block)
}

func parse_indented_block(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult
	if file_extension == "flu" {
		var block []lexer.Token = []lexer.Token{}
		for {
			for is_semicolon(tokens, file_extension) {
				eat(tokens, file_extension)
			}

			if (*tokens)[0].TokenType != lexer.Tab {
				break
			}

			*tokens = (*tokens)[1:]
			for not_semicolon_or_eof(tokens, file_extension) {
				block = append(block, (*tokens)[0])
				*tokens = (*tokens)[1:]
			}

			if is_eof(tokens, file_extension) {
				break
			}

			block = append(block, eat(tokens, file_extension))
		}

		block = append(block, lexer.Token{TokenType: lexer.EOF, Value: "EOF"})
		rt = ProduceBlock(&block, file_extension, false)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	} else {
		for is_semicolon(tokens, file_extension) {
			eat(tokens, file_extension)
		}

		rt = expect(tokens, file_extension, lexer.LeftBrace)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		rt = ProduceBlock(tokens, file_extension, true)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		for is_semicolon(tokens, file_extension) {
			eat(tokens, file_extension)
		}

		return runtime.Success(rt.Result)
	}
}

func parse_statement(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult
	var flag bool = true
	switch at(tokens, file_extension).TokenType {
	case lexer.If:
		rt = parse_if_statement(tokens, file_extension)
		flag = false
	case lexer.Fn:
		rt = parse_function_declaration(tokens, file_extension)
		flag = false
	default:
		rt = parse_expression(tokens, file_extension)
	}

	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var returned Statement = rt.Result.(Statement)
	if not_semicolon_or_eof(tokens, file_extension) && flag {
		return runtime.Failure(runtime.SyntaxError{Reason: fmt.Sprintf("Expected %d or %d, got '%s'", lexer.Semicolon, lexer.EOF, at(tokens, file_extension).Value)})
	}

	return runtime.Success(returned)
}

func parse_if_statement(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	eat(tokens, file_extension)
	var rt runtime.RuntimeResult = parse_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var condition Expression = rt.Result.(Expression)

	rt = expect(tokens, file_extension, lexer.Do)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	rt = parse_indented_block(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var block Block = rt.Result.(Block)
	if at(tokens, file_extension).TokenType == lexer.Unless {
		rt = parse_if_statement(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		__, _ := rt.Result.(IfStatement)
		return runtime.Success(IfStatement{
			Condition: condition,
			Body: block,
			Next: &__,
		})
	}

	if at(tokens, file_extension).TokenType == lexer.Else {
		eat(tokens, file_extension)
		if at(tokens, file_extension).TokenType == lexer.If {
			rt = parse_if_statement(tokens, file_extension)
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			__, _ := rt.Result.(IfStatement)
			return runtime.Success(IfStatement{
				Condition: condition,
				Body: block,
				Next: &__,
			})
		}

		rt = expect(tokens, file_extension, lexer.Do)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		rt = parse_indented_block(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		var else_block Block = rt.Result.(Block)
		return runtime.Success(IfStatement{
			Condition: condition,
			Body: block,
			Next: &IfStatement{
				Condition: Boolean{Value: true},
				Body: else_block,
				Next: nil,
			},
		})
	}

	return runtime.Success(IfStatement{
		Condition: condition,
		Body: block,
		Next: &IfStatement{
			Condition: condition,
			Body: block,
			Next: nil,
		},
	})
}

func parse_function_declaration(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	eat(tokens, file_extension)
	var rt runtime.RuntimeResult = expect(tokens, file_extension, lexer.Identifier)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var function_name string = rt.Result.(lexer.Token).Value

	rt = expect(tokens, file_extension, lexer.LeftParentheses)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var arguments []string = []string{}
	for at(tokens, file_extension).TokenType != lexer.RightParentheses {
		rt = expect(tokens, file_extension, lexer.Identifier)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		arguments = append(arguments, rt.Result.(lexer.Token).Value)
		if at(tokens, file_extension).TokenType == lexer.RightParentheses {
			break
		}

		rt = expect(tokens, file_extension, lexer.Comma)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}
	}

	eat(tokens, file_extension)
	expect(tokens, file_extension, lexer.Do)

	rt = parse_indented_block(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var body Block = rt.Result.(Block)
	return runtime.Success(FunctionDeclaration{
		FunctionName: function_name,
		Arguments: arguments,
		Body: body,
	})
}

func parse_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = parse_assignment_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(rt.Result)
}

func parse_assignment_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = parse_comparative_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var left Expression = rt.Result.(Expression)
	if at(tokens, file_extension).TokenType != lexer.Equals {
		return runtime.Success(left)
	}

	eat(tokens, file_extension)
	rt = parse_assignment_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(AssignmentExpression{
		Assignee: left,
		Value: rt.Result.(Expression),
	})
}

func parse_comparative_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = parse_additive_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var left Expression = rt.Result.(Expression)
	var operators helpers.IntSet = helpers.IntSet{Set: make(map[int]bool)}
	operators.Add(lexer.DoubleEquals)
	operators.Add(lexer.NotEquals)
	operators.Add(lexer.GreaterThan)
	operators.Add(lexer.SmallerThan)
	operators.Add(lexer.GreaterThanOrEquals)
	operators.Add(lexer.SmallerThanOrEquals)

	for operators.Contains(at(tokens, file_extension).TokenType) {
		var operator = eat(tokens, file_extension).TokenType
		rt = parse_additive_expression(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		left = BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    rt.Result.(Expression),
		}
	}

	return runtime.Success(left)
}

func parse_additive_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = parse_multiplicative_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var left Expression = rt.Result.(Expression)
	var operators helpers.IntSet = helpers.IntSet{Set: make(map[int]bool)}
	operators.Add(lexer.Plus)
	operators.Add(lexer.Minus)

	for operators.Contains(at(tokens, file_extension).TokenType) {
		var operator = eat(tokens, file_extension).TokenType
		rt = parse_multiplicative_expression(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		left = BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    rt.Result.(Expression),
		}
	}

	return runtime.Success(left)
}

func parse_multiplicative_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = parse_unary_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var left Expression = rt.Result.(Expression)
	var operators helpers.IntSet = helpers.IntSet{Set: make(map[int]bool)}
	operators.Add(lexer.Multiply)
	operators.Add(lexer.Divide)
	operators.Add(lexer.Modulus)

	for operators.Contains(at(tokens, file_extension).TokenType) {
		var operator = eat(tokens, file_extension).TokenType
		rt = parse_unary_expression(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		left = BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    rt.Result.(Expression),
		}
	}

	return runtime.Success(left)
}

func parse_unary_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var signs helpers.IntSet = helpers.IntSet{Set: make(map[int]bool)}
	signs.Add(lexer.Plus)
	signs.Add(lexer.Minus)

	var rt runtime.RuntimeResult
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	if !signs.Contains(at(tokens, file_extension).TokenType) {
		rt = parse_factorial_expression(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	}

	var sign = lexer.Plus
	for signs.Contains(at(tokens, file_extension).TokenType) {
		if eat(tokens, file_extension).TokenType == lexer.Plus {
			continue
		}

		if sign == lexer.Plus {
			sign = lexer.Minus
		} else {
			sign = lexer.Plus
		}
	}

	rt = parse_factorial_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(UnaryExpression{
		Sign:  sign,
		Value: rt.Result.(Expression),
	})
}

func parse_factorial_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = parse_call_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var level uint64 = 0
	for at(tokens, file_extension).TokenType == lexer.Exclamation {
		level++
		eat(tokens, file_extension)
	}

	if level == 0 {
		return runtime.Success(rt.Result)
	}

	return runtime.Success(Factorial{
		Value: rt.Result.(Expression),
		Level: level,
	})
}

func parse_call_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = parse_primary_expression(tokens, file_extension)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	if at(tokens, file_extension).TokenType != lexer.LeftParentheses {
		return runtime.Success(rt.Result)
	}

	var callee Expression = rt.Result.(Expression)
	eat(tokens, file_extension)

	var arguments []Expression = []Expression{}
	for at(tokens, file_extension).TokenType != lexer.RightParentheses {
		rt = parse_expression(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		arguments = append(arguments, rt.Result.(Expression))
		if at(tokens, file_extension).TokenType == lexer.RightParentheses {
			break
		}

		rt = expect(tokens, file_extension, lexer.Comma)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}
	}

	eat(tokens, file_extension)
	return runtime.Success(CallExpression{
		Callee: callee,
		Arguments: arguments,
	})
}

func parse_primary_expression(tokens *[]lexer.Token, file_extension string) runtime.RuntimeResult {
	var token lexer.Token = eat(tokens, file_extension)
	switch token.TokenType {
	case lexer.Int:
		value, _ := strconv.ParseInt(token.Value, 10, 64)
		return runtime.Success(Int{Value: value})
	case lexer.Double:
		value, _ := strconv.ParseFloat(token.Value, 64)
		return runtime.Success(Double{Value: value})
	case lexer.True:
		return runtime.Success(Boolean{Value: true})
	case lexer.False:
		return runtime.Success(Boolean{Value: false})
	case lexer.LeftParentheses:
		var rt runtime.RuntimeResult = parse_expression(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		var returned Expression = rt.Result.(Expression)
		rt = expect(tokens, file_extension, lexer.RightParentheses)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(returned)
	case lexer.Identifier:
		return runtime.Success(Identifier{VariableName: token.Value})
	case lexer.Pipe:
		var rt runtime.RuntimeResult = parse_expression(tokens, file_extension)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		var returned Expression = rt.Result.(Expression)
		rt = expect(tokens, file_extension, lexer.Pipe)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(AbsoluteValue{Value: returned})
	default:
		return runtime.Failure(runtime.SyntaxError{Reason: fmt.Sprintf("Invalid syntax. Token: {%d %s}", token.TokenType, token.Value)})
	}
}