package db

import (
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/alexmorten/mgraph/chlog"
	"github.com/google/btree"
)

var (
	//ErrNodeNotFound will be returned when the node was searched for  but couldn't be found
	ErrNodeNotFound = errors.New("Node could not be found")

	//ErrUnexpectedChangeType ...
	ErrUnexpectedChangeType = errors.New("unexpected type of change when creating DB")
)

//DB resides in memory
type DB struct {
	tree         *btree.BTree
	changeWriter *chlog.Writer
}

//NewDB with initialized tree
func NewDB() *DB {
	registerGobTypes()
	return &DB{
		tree: btree.New(2),
	}
}

//Init DB from disk and set changeWriter
func (db *DB) Init() error {
	logConfig := chlog.DefaultConfig()
	changeChan, err := chlog.ReadAllChanges(logConfig)

	if err != nil && err != chlog.ErrNoChangeLog {
		return err
	}

	if err == chlog.ErrNoChangeLog {
		fmt.Println(err.Error(), "starting with empty DB")
	} else {
		fmt.Println("Initializing DB from changelog...")
		for change := range changeChan {
			data := change.Data
			switch data.(type) {
			case *Node:
				n := data.(*Node)
				db.addNode(n)
			default:
				return ErrUnexpectedChangeType
			}
		}
		fmt.Println("DB filled with values")
	}

	changeWriter, err := chlog.NewWriter(logConfig)
	if err != nil {
		return err
	}
	db.changeWriter = changeWriter

	return nil
}

//Print the contents of the DB
func (db *DB) Print() {
	db.tree.Ascend(func(item btree.Item) bool {
		fmt.Println(item.(*Node))
		return true
	})
}

//AddNode to DB
func (db *DB) AddNode(n *Node) error {
	err := db.changeWriter.WriteChange(&chlog.Change{
		Ts:   time.Now(),
		Data: n,
	})
	if err != nil {
		return err
	}
	db.addNode(n)

	return nil
}

func (db *DB) addNode(n *Node) {
	db.tree.ReplaceOrInsert(n)
}

//ReplaceNode in DB
func (db *DB) ReplaceNode(n *Node) error {
	err := db.replaceNode(n)
	if err != nil {
		return err
	}

	err = db.changeWriter.WriteChange(&chlog.Change{
		Ts:   time.Now(),
		Data: n,
	})
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) replaceNode(n *Node) error {
	if db.tree.Has(n) {
		db.tree.ReplaceOrInsert(n)
		return nil
	}

	return ErrNodeNotFound
}

//AddRelation to nodes
func (db *DB) AddRelation(r Relation) error {
	return db.addRelation(r, true)
}

//AddRelation to nodes
func (db *DB) addRelation(r Relation, writeChanges bool) error {
	fromItem := db.tree.Get(r.From)
	toItem := db.tree.Get(r.To)

	if fromItem == nil && toItem == nil {
		return ErrNodeNotFound
	}

	if fromItem != nil {
		fmt.Println("from")
		fromNode := fromItem.(*Node)
		r.from = fromNode
		fromNode.AddOrReplaceRelation(r)

		if writeChanges {
			err := db.changeWriter.WriteChange(&chlog.Change{
				Ts:   time.Now(),
				Data: fromNode,
			})
			if err != nil {
				return err
			}
		}
	}

	if toItem != nil {
		fmt.Println("to")
		toNode := toItem.(*Node)
		r.from = toNode
		toNode.AddOrReplaceRelation(r)

		if writeChanges {
			err := db.changeWriter.WriteChange(&chlog.Change{
				Ts:   time.Now(),
				Data: toNode,
			})
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func registerGobTypes() {
	gob.Register(&Node{})
	gob.Register(&Relation{})
	gob.Register(&chlog.Change{})
}
