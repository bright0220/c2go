package ast

type BuiltinType struct {
	Address string
	Type    string
}

func ParseBuiltinType(line string) BuiltinType {
	groups := groupsFromRegex(
		"'(?P<type>.*?)'",
		line,
	)

	return BuiltinType{
		Address: groups["address"],
		Type: groups["type"],
	}
}
