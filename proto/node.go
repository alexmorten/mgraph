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
