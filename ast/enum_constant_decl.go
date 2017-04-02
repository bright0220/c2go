package ast

type EnumConstantDecl struct {
	Address   string
	Position  string
	Position2 string
	Name      string
	Type      string
}

func parseEnumConstantDecl(line string) EnumConstantDecl {
	groups := groupsFromRegex(
		"<(?P<position>.*)>(?P<position2> [^ ]+)? (?P<name>.+) '(?P<type>.+?)'",
		line,
	)

	return EnumConstantDecl{
		Address: groups["address"],
		Position: groups["position"],
		Position2: groups["position2"],
		Name: groups["name"],
		Type: groups["type"],
	}
}
