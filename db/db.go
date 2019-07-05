package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alexmorten/mgraph/proto"
	"github.com/dgraph-io/badger"
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
	response := &proto.QueryResponse{}

	err := db.store.Update(func(txn *badger.Txn) error {
		for _, statement := range query.Statements {
			s := statement.GetCreate()
			ctx := newWriteContext(txn)
			_, writtenRoot := ctx.descendNode(s.Root)
			err := ctx.write()
			if err != nil {
				return err
			}
			response.Result = append(response.Result, &proto.StatementResult{Root: writtenRoot})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}

//Find Node
func (db *DB) Find(s *proto.Statement_Find) (*proto.QueryResponse, error) {
	response := &proto.QueryResponse{}

	err := db.store.View(func(txn *badger.Txn) error {
		ctx := &readContext{
			txn: txn,
		}

		readRoot := ctx.descendNode(s.Find.Root, s.Find.Root.Key)
		if readRoot != nil {
			response.Result = append(response.Result, &proto.StatementResult{Root: readRoot})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
