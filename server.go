package mgraph

import (
	"fmt"

	"github.com/alexmorten/mgraph/db"
	"github.com/alexmorten/mgraph/proto"
	json "github.com/golang/protobuf/jsonpb"

	"bytes"
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
	var lastResponse *proto.QueryResponse
	for i := 0; i < 10; i++ {
		before := time.Now()
		response, err := db.Update(q)
		fmt.Println(float64(time.Since(before).Nanoseconds())/1000000.0, "ms")
		he(err)
		m := json.Marshaler{Indent: " "}
		b := &bytes.Buffer{}
		he(m.Marshal(b, response))
		// fmt.Println(b.String())
		lastResponse = response
	}

	married := &proto.QueryRelation{Direction: &proto.QueryRelation_To{To: &proto.QueryNode{}}}

	tom := &proto.QueryNode{Key: lastResponse.Result[0].Root.Key, Relations: []*proto.QueryRelation{married}}
	findStatement := &proto.FindStatement{Root: tom}
	before := time.Now()
	response, err := db.Find(findStatement)
	fmt.Println("Find time: ")

	fmt.Println(float64(time.Since(before).Nanoseconds())/1000000.0, "ms")

	if err != nil {
		panic(err)
	}

	m := json.Marshaler{Indent: " "}
	b := &bytes.Buffer{}
	he(m.Marshal(b, response))
	fmt.Println(b.String())

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
