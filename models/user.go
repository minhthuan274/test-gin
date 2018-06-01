package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Avatar    string        `json:"avatar" bson:"avatar" `
	Statistic Statistic     `json:"statistic" bson:"statistic"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
}

type Statistic struct {
	Checkin       int32     `json:"checkin" bson:"checkin"`
	Bill          int32     `json:"bill" bson:"bill"`
	Coin          int32     `json:"coin" bson:"coin"`
	CheckinCoin   int32     `json:"checkinCoin" bson:"checkinCoin"`
	BilledZcoin   int32     `json:"billedZcoin" bson:"billedZcoin"`
	Expense       int32     `json:"expense" bson:"expense"`
	Reward        int32     `json:"reward" bson:"reward"`
	Badge         int32     `json:"badge" bson:"badge"`
	LastCheckinAt time.Time `json:"lastCheckinAt" bson:"lastCheckinAt"`
	LastBillAt    time.Time `json:"lastBillAt" bson:"lastBillAt"`
}
