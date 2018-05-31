package models

import "gopkg.in/mgo.v2/bson"

const (
	CollectionReview = "reviews"
)

type Review struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User     bson.ObjectId `json:"user" bson:"user,omitempty"`
	Merchant bson.ObjectId `json:"merchant" bson:"merchant,omitempty"`
	Feedback string        `json:"feedback" bson:"feedback"`
	Point    int           `json:"point" bson:"point"`
}
