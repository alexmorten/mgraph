package mgraph

import (
	"fmt"

	"github.com/alexmorten/mgraph/db"
	"github.com/alexmorten/mgraph/proto"

	"time"
)

//Server ..
type Server struct {
	db *db.DB
}

//NewServer with initialized DB
func NewServer() *Server {
	db, err := db.NewDB(db.DefaultConfig())
	he(err)
	writeStatement := &proto.Statement_Create{
		Create: &proto.CreateStatement{
			Root: &proto.QueryNode{
				Attributes: map[string]*proto.AttributeValue{
					"name": &proto.AttributeValue{Value: &proto.AttributeValue_StringValue{StringValue: "Tom"}},
				},
				Relations: []*proto.QueryRelation{
					&proto.QueryRelation{
						Type: "MARRIED",
						Attributes: map[string]*proto.AttributeValue{
							"since": &proto.AttributeValue{Value: &proto.AttributeValue_IntValue{IntValue: time.Now().Unix()}},
						},
						Direction: &proto.QueryRelation_To{To: &proto.QueryNode{
							Attributes: map[string]*proto.AttributeValue{
								"name": &proto.AttributeValue{Value: &proto.AttributeValue_StringValue{StringValue: "Jenny"}},
							},
						},
						},
					},
				},
			},
		},
	}

	q := &proto.Query{
		Statements: []*proto.Statement{&proto.Statement{Type: writeStatement}},
	}

	fmt.Println(db.Update(q))

	db.Find("50e687b9-3c36-414e-adc0-73d49e0aa57f")
	db.Shutdown()
	return &Server{
		db: db,
	}
}

func he(err error) {
	if err != nil {
		panic(err)
	}
}
