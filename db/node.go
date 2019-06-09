package db

import (
	"fmt"
	"time"

	"github.com/google/btree"
)

//Node ..
type Node struct {
	Key
	WrittenAt time.Time
	Relations []Relation
}

//NewNode with random key and WrittenAt now
func NewNode() *Node {
	return &Node{
		Key:       NewKey(),
		WrittenAt: time.Now(),
	}
}

//Less than other btree.Item
func (n *Node) Less(other btree.Item) bool {
	return n.Key.Less(other)
}

//AddOrReplaceRelation adds the relation to the node if the key does not exist in the relations yet.
//Otherwise it replaces the existing relation with the given relation.
func (n *Node) AddOrReplaceRelation(r Relation) {
	for index, existingR := range n.Relations {
		if existingR.Key == r.Key {
			n.Relations[index] = r
			return
		}
	}

	n.Relations = append(n.Relations, r)
}

func (n *Node) String() string {
	nodeInfo := fmt.Sprintf("key: %s writtenAt: %s \n", n.Key, n.WrittenAt.String())
	nodeInfo += "Relations: \n"
	for _, r := range n.Relations {
		if r.From == n.Key {
			nodeInfo += fmt.Sprintf("  -[%v]-> %v \n", r.Type, r.To)
		} else {
			nodeInfo += fmt.Sprintf("  <-[%v]- %v \n", r.Type, r.From)
		}
	}
	return nodeInfo
}
