package db

import (
	"github.com/alexmorten/mgraph/proto"
	"github.com/dgraph-io/badger"
)

type writeContext struct {
	txn            *badger.Txn
	nodeWrites     map[string]*proto.Node
	relationWrites []*proto.Relation
}

func newWriteContext(txn *badger.Txn) *writeContext {
	return &writeContext{
		txn:        txn,
		nodeWrites: map[string]*proto.Node{},
	}
}

func (c *writeContext) descendNode(writeStatementNode *proto.QueryNode) (*proto.Node, *proto.QueryNode) {
	n := proto.ConstructNode(writeStatementNode)
	wn := proto.ConstructQueryNode(n)
	for _, qr := range writeStatementNode.Relations {
		r, wr := c.descendRelation(qr, n.Key)
		c.relationWrites = append(c.relationWrites, r)
		wn.Relations = append(wn.Relations, wr)
	}

	c.nodeWrites[n.Key] = n

	return n, wn
}

func (c *writeContext) descendRelation(qr *proto.QueryRelation, parentKey string) (*proto.Relation, *proto.QueryRelation) {
	from := qr.GetFrom()
	if from != nil {
		n, wn := c.descendNode(from)

		r := qr.ConstructRelation()
		r.From = n.Key
		r.To = parentKey
		c.relationWrites = append(c.relationWrites, r)

		wr := r.ConstructQueryRelation(n)
		wr.Direction = &proto.QueryRelation_From{From: wn}

		return r, wr
	}

	to := qr.GetTo()
	if to != nil {
		n, wn := c.descendNode(to)

		r := qr.ConstructRelation()
		r.To = n.Key
		r.From = parentKey
		c.relationWrites = append(c.relationWrites, r)

		wr := r.ConstructQueryRelation(n)
		wr.Direction = &proto.QueryRelation_To{To: wn}

		return r, wr
	}

	panic("relation without from or to")
}

func (c *writeContext) write() error {
	for _, n := range c.nodeWrites {
		err := writeNodeIntoIndex(c.txn, n)
		if err != nil {
			return err
		}
	}

	for _, relation := range c.relationWrites {
		err := writeRelationIntoIndex(c.txn, relation)
		if err != nil {
			return err
		}
	}

	return nil
}
