package mgraph

import (
	"fmt"

	"github.com/alexmorten/mgraph/db"
	"github.com/alexmorten/mgraph/proto"
	"github.com/alexmorten/mgraph/query"
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

	q := buildCreateQuery()
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

	before := time.Now()
	response, err := db.Find(buildFindQuery(lastResponse.Result[0].Root.Key))
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

func buildCreateQuery() *proto.Query {
	q := query.New()
	s := q.Create()
	n := s.Root()
	n.StringAttr("name", "Tom").IntAttr("age", 50).BoolAttr("hungry", true)
	r := n.Related().Type("MARRIED")
	r.IntAttr("since", time.Now().Unix())
	r.To().StringAttr("name", "Jenny")

	return q.Build()
}

func buildFindQuery(key string) *proto.Statement_Find {
	s := query.NewFind()
	n := s.Root().Key(key)
	n.Related().Type("MARRIED").To()
	return s.Query()
}
