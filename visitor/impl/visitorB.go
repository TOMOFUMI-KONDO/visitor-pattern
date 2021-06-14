package impl

import (
	"fmt"

	acceptorInterface "example.com/visitor-pattern/acceptor/_interface"
	acceptorImpl "example.com/visitor-pattern/acceptor/impl"
)

func NewVisitorB() acceptorInterface.Visitor {
	return implB{name: "visitorB"}
}

type implB struct {
	name string
}

func (impl implB) Visit(acceptor acceptorInterface.Acceptor) {
	switch acceptor.(type) {
	case acceptorImpl.AcceptorA:
		fmt.Println(impl.name + "が" + acceptor.(acceptorImpl.AcceptorA).GetName() + "を訪問しました。")
	case acceptorImpl.AcceptorB:
		fmt.Println(impl.name + "が" + acceptor.(acceptorImpl.AcceptorB).GetName() + "を訪問しました。")
	}
}
