package interpreter

import (
	"fluentix/backend/scopes"
	"fluentix/backend/value_types"
	"fluentix/frontend/lexer"
	"fluentix/frontend/parser"
	"fluentix/runtime"
	"fluentix/backend/initializer"
	"fmt"
)

func EvaluateInt(ast_node *parser.Int, scope *scopes.Scope) runtime.RuntimeResult {
	return runtime.Success(initializer.MakeInt(ast_node.Value))
}

func EvaluateDouble(ast_node *parser.Double, scope *scopes.Scope) runtime.RuntimeResult {
	return runtime.Success(initializer.MakeDouble(ast_node.Value))
}

func EvaluateBoolean(ast_node *parser.Boolean, scope *scopes.Scope) runtime.RuntimeResult {
	return runtime.Success(initializer.MakeBoolean(ast_node.Value))
}

func EvaluateBinaryExpression(ast_node *parser.BinaryExpression, scope *scopes.Scope) runtime.RuntimeResult {
	__, _ := ast_node.Left.(parser.Statement)
	var rt runtime.RuntimeResult = Evaluate(&__, scope)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var left value_types.Object = rt.Result.(value_types.Object)

	__, _ = ast_node.Right.(parser.Statement)
	rt = Evaluate(&__, scope)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var right value_types.Object = rt.Result.(value_types.Object)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	switch ast_node.Operator {
	case lexer.Plus:
		value, ok_left := left.Properties_()["Plus"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '+' for %s and %s", left.Name_(), right.Name_())})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '+' for %s and %s", left.Name_(), right.Name_())})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.Minus:
		value, ok_left := left.Properties_()["Minus"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '-' for %s and %s", left.Name_(), right.Name_())})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '-' for %s and %s", left.Name_(), right.Name_())})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.Multiply:
		value, ok_left := left.Properties_()["Multiply"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '*' for %s and %s", left.Name_(), right.Name_())})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '*' for %s and %s", left.Name_(), right.Name_())})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.Divide:
		value, ok_left := left.Properties_()["Divide"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '/' for %s and %s", left.Name_(), right.Name_())})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '/' for %s and %s", left.Name_(), right.Name_())})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}

	case lexer.Modulus:
		value, ok_left := left.Properties_()["Modulus"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '%' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '%' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.DoubleEquals:
		value, ok_left := left.Properties_()["Equals"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '=' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '=' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.NotEquals:
		value, ok_left := left.Properties_()["NotEquals"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '!=' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '!=' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.GreaterThan:
		value, ok_left := left.Properties_()["Greater"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '>' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '>' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.SmallerThan:
		value, ok_left := left.Properties_()["Smaller"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '<' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '<' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.GreaterThanOrEquals:
		value, ok_left := left.Properties_()["GreaterThanOrEquals"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '>=' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '>=' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	case lexer.SmallerThanOrEquals:
		value, ok_left := left.Properties_()["SmallerThanOrEquals"]
		if !ok_left {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '<=' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: "Invalid operation '<=' for " + left.Name_() + " and " + right.Name_()})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&right})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	}

	return rt
}

func EvaluateUnaryExpression(ast_node *parser.UnaryExpression, scope *scopes.Scope) runtime.RuntimeResult {
	__, _ := ast_node.Value.(parser.Statement)
	var rt runtime.RuntimeResult = Evaluate(&__, scope)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	if ast_node.Sign == lexer.Plus {
		value, ok := rt.Result.(value_types.Object).Properties_()["Positive"]
		if !ok {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '+' for %s", rt.Result.(value_types.Object).Name_())})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '+' for %s", rt.Result.(value_types.Object).Name_())})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			___, _ := rt.Result.(value_types.Object)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&___})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	}

	if ast_node.Sign == lexer.Minus {
		value, ok := rt.Result.(value_types.Object).Properties_()["Negative"]
		if !ok {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '-' for %s", rt.Result.(value_types.Object).Name_())})
		}

		if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '-' for %s", rt.Result.(value_types.Object).Name_())})
		}

		if value.Name_() == "builtin_function" {
			var new = value.(value_types.ValuedObject)
			___, _ := rt.Result.(value_types.Object)
			rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&___})
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}
	}

	return runtime.RuntimeResult{}
}

func EvaluateAbsoluteValueExpression(ast_node *parser.AbsoluteValue, scope *scopes.Scope) runtime.RuntimeResult {
	__, _ := ast_node.Value.(parser.Statement)
	var rt runtime.RuntimeResult = Evaluate(&__, scope)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	value, ok := rt.Result.(value_types.Object).Properties_()["Absolute"]
	if !ok {
		return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation 'abs' for %s", rt.Result.(value_types.Object).Name_())})
	}

	if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
		return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation 'abs' for %s", rt.Result.(value_types.Object).Name_())})
	}

	if value.Name_() == "builtin_function" {
		var new = value.(value_types.ValuedObject)
		___, _ := rt.Result.(value_types.Object)
		rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&___})
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	}

	return runtime.RuntimeResult{}
}

func EvaluateFactorialExpression(ast_node *parser.Factorial, scope *scopes.Scope) runtime.RuntimeResult {
	__, _ := ast_node.Value.(parser.Statement)
	var rt runtime.RuntimeResult = Evaluate(&__, scope)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	value, ok := rt.Result.(value_types.Object).Properties_()["Factorial"]
	if !ok {
		return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '!' for %s", rt.Result.(value_types.Object).Name_())})
	}

	if value.Name_() != "builtin_function" && value.Name_() != "defined_function" {
		return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Invalid operation '!' for %s", rt.Result.(value_types.Object).Name_())})
	}

	if value.Name_() == "builtin_function" {
		var new = value.(value_types.ValuedObject)
		___, _ := rt.Result.(value_types.Object)
		
		var object value_types.Object = interface{}(initializer.MakeInt(int64(ast_node.Level))).(value_types.Object)
		rt = new.Value.(func([]*value_types.Object) runtime.RuntimeResult)([]*value_types.Object{&object, &___})
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	}

	return runtime.RuntimeResult{}
}

func EvaluateAssignmentExpression(ast_node *parser.AssignmentExpression, scope *scopes.Scope) runtime.RuntimeResult {
	var rt runtime.RuntimeResult
	switch ast_node.Assignee.Kind() {
	case parser.IdentifierExpr:
		var variable_name string = ast_node.Assignee.(parser.Identifier).VariableName
		__, _ := ast_node.Value.(parser.Statement)
		rt = Evaluate(&__, scope)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		var returned value_types.Object = rt.Result.(value_types.Object)

		rt = scope.Assign(variable_name, rt.Result.(value_types.Object))
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(returned)
	default:
		return runtime.Failure(runtime.SyntaxError{Reason: "Invalid syntax!"})
	}
}

func EvaluateIdentifier(ast_node *parser.Identifier, scope *scopes.Scope) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = scope.Resolve(ast_node.VariableName)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(rt.Result)
}