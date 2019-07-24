package db

import (
	"github.com/alexmorten/mgraph/proto"
	"github.com/dgraph-io/badger"
	pb "github.com/golang/protobuf/proto"
)

const relationIndexPrefix = "relations/"
const relationPrimaryPrefix = relationIndexPrefix + "primary/"
const relationToPrefix = relationIndexPrefix + "to/"
const relationFromPrefix = relationIndexPrefix + "from/"

func pathForRelation(key string) []byte {
	return []byte(relationPrimaryPrefix + key)
}

func relationToPrefixForNode(key string) []byte {
	return []byte(relationToPrefix + key + "/")
}

func relationToIndexPath(fromKey, toKey, relationKey string) []byte {
	return []byte(relationToPrefix + fromKey + "/" + toKey + "/" + relationKey)
}

func relationFromPrefixForNode(key string) []byte {
	return []byte(relationFromPrefix + key + "/")
}

func relationFromIndexPath(toKey, fromKey, relationKey string) []byte {
	return []byte(relationToPrefix + toKey + "/" + fromKey + "/" + relationKey)
}

func writeRelationIntoIndex(txn *badger.Txn, r *proto.Relation) error {
	b, err := pb.Marshal(r)
	if err != nil {
		return err
	}

	err = txn.Set(pathForRelation(r.Key), b)
	if err != nil {
		return err
	}

	toPath := relationToIndexPath(r.From, r.To, r.Key)
	err = txn.Set(toPath, []byte{})
	if err != nil {
		return err
	}

	fromPath := relationToIndexPath(r.To, r.From, r.Key)
	err = txn.Set(fromPath, []byte{})
	if err != nil {
		return err
	}

	return nil
}
