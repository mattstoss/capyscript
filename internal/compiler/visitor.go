package compiler

type Node interface {
	Accept(Visitor)
}

type Visitor interface {
}
