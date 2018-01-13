package visitor

import (
	"fmt"
	"reflect"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

// Dumper prints ast hierarchy to stdout
// Also prints comments and positions attached to nodes
type Dumper struct {
	Indent    string
	Comments  comment.Comments
	Positions position.Positions
}

// EnterNode is invoked at every node in heirerchy
func (d Dumper) EnterNode(n node.Node) bool {

	fmt.Printf("%v%v", d.Indent, reflect.TypeOf(n))
	if p := d.Positions[n]; p != nil {
		fmt.Printf(" %v", *p)
	}
	if a := n.Attributes(); len(a) > 0 {
		fmt.Printf(" %v", a)
	}
	fmt.Println()

	if c := d.Comments[n]; len(c) > 0 {
		fmt.Printf("%vComments:\n", d.Indent+"  ")
		for _, cc := range c {
			fmt.Printf("%v%q\n", d.Indent+"    ", cc)
		}
	}

	return true
}

// GetChildrenVisitor is invoked at every node parameter that contains children nodes
func (d Dumper) GetChildrenVisitor(key string) node.Visitor {
	fmt.Printf("%v%q:\n", d.Indent+"  ", key)
	return Dumper{d.Indent + "    ", d.Comments, d.Positions}
}

// LeaveNode is invoked after node process
func (d Dumper) LeaveNode(n node.Node) {
	// do nothing
}