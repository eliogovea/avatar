package model

import "gopkg.in/mgo.v2/bson"

type Avatar struct {
	ID         bson.ObjectId `json:"id" bson:"id"`
	Status     string        `json:"status" bson:"status"` // pending or approved
	Username   string        `json:"username" bson:"username"`
	ApprovedBy string        `json:"approved_by" bson:"approved_by"` // empty if status is  pending
	File       []byte        `json`
}
