package query

import "github.com/alexmorten/mgraph/proto"

//Attributes is helper type for handling proto.Attributes in the Node and relation
type Attributes map[string]*proto.AttributeValue

//StringAttr defines a string attribute on the Attributes
func (a Attributes) StringAttr(k, v string) Attributes {
	a[k] = &proto.AttributeValue{
		Value: &proto.AttributeValue_StringValue{
			StringValue: v,
		},
	}
	return a
}

//IntAttr defines an int attribute on the Attributes
func (a Attributes) IntAttr(k string, v int64) Attributes {
	a[k] = &proto.AttributeValue{
		Value: &proto.AttributeValue_IntValue{
			IntValue: v,
		},
	}
	return a
}

//FloatAttr defines a float attribute on the Attributes
func (a Attributes) FloatAttr(k string, v float32) Attributes {
	a[k] = &proto.AttributeValue{
		Value: &proto.AttributeValue_FloatValue{
			FloatValue: v,
		},
	}
	return a
}

//DoubleAttr defines a double attribute on the Attributes
func (a Attributes) DoubleAttr(k string, v float64) Attributes {
	a[k] = &proto.AttributeValue{
		Value: &proto.AttributeValue_DoubleValue{
			DoubleValue: v,
		},
	}
	return a
}

//BoolAttr defines a bool attribute on the Attributes
func (a Attributes) BoolAttr(k string, v bool) Attributes {
	a[k] = &proto.AttributeValue{
		Value: &proto.AttributeValue_BooleanValue{
			BooleanValue: v,
		},
	}
	return a
}
