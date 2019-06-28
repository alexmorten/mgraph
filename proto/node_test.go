package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Matches(t *testing.T) {
	t.Run("Empty nodes match", func(t *testing.T) {
		qn1 := &QueryNode{}
		qn2 := &QueryNode{}
		assert.True(t, qn1.Matches(qn2))
	})

	t.Run("Nodes should match themsselves", func(t *testing.T) {
		qns := []*QueryNode{&QueryNode{}, &QueryNode{Type: "some_type"}, &QueryNode{Attributes: map[string]*AttributeValue{"some_attribute_string": &AttributeValue{Value: &AttributeValue_StringValue{}}}}}
		for _, qn := range qns {
			assert.True(t, qn.Matches(qn))
		}
	})

	t.Run("A valid overlap matches", func(t *testing.T) {
		qn := &QueryNode{Attributes: map[string]*AttributeValue{
			"some_attribute_string": &AttributeValue{Value: &AttributeValue_StringValue{StringValue: "some string"}},
		}}

		qn2 := &QueryNode{Attributes: map[string]*AttributeValue{
			"some_attribute_string": &AttributeValue{Value: &AttributeValue_StringValue{StringValue: "some string"}},
			"some_attribute_int":    &AttributeValue{Value: &AttributeValue_IntValue{IntValue: 12345678}},
		}}
		assert.True(t, qn.Matches(qn2))
	})

	t.Run("A no overlap doesn't match", func(t *testing.T) {
		qn := &QueryNode{Attributes: map[string]*AttributeValue{
			"some_attribute_string": &AttributeValue{Value: &AttributeValue_StringValue{StringValue: "some string"}},
			"some_attribute_int":    &AttributeValue{Value: &AttributeValue_IntValue{IntValue: 12345678}},
		}}
		qn2 := &QueryNode{Attributes: map[string]*AttributeValue{"some_other_string": &AttributeValue{Value: &AttributeValue_StringValue{StringValue: "some other string"}}}}
		assert.False(t, qn.Matches(qn2))
	})
}
