package db

import (
	"github.com/google/btree"
	"github.com/google/uuid"
)

//Key unique for each node and relation. Used for the db-tree and relationship lookups
type Key string

//NewKey returns a random uuid string key
func NewKey() Key {
	return Key(uuid.New().String())
}

//Less than other btree.Item
func (k Key) Less(other btree.Item) bool {
	if otherKey, ok := other.(Key); ok {
		return k < otherKey
	}

	if node, ok := other.(*Node); ok {
		return k < node.Key
	}

	return false
}
