package parser

type Block struct {
	Body []Statement
}

func (b Block) Kind() int {
	return BlockStmt
}

type IfStatement struct {
	// this is a node to a linked list of if statement nodes for an if statement (probably full or not)
	Condition Expression // if this is an else, then this value will be set to true
	Body Block // yk what that is lol
	Next *IfStatement // nil if there's nothing left...
}

func (i IfStatement) Kind() int {
	return IfStmt
}

type FunctionDeclaration struct {
	FunctionName string
	Arguments []string
	Body Block
}

func (f FunctionDeclaration) Kind() int {
	return FunctionDeclarationStmt
}