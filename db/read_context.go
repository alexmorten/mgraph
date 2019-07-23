package db

import (
	"github.com/alexmorten/mgraph/proto"
	"github.com/dgraph-io/badger"
	pb "github.com/golang/protobuf/proto"
)

type readContext struct {
	txn *badger.Txn
}

func (c *readContext) descendNode(readStatementNode *proto.QueryNode, key string) *proto.QueryNode {
	if readStatementNode.Key != "" && readStatementNode.Key != key {
		return nil
	}

	n := c.findNode(key)
	if n == nil {
		return nil
	}

	qn := proto.ConstructQueryNode(n)
	if !readStatementNode.Matches(qn) {
		return nil
	}
	for _, statementRelation := range readStatementNode.Relations {
		foundAtLeastOneMatch := false
		omitNodeKey := func(k string) bool {
			if statementRelation.RelatedNode() == nil || statementRelation.RelatedNode().Key == "" {
				return false
			}

			return statementRelation.RelatedNode().Key != k
		}
		for relationIteration := range iterateRelationsForNode(c.txn, n.Key, omitNodeKey) {
			if relationIteration.Err != nil {
				panic(relationIteration.Err)
			}
			relation := relationIteration.Relation
			qr := relation.ConstructQueryRelation(n)
			if statementRelation.Matches(qr) {
				relatedNode := statementRelation.RelatedNode()
				if relatedNode == nil {
					continue
				}
				relatedKey := relatedNode.Key

				if relatedKey == "" {
					relatedKey = relation.OtherSideKey(n.Key)
				}

				matchedNode := c.descendNode(relatedNode, relatedKey)
				if matchedNode != nil {
					matchedRelation := relation.ConstructQueryRelation(n)
					matchedRelation.OverwriteRelatedNode(matchedNode)

					qn.Relations = append(qn.Relations, matchedRelation)

					foundAtLeastOneMatch = true
				}
			}
		}

		if !foundAtLeastOneMatch {
			return nil
		}
	}

	return qn
}

func (c *readContext) findNode(key string) *proto.Node {
	foundItem, err := c.txn.Get(pathForNode(key))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil
		}
		panic(err)
	}

	n, err := parseNode(foundItem)
	if err != nil {
		panic(err)
	}
	return n
}

func parseNode(item *badger.Item) (*proto.Node, error) {
	b, err := item.Value()
	if err != nil {
		return nil, err
	}
	n := &proto.Node{}
	err = pb.Unmarshal(b, n)
	if err != nil {
		return nil, err
	}

	return n, nil
}
