package query

import (
	"github.com/alexmorten/mgraph/proto"
)

//Node for building queries
type Node struct {
	qn        *proto.QueryNode
	relations []*Relation

	Attributes
}

//NewNode with initilized pointers
func NewNode() *Node {
	return &Node{
		qn:         &proto.QueryNode{},
		Attributes: Attributes{},
	}
}

// Key defines the key of the node in the query
func (n *Node) Key(k string) *Node {
	n.qn.Key = k
	return n
}

//Type defines the type of the node in the query
func (n *Node) Type(t string) *Node {
	n.qn.Type = t
	return n
}

//Related adds a to the node
func (n *Node) Related() *Relation {
	r := NewRelation()
	n.relations = append(n.relations, r)
	return r
}

//GetQueryNode from Node query builder
func (n *Node) GetQueryNode() *proto.QueryNode {
	n.qn.Attributes = n.Attributes
	for _, r := range n.relations {

		n.qn.Relations = append(n.qn.Relations, r.GetQueryRelation())
	}
	return n.qn
}
