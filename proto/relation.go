package proto

//ConstructRelation from QueryRelation
func ConstructRelation(qr *QueryRelation) *Relation {
	r := &Relation{
		Key:        qr.Key,
		Type:       qr.Type,
		Attributes: qr.Attributes,
	}

	if r.Key == "" {
		r.Key = newKey()
	}

	return r
}
