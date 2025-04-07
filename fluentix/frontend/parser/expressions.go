package parser

type BinaryExpression struct {
	Left Expression
	Operator int
	Right Expression
}

func (b BinaryExpression) Kind() int {
	return BinaryExpr
}

func (b BinaryExpression) ExpressionConfirm() {

}

type UnaryExpression struct {
	Sign int
	Value Expression
}

func (u UnaryExpression) Kind() int {
	return UnaryExpr
}

func (u UnaryExpression) ExpressionConfirm() {

}

type AbsoluteValue struct {
	Value Expression
}

func (a AbsoluteValue) Kind() int {
	return AbsoluteValueExpr
}

func (a AbsoluteValue) ExpressionConfirm() {

}

type Factorial struct {
	Value Expression
	Level uint64
}

func (f Factorial) Kind() int {
	return FactorialExpr
}

func (f Factorial) ExpressionConfirm() {

}

type AssignmentExpression struct {
	Assignee Expression
	Value Expression
}

func (a AssignmentExpression) Kind() int {
	return AssignmentExpr
}

func (a AssignmentExpression) ExpressionConfirm() {

}

type CallExpression struct {
	Callee Expression
	Arguments []Expression
}

func (c CallExpression) Kind() int {
	return CallExpr
}

func (c CallExpression) ExpressionConfirm() {
	
}

type Identifier struct {
	VariableName string
}

func (i Identifier) Kind() int {
	return IdentifierExpr
}

func (i Identifier) ExpressionConfirm() {

}

type Int struct {
	Value int64
}

func (i Int) Kind() int {
	return IntExpr
}

func (i Int) ExpressionConfirm() {

}

type Double struct {
	Value float64
}

func (d Double) Kind() int {
	return DoubleExpr
}

func (d Double) ExpressionConfirm() {

}

type Boolean struct {
	Value bool
}

func (b Boolean) Kind() int {
	return BooleanExpr
}

func (b Boolean) ExpressionConfirm() {

}