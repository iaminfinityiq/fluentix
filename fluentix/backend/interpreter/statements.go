package interpreter

import (
	"fluentix/backend/scopes"
	"fluentix/backend/value_types"
	"fluentix/frontend/parser"
	"fluentix/runtime"
	"fmt"
)

func EvaluateBlock(ast_node *parser.Block, scope *scopes.Scope) runtime.RuntimeResult {
	var last_evaluated any = nil
	for _, statement := range ast_node.Body {
		var rt runtime.RuntimeResult = Evaluate(&statement, scope)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		last_evaluated = rt.Result
	}

	return runtime.Success(last_evaluated)
}

func EvaluateIfStatement(ast_node *parser.IfStatement, scope *scopes.Scope) runtime.RuntimeResult {
	for {
		__, _ := ast_node.Condition.(parser.Statement)
		var rt runtime.RuntimeResult = Evaluate(&__, scope)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		if rt.Result.(value_types.Object).Name_() != "boolean" {
			return runtime.Failure(runtime.DataTypeError{Reason: fmt.Sprintf("Expected a boolean in the condition, got %s", rt.Result.(value_types.Object).Name_())})
		}

		if rt.Result.(value_types.ValuedObject).Value.(bool) {
			rt = EvaluateBlock(&ast_node.Body, scope)
			if rt.Error != nil {
				return runtime.Failure(*rt.Error)
			}

			return runtime.Success(rt.Result)
		}

		if ast_node.Next == nil {
			return runtime.Success(nil)
		}

		ast_node = ast_node.Next
	}
}

func EvaluateFunctionDeclaration(ast_node *parser.FunctionDeclaration, scope *scopes.Scope) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = scope.Assign(ast_node.FunctionName, MakeDefinedFunction(ast_node.Arguments, ast_node.Body, scope))
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(nil)
}