package query

import (
	"github.com/alexmorten/mgraph/proto"
)

//Statement is a general query builder statement
type Statement interface {
	Statement() *proto.Statement
}

//CreateStatement for building create statements
type CreateStatement struct {
	root *Node
}

//NewCreate statement for building queries
func NewCreate() *CreateStatement {
	return &CreateStatement{}
}

//Root creates and returns a root node for building a create query
func (s *CreateStatement) Root() *Node {
	s.root = NewNode()
	return s.root
}

//Query returns the built statement itself
func (s *CreateStatement) Query() *proto.Statement_Create {
	return &proto.Statement_Create{
		Create: &proto.CreateStatement{
			Root: s.root.GetQueryNode(),
		},
	}
}

//Statement returns the statement in the proto.Statement wrapper
func (s *CreateStatement) Statement() *proto.Statement {
	return &proto.Statement{
		Type: s.Query(),
	}
}

//FindStatement for building create statements
type FindStatement struct {
	root *Node
}

//NewFind statement for building queries
func NewFind() *FindStatement {
	return &FindStatement{}
}

//Root creates and returns a root node for building a find query
func (s *FindStatement) Root() *Node {
	s.root = NewNode()
	return s.root
}

//Query returns the built statement itself
func (s *FindStatement) Query() *proto.Statement_Find {
	return &proto.Statement_Find{
		Find: &proto.FindStatement{
			Root: s.root.GetQueryNode(),
		},
	}
}

//Statement returns the statement in the proto.Statement wrapper
func (s *FindStatement) Statement() *proto.Statement {
	return &proto.Statement{
		Type: s.Query(),
	}
}
