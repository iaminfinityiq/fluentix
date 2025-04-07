package parser

type Statement interface {
	Kind() int
}

type Expression interface {
	Statement
	ExpressionConfirm()
}