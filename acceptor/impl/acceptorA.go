package impl

import (
	"example.com/visitor-pattern/acceptor/_interface"
)

func NewAcceptorA() _interface.Acceptor {
	return AcceptorA{name: "acceptorA"}
}

type AcceptorA struct {
	name string
}

func (impl AcceptorA) Accept(visitor _interface.Visitor) {
	visitor.Visit(impl)
}

func (impl AcceptorA) GetName() string {
	return impl.name
}
