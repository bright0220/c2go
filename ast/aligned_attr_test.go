package ast

import (
	"testing"
)

func TestAlignedAttr(t *testing.T) {
	nodes := map[string]Node{
		`0x7f8a1d8ccfd0 <col:47, col:57> aligned`: &AlignedAttr{
			Addr:       0x7f8a1d8ccfd0,
			Position:   "col:47, col:57",
			IsAligned:  true,
			ChildNodes: []Node{},
		},
		`0x2c8ba10 <col:42>`: &AlignedAttr{
			Addr:       0x2c8ba10,
			Position:   "col:42",
			IsAligned:  false,
			ChildNodes: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
