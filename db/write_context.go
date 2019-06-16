package db

import (
	"github.com/alexmorten/mgraph/proto"
	"github.com/dgraph-io/badger"
	pb "github.com/golang/protobuf/proto"
)

type writeContext struct {
	txn        *badger.Txn
	nodeWrites map[string]*proto.Node
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
		n.Relations = append(n.Relations, r)
		wn.Relations = append(wn.Relations, wr)
	}

	c.nodeWrites[n.Key] = n

	return n, wn
}

func (c *writeContext) descendRelation(qr *proto.QueryRelation, parentKey string) (*proto.Relation, *proto.QueryRelation) {
	from := qr.GetFrom()
	if from != nil {
		n, wn := c.descendNode(from)

		r := proto.ConstructRelation(qr)
		r.From = n.Key
		r.To = parentKey
		n.Relations = append(n.Relations, r)

		wr := proto.ConstructQueryRelation(r)
		wr.Direction = &proto.QueryRelation_From{From: wn}

		return r, wr
	}

	to := qr.GetTo()
	if to != nil {
		n, wn := c.descendNode(to)

		r := proto.ConstructRelation(qr)
		r.To = n.Key
		r.From = parentKey
		n.Relations = append(n.Relations, r)

		wr := proto.ConstructQueryRelation(r)
		wr.Direction = &proto.QueryRelation_To{To: wn}

		return r, wr
	}

	panic("relation without from or to")
}

func (c *writeContext) write() error {
	for _, n := range c.nodeWrites {
		b, err := pb.Marshal(n)
		if err != nil {
			return err
		}

		err = c.txn.Set([]byte(n.Key), b)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
