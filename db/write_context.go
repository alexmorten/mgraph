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

func (c *writeContext) descendNode(qn *proto.QueryNode) *proto.Node {
	n := proto.ConstructNode(qn)

	for _, qr := range qn.Relations {
		r := c.descendRelation(qr, n.Key)
		n.Relations = append(n.Relations, r)
	}

	c.nodeWrites[n.Key] = n

	return n
}

func (c *writeContext) descendRelation(qr *proto.QueryRelation, parentKey string) *proto.Relation {
	from := qr.GetFrom()
	if from != nil {
		n := c.descendNode(from)

		r := proto.ConstructRelation(qr)
		r.From = n.Key
		r.To = parentKey

		n.Relations = append(n.Relations, r)

		return r
	}

	to := qr.GetTo()
	if to != nil {
		n := c.descendNode(to)

		r := proto.ConstructRelation(qr)
		r.To = n.Key
		r.From = parentKey

		n.Relations = append(n.Relations, r)

		return r
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
