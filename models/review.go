package models

import "gopkg.in/mgo.v2/bson"

const (
	CollectionReview = "reviews"
)

type Review struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User     bson.ObjectId `json:"user" bson:"user,omitempty"`
	Merchant bson.ObjectId `json:"merchant" bson:"merchant,omitempty" form:"merchant"`
	Feedback string        `json:"feedback" bson:"feedback" form:"feedback"`
	Point    int           `json:"point" bson:"point" form:"point"`
}

type ReviewJson struct {
	Merchant string `json:"merchant" binding:"required" form:"merchant"`
	Feedback string `json:"feedback" binding:"required" form:"feedback"`
	Point    int    `json:"point" binding:"required" form:"point"`
}
