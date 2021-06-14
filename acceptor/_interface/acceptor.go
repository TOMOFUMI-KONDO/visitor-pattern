package _interface

type Acceptor interface {
	Accept(visitor Visitor)
}
