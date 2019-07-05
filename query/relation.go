package query

import (
	"github.com/alexmorten/mgraph/proto"
)

//Relation for building queries
type Relation struct {
	qr *proto.QueryRelation
	Attributes
	to   *Node
	from *Node
}

//NewRelation with initilized pointers
func NewRelation() *Relation {
	return &Relation{
		qr:         &proto.QueryRelation{},
		Attributes: Attributes{},
	}
}

// Key defines the key of the Relation in the query
func (r *Relation) Key(k string) *Relation {
	r.qr.Key = k
	return r
}

//Type defines the type of the Relation in the query
func (r *Relation) Type(t string) *Relation {
	r.qr.Type = t
	return r
}

//To declares a Node the relation is pointing to
func (r *Relation) To() *Node {
	if r.from != nil {
		panic("called *Relation.To() on a relation with already set from")
	}

	r.to = NewNode()
	return r.to
}

//From declares a Node the relation is originating from
func (r *Relation) From() *Node {
	if r.to != nil {
		panic("called *Relation.To() on a relation with already set from")
	}

	r.to = NewNode()
	return r.to
}

//GetQueryRelation ...
func (r *Relation) GetQueryRelation() *proto.QueryRelation {
	r.qr.Attributes = r.Attributes
	if r.to != nil {
		r.qr.Direction = &proto.QueryRelation_To{
			To: r.to.GetQueryNode(),
		}
	}

	if r.from != nil {
		r.qr.Direction = &proto.QueryRelation_From{
			From: r.from.GetQueryNode(),
		}
	}

	return r.qr
}
