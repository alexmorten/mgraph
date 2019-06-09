package mgraph

import (
	"fmt"

	"github.com/alexmorten/mgraph/db"
)

//Server ..
type Server struct {
	db *db.DB
}

//NewServer with initialized DB
func NewServer() *Server {
	d := db.NewDB()
	fmt.Println("before init: ")
	d.Print()
	err := d.Init()
	if err != nil {
		panic(err)
	}
	fmt.Println("after init: ")
	d.Print()
	n1 := db.NewNode()
	n2 := db.NewNode()

	he(d.AddNode(n1))
	he(d.AddNode(n2))

	err = d.AddRelation(db.NewRelation(n1.Key, "BELONGS_TO", n2.Key))
	he(err)
	fmt.Println("after adding stuff: ")
	d.Print()

	return &Server{
		db: d,
	}
}

func he(err error) {
	if err != nil {
		panic(err)
	}
}
