package main

import (
	acceptor "example.com/visitor-pattern/acceptor/impl"
	visitor "example.com/visitor-pattern/visitor/impl"
)

func main() {
	visitorA := visitor.NewVisitorA()
	visitorB := visitor.NewVisitorB()
	visitorC := visitor.NewVisitorC()
	acceptorA := acceptor.NewAcceptorA()
	acceptorB := acceptor.NewAcceptorB()

	acceptorA.Accept(visitorA)
	acceptorA.Accept(visitorB)
	acceptorA.Accept(visitorC)
	acceptorB.Accept(visitorA)
	acceptorB.Accept(visitorB)
	acceptorB.Accept(visitorC)
}
