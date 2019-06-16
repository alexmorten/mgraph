package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alexmorten/mgraph/proto"
	"github.com/dgraph-io/badger"
	pb "github.com/golang/protobuf/proto"
)

var (
	//ErrNodeNotFound will be returned when the node was searched for  but couldn't be found
	ErrNodeNotFound = errors.New("Node could not be found")

	//ErrUnexpectedChangeType ...
	ErrUnexpectedChangeType = errors.New("unexpected type of change when creating DB")
)

//DB resides in memory
type DB struct {
	store *badger.DB
}

//Config for DB
type Config struct {
	Dir string
}

//DefaultConfig for DB
func DefaultConfig() Config {
	return Config{
		Dir: "data",
	}
}

//NewDB with initialized tree
func NewDB(c Config) (*DB, error) {
	opts := badger.DefaultOptions
	opts.Dir = c.Dir
	opts.ValueDir = c.Dir

	store, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	db := &DB{
		store: store,
	}
	info := db.store.Tables()
	b, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		return nil, err
	}
	fmt.Println("table info:")
	fmt.Println(string(b))

	return db, nil
}

//Shutdown the db
func (db *DB) Shutdown() error {
	return db.store.Close()
}

//Update the db with the given query
func (db *DB) Update(query *proto.Query) (*proto.QueryResponse, error) {
	err := db.store.Update(func(txn *badger.Txn) error {
		for _, statement := range query.Statements {
			s := statement.GetCreate()
			ctx := newWriteContext(txn)
			n := ctx.descendNode(s.Root)
			fmt.Println(n)
			err := ctx.write()
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	response := &proto.QueryResponse{}
	return response, nil
}

//Find Node
func (db *DB) Find(key string) {
	db.store.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		b, err := item.Value()
		if err != nil {
			return err
		}
		n := &proto.Node{}
		err = pb.Unmarshal(b, n)
		if err != nil {
			return err
		}

		fmt.Println(n)
		return nil
	})
}
