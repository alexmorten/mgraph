package db

import "time"

//Relation ..
type Relation struct {
	Key
	WrittenAt time.Time
	Type      string
	From      Key
	from      *Node
	To        Key
	to        *Node
}

//NewRelation with random key
func NewRelation(from Key, t string, to Key) Relation {
	return Relation{
		Key:       NewKey(),
		WrittenAt: time.Now(),
		Type:      t,
		From:      from,
		To:        to,
	}
}
