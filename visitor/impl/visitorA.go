package impl

import (
	"fmt"

	acceptorInterface "example.com/visitor-pattern/acceptor/_interface"
	acceptorImpl "example.com/visitor-pattern/acceptor/impl"
)

func NewVisitorA() acceptorInterface.Visitor {
	return implA{name: "visitorA"}
}

type implA struct {
	name string
}

func (impl implA) Visit(acceptor acceptorInterface.Acceptor) {
	switch acceptor.(type) {
	case acceptorImpl.AcceptorA:
		fmt.Println(impl.name + "が" + acceptor.(acceptorImpl.AcceptorA).GetName() + "を訪問しました。")
	case acceptorImpl.AcceptorB:
		fmt.Println(impl.name + "が" + acceptor.(acceptorImpl.AcceptorB).GetName() + "を訪問しました。")
	}
}
