package model

import "gopkg.in/mgo.v2/bson"

type Admin struct {
	ID         bson.ObjectId `bson:"_id"`
	Username   string        `bson:"username" json:"username"`
	ApprovedBy string        `bson:"approved_by" json:"approved_by"`
}
