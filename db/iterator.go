package db

import (
	"errors"

	"github.com/alexmorten/mgraph/proto"
	"github.com/dgraph-io/badger"

	"strings"

	pb "github.com/golang/protobuf/proto"
)

type relationIteration struct {
	Relation *proto.Relation
	Err      error
}

func iterateRelationsOverPrefix(txn *badger.Txn, prefix []byte, omitOtherNodeKey func(k string) bool) <-chan relationIteration {
	iteratorChan := make(chan relationIteration, 0)
	go func() {
		defer close(iteratorChan)

		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			keyOnlyitem := it.Item()
			k := keyOnlyitem.Key()

			otherNodeAndRelationKey := string(k[len(prefix):])
			keys := strings.Split(otherNodeAndRelationKey, "/")
			if len(keys) != 2 {
				err := errors.New("found invalid key when iterating over relationship index: " + string(k))
				iteratorChan <- relationIteration{Err: err}
				break
			}

			otherNodeKey := keys[0]
			relationKey := keys[1]

			if omitOtherNodeKey != nil && omitOtherNodeKey(otherNodeKey) {
				continue
			}
			relationItem, err := txn.Get(pathForRelation(relationKey))
			if err != nil {
				iteratorChan <- relationIteration{Err: err}
				break
			}

			v, err := relationItem.Value()
			if err != nil {
				iteratorChan <- relationIteration{Err: err}
				break
			}

			r := &proto.Relation{}
			err = pb.Unmarshal(v, r)
			if err != nil {
				iteratorChan <- relationIteration{Err: err}
				break
			}

			iteratorChan <- relationIteration{Relation: r}
		}
	}()

	return iteratorChan
}

func iterateRelationsForNode(txn *badger.Txn, nodeKey string, omitOtherNodeKey func(k string) bool) <-chan relationIteration {
	collectedChan := make(chan relationIteration, 0)

	go func() {
		defer close(collectedChan)

		toPrefix := relationToPrefixForNode(nodeKey)
		toItChan := iterateRelationsOverPrefix(txn, toPrefix, omitOtherNodeKey)
		for iteration := range toItChan {
			collectedChan <- iteration

			if iteration.Err != nil {
				return
			}
		}

		fromPrefix := relationFromPrefixForNode(nodeKey)
		fromItChan := iterateRelationsOverPrefix(txn, fromPrefix, omitOtherNodeKey)
		for iteration := range fromItChan {
			collectedChan <- iteration

			if iteration.Err != nil {
				return
			}
		}
	}()

	return collectedChan
}
