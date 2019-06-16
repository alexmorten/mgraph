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

func newKey() string {
	return uuid.New().String()
}
