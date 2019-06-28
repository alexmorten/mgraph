package proto

import (
	"github.com/google/uuid"
)

//ConstructNode from QueryNode
func ConstructNode(qn *QueryNode) *Node {
	n := &Node{
		Key:        qn.Key,
		Type:       qn.Type,
		Attributes: qn.Attributes,
	}

	if n.Key == "" {
		n.Key = newKey()
	}

	return n
}

// ConstructQueryNode from node
func ConstructQueryNode(n *Node) *QueryNode {
	return &QueryNode{
		Key:        n.Key,
		Type:       n.Type,
		Attributes: n.Attributes,
	}
}

func newKey() string {
	return uuid.New().String()
}

//Matches checks if all properties given in n match the properties of n2
func (n *QueryNode) Matches(n2 *QueryNode) bool {
	if n.Key != "" && n.Key != n2.Key {
		return false
	}

	if n.Type != "" && n.Type != n2.Type {
		return false
	}

	return AttributesMatch(n.Attributes, n2.Attributes)
}
