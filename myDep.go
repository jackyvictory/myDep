package myDep

import "gopkg.in/mgo.v2/bson"

// MyStruct for test
type MyStruct struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
}
