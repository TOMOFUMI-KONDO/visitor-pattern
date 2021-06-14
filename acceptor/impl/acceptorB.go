package impl

import (
	"example.com/visitor-pattern/acceptor/_interface"
)

func NewAcceptorB() _interface.Acceptor {
	return AcceptorB{name: "acceptorB"}
}

type AcceptorB struct {
	name string
}

func (impl AcceptorB) Accept(visitor _interface.Visitor) {
	visitor.Visit(impl)
}

func (impl AcceptorB) GetName() string {
	return impl.name
}
