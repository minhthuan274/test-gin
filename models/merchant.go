package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionMerchant = "merchants"
)

type MerchantDetailReview struct {
	ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name    string        `json:"name" bson:"name"`
	Address string        `json:"address" bson:"address"`
	Covers  []string      `json:"covers" bson:"covers"`
}
