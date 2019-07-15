package db_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/alexmorten/mgraph/db"
	"github.com/alexmorten/mgraph/proto"
	"github.com/alexmorten/mgraph/query"
	fuzz "github.com/google/gofuzz"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testModel struct {
	SomeBool    bool
	SomeInt     int
	SomeFloat   float32
	SomeFloat64 float64
	SomeString  string
}

const testDataDir = "test_data"

func Test_Update(t *testing.T) {
	conf := db.DefaultConfig()
	conf.Dir = testDataDir
	defer func() {
		os.RemoveAll(testDataDir)
	}()
	db, err := db.New(conf)
	require.NoError(t, err)

	t.Run("simple insertions work", func(t *testing.T) {
		t.Parallel()
		fuzzer := fuzz.New()
		for i := 0; i < 1000; i++ {
			n := &testModel{}
			fuzzer.Fuzz(n)

			insertQuery := query.New()
			qn := insertQuery.Create().Root()
			assignQueryAttributes(&qn.Attributes, n)

			insertResponse, err := db.Update(insertQuery.Build())
			require.NoError(t, err)
			require.Len(t, insertResponse.Result, 1)
			k := insertResponse.Result[0].Root.Key

			findStatement := query.NewFind()
			findStatement.Root().Key(k)
			findResponse, err := db.Find(findStatement.Query())
			require.NoError(t, err)
			require.Len(t, findResponse.Result, 1)

			foundNode := findResponse.Result[0].Root

			assertAttributesMatch(t, n, foundNode.Attributes)
		}
	})

	t.Run("nested insertions work", func(t *testing.T) {
		t.Parallel()
		fuzzer := fuzz.New()
		for i := 0; i < 100; i++ {
			rootN := &testModel{}
			relation := &testModel{}
			nestedN := &testModel{}
			fuzzer.Fuzz(rootN)
			fuzzer.Fuzz(relation)
			fuzzer.Fuzz(nestedN)
			fmt.Println(relation)

			insertQuery := query.New()
			qn := insertQuery.Create().Root()
			r := qn.Related().Type("nested")
			nestedQn := r.To()

			assignQueryAttributes(&qn.Attributes, rootN)
			assignQueryAttributes(&r.Attributes, relation)
			assignQueryAttributes(&nestedQn.Attributes, nestedN)

			insertResponse, err := db.Update(insertQuery.Build())
			require.NoError(t, err)
			require.Len(t, insertResponse.Result, 1)
			k := insertResponse.Result[0].Root.Key

			findStatement := query.NewFind()
			findStatement.Root().Key(k).Related().Type("nested").To()
			findResponse, err := db.Find(findStatement.Query())
			require.NoError(t, err)
			require.Len(t, findResponse.Result, 1)

			foundRoot := findResponse.Result[0].Root
			assertAttributesMatch(t, rootN, foundRoot.Attributes)
			require.Len(t, foundRoot.Relations, 1)
			foundRelation := foundRoot.Relations[0]
			assertAttributesMatch(t, relation, foundRelation.Attributes)

			foundNested := foundRelation.GetTo()
			require.NotNil(t, foundNested)
			assertAttributesMatch(t, nestedN, foundNested.Attributes)
		}
	})
}

func assignQueryAttributes(qn *query.Attributes, n *testModel) {
	qn.BoolAttr("some_bool", n.SomeBool)
	qn.IntAttr("some_int", int64(n.SomeInt))
	qn.FloatAttr("some_float", n.SomeFloat)
	qn.DoubleAttr("some_double", n.SomeFloat64)
	qn.StringAttr("some_string", n.SomeString)
}

func assertAttributesMatch(t *testing.T, n *testModel, attr map[string]*proto.AttributeValue) {
	assert.Equal(t, n.SomeBool, attr["some_bool"].GetBooleanValue())
	assert.Equal(t, n.SomeInt, int(attr["some_int"].GetIntValue()))
	assert.Equal(t, n.SomeFloat, attr["some_float"].GetFloatValue())
	assert.Equal(t, n.SomeFloat64, attr["some_double"].GetDoubleValue())
	assert.Equal(t, n.SomeString, attr["some_string"].GetStringValue())
}
