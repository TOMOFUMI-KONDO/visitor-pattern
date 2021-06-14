package _interface

type Visitor interface {
	Visit(acceptor Acceptor)
}
