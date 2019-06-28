package proto

import fmt "fmt"

//AttributesMatch checks if all values in a exist in a2 and have the same value
func AttributesMatch(a, a2 map[string]*AttributeValue) bool {

	for key, value := range a {
		v := value.GetValue()
		switch v.(type) {
		case *AttributeValue_StringValue:
			v2 := a2[key].GetStringValue()
			if value.GetStringValue() != v2 {
				return false
			}

		case *AttributeValue_IntValue:
			v2 := a2[key].GetIntValue()
			if value.GetIntValue() != v2 {
				return false
			}

		case *AttributeValue_FloatValue:
			v2 := a2[key].GetFloatValue()
			if value.GetFloatValue() != v2 {
				return false
			}

		case *AttributeValue_DoubleValue:
			v2 := a2[key].GetDoubleValue()
			if value.GetDoubleValue() != v2 {
				return false
			}

		default:
			panic(fmt.Errorf("type %s not handled in Matches()", v))
		}
	}

	return true
}
