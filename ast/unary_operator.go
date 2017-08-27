package ast

type UnaryOperator struct {
	Addr     Address
	Position string
	Type     string
	IsLvalue bool
	IsPrefix bool
	Operator string
	Children []Node
}

func parseUnaryOperator(line string) *UnaryOperator {
	groups := groupsFromRegex(
		`<(?P<position>.*)>
		 '(?P<type>.*?)'
		(?P<lvalue> lvalue)?
		(?P<prefix> prefix)?
		(?P<postfix> postfix)?
		 '(?P<operator>.*?)'`,
		line,
	)

	return &UnaryOperator{
		Addr:     ParseAddress(groups["address"]),
		Position: groups["position"],
		Type:     groups["type"],
		IsLvalue: len(groups["lvalue"]) > 0,
		IsPrefix: len(groups["prefix"]) > 0,
		Operator: groups["operator"],
		Children: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *UnaryOperator) AddChild(node Node) {
	n.Children = append(n.Children, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *UnaryOperator) Address() Address {
	return n.Addr
}
