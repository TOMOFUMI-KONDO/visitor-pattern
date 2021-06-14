package impl

import (
	"fmt"

	acceptorInterface "example.com/visitor-pattern/acceptor/_interface"
	acceptorImpl "example.com/visitor-pattern/acceptor/impl"
)

func NewVisitorC() acceptorInterface.Visitor {
	return implC{name: "visitorC", present: "花束"}
}

type implC struct {
	name    string
	present string
}

func (impl implC) Visit(acceptor acceptorInterface.Acceptor) {
	switch acceptor.(type) {
	case acceptorImpl.AcceptorA:
		fmt.Println(impl.name + "が" + impl.present + "を渡しに" + acceptor.(acceptorImpl.AcceptorA).GetName() + "を訪問しました。")
	case acceptorImpl.AcceptorB:
		fmt.Println(impl.name + "が" + impl.present + "を渡しに" + acceptor.(acceptorImpl.AcceptorB).GetName() + "を訪問しました。")
	}
}
