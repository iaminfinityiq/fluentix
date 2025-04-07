package interpreter

import (
	"fluentix/backend/scopes"
	"fluentix/frontend/parser"
	"fluentix/runtime"
)

func Evaluate(ast_node *parser.Statement, scope *scopes.Scope) runtime.RuntimeResult {
	var rt runtime.RuntimeResult
	switch (*ast_node).Kind() {
	case parser.BlockStmt:
		__, _ := (*ast_node).(parser.Block)
		rt = EvaluateBlock(&__, scope)
	case parser.IfStmt:
		__, _ := (*ast_node).(parser.IfStatement)
		rt = EvaluateIfStatement(&__, scope)
	case parser.FunctionDeclarationStmt:
		__, _ := (*ast_node).(parser.FunctionDeclaration)
		rt = EvaluateFunctionDeclaration(&__, scope)
	case parser.IntExpr:
		__, _ := (*ast_node).(parser.Int)
		rt = EvaluateInt(&__, scope)
	case parser.DoubleExpr:
		__, _ := (*ast_node).(parser.Double)
		rt = EvaluateDouble(&__, scope)
	case parser.BooleanExpr:
		__, _ := (*ast_node).(parser.Boolean)
		rt = EvaluateBoolean(&__, scope)
	case parser.BinaryExpr:
		__, _ := (*ast_node).(parser.BinaryExpression)
		rt = EvaluateBinaryExpression(&__, scope)
	case parser.UnaryExpr:
		__, _ := (*ast_node).(parser.UnaryExpression)
		rt = EvaluateUnaryExpression(&__, scope)
	case parser.AbsoluteValueExpr:
		__, _ := (*ast_node).(parser.AbsoluteValue)
		rt = EvaluateAbsoluteValueExpression(&__, scope)
	case parser.FactorialExpr:
		__, _ := (*ast_node).(parser.Factorial)
		rt = EvaluateFactorialExpression(&__, scope)
	case parser.AssignmentExpr:
		__, _ := (*ast_node).(parser.AssignmentExpression)
		rt = EvaluateAssignmentExpression(&__, scope)
	case parser.IdentifierExpr:
		__, _ := (*ast_node).(parser.Identifier)
		rt = EvaluateIdentifier(&__, scope)
	}

	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(rt.Result)
}