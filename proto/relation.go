package proto

//ConstructQueryRelation from QueryRelation. Direction will be set according the the key in n
func (r *Relation) ConstructQueryRelation(n *Node) *QueryRelation {
	var direction isQueryRelation_Direction

	if r.To == n.Key {
		direction = &QueryRelation_From{
			From: &QueryNode{Key: r.From},
		}
	} else if r.From == n.Key {
		direction = &QueryRelation_To{
			To: &QueryNode{Key: r.To},
		}
	} else {
		panic("ConstructQueryRelation called with node that is not set as from or to")
	}

	return &QueryRelation{
		Key:        r.Key,
		Type:       r.Type,
		Attributes: r.Attributes,
		Direction:  direction,
	}
}

//OtherSideKey returns the key of the other side of the relation
func (r *Relation) OtherSideKey(key string) string {
	if r.From == key {
		return r.To
	}

	if r.To == r.Key {
		return r.From
	}
	return ""
}

//ConstructRelation from QueryRelation
func (qr *QueryRelation) ConstructRelation() *Relation {
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

//Matches checks if all properties given in n match the properties of n2
func (qr *QueryRelation) Matches(r2 *QueryRelation) bool {
	if qr.Key != "" && qr.Key != r2.Key {
		return false
	}

	if qr.Type != "" && qr.Type != r2.Type {
		return false
	}

	if !AttributesMatch(qr.Attributes, r2.Attributes) {
		return false
	}

	if rFrom, ok := qr.Direction.(*QueryRelation_From); ok {
		if r2From, ok := r2.Direction.(*QueryRelation_From); ok {
			rFrom.From.Matches(r2From.From)
		} else {
			return false
		}
	}

	if rTo, ok := qr.Direction.(*QueryRelation_To); ok {
		if r2To, ok := r2.Direction.(*QueryRelation_To); ok {
			rTo.To.Matches(r2To.To)
		} else {
			return false
		}
	}

	return true
}

//RelatedNode is the QueryNode, that direction is pointing to
func (qr *QueryRelation) RelatedNode() *QueryNode {
	if rTo, ok := qr.Direction.(*QueryRelation_To); ok {
		return rTo.To
	}

	if rFrom, ok := qr.Direction.(*QueryRelation_From); ok {
		return rFrom.From
	}
	return nil
}

//OverwriteRelatedNode with the given QueryNode. If direction is not currently set, this is a no-op
func (qr *QueryRelation) OverwriteRelatedNode(qn *QueryNode) {
	if rTo, ok := qr.Direction.(*QueryRelation_To); ok {
		rTo.To = qn
	}

	if rFrom, ok := qr.Direction.(*QueryRelation_From); ok {
		rFrom.From = qn
	}
}
