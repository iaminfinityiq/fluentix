package parser

const (
	BlockStmt = iota
	
	// Statements
	IfStmt
	FunctionDeclarationStmt

	// Expressions
	BinaryExpr
	UnaryExpr
	AbsoluteValueExpr
	FactorialExpr
	AssignmentExpr

	// Value types
	IdentifierExpr // This is special
	IntExpr
	DoubleExpr
	BooleanExpr
	CallExpr
)