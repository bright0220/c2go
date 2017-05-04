package ast

type DoStmt struct {
	Address  string
	Position string
	Children []interface{}
}

func parseDoStmt(line string) *DoStmt {
	groups := groupsFromRegex(
		"<(?P<position>.*)>",
		line,
	)

	return &DoStmt{
		Address:  groups["address"],
		Position: groups["position"],
		Children: []interface{}{},
	}
}

func (n *DoStmt) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
