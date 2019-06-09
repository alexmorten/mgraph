package chlog

import (
	"time"
)

//Change of the db at a point in time
type Change struct {
	Ts   time.Time
	Data interface{}
}
