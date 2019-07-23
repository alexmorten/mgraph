package db

import (
	"github.com/alexmorten/mgraph/proto"
	"github.com/dgraph-io/badger"
	pb "github.com/golang/protobuf/proto"
)

const nodeIndexPrefix = "nodes/"

func pathForNode(k string) []byte {
	return []byte(nodeIndexPrefix + k)
}

func writeNodeIntoIndex(txn *badger.Txn, n *proto.Node) error {
	b, err := pb.Marshal(n)
	if err != nil {
		return err
	}
	return txn.Set(pathForNode(n.Key), b)
}
