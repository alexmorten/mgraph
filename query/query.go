package query

import (
	"github.com/alexmorten/mgraph/proto"
)

//Query for building DB queries
type Query struct {
	query      *proto.Query
	statements []Statement
}

//New Query for building DB queries
func New() *Query {
	return &Query{}
}

//Create adds a CreateStatement to the query
func (q *Query) Create() *CreateStatement {
	s := NewCreate()
	q.statements = append(q.statements, s)

	return s
}

//Find adds a FindStatement to the query
func (q *Query) Find() *FindStatement {
	s := NewFind()
	q.statements = append(q.statements, s)

	return s
}

//Build proto Query
func (q *Query) Build() *proto.Query {
	q.query = &proto.Query{}
	for _, s := range q.statements {
		q.query.Statements = append(q.query.Statements, s.Statement())
	}

	return q.query
}
