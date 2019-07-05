package db_test

import (
	"os"
	"testing"

	"github.com/alexmorten/mgraph/db"
	"github.com/alexmorten/mgraph/query"
	fuzz "github.com/google/gofuzz"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testNode struct {
	someBool    bool
	someInt     int
	someFloat   float32
	someFloat64 float64
	someString  string
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
		fuzzer := fuzz.New()
		for i := 0; i < 1000; i++ {
			n := &testNode{}
			fuzzer.Fuzz(n)

			insertQuery := query.New()
			qn := insertQuery.Create().Root()

			qn.BoolAttr("some_bool", n.someBool)
			qn.IntAttr("some_int", int64(n.someInt))
			qn.FloatAttr("some_float", n.someFloat)
			qn.DoubleAttr("some_double", n.someFloat64)
			qn.StringAttr("some_string", n.someString)

			insertResponse, err := db.Update(insertQuery.Build())
			require.NoError(t, err)
			require.Len(t, insertResponse.Result, 1)
			k := insertResponse.Result[0].Root.Key

			findStatement := query.NewFind()
			findStatement.Root().Key(k)
			findResponse, err := db.Find(findStatement.Query())
			require.NoError(t, err)
			require.Len(t, insertResponse.Result, 1)

			foundNode := findResponse.Result[0].Root

			assert.Equal(t, n.someBool, foundNode.Attributes["some_bool"].GetBooleanValue())
			assert.Equal(t, n.someInt, int(foundNode.Attributes["some_int"].GetIntValue()))
			assert.Equal(t, n.someFloat, foundNode.Attributes["some_float"].GetFloatValue())
			assert.Equal(t, n.someFloat64, foundNode.Attributes["some_double"].GetDoubleValue())
			assert.Equal(t, n.someString, foundNode.Attributes["some_string"].GetStringValue())

		}
	})
}
