package api

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User struct use for api
type User struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Email string        `json:"email" bson:"email"`
	Time  time.Time     `json:"time" bson:"time"`
}

// UserForUpdate struct use for update api
type UserForUpdate struct {
	Name  string    `json:"name" bson:"name"`
	Email string    `json:"email" bson:"email"`
	Time  time.Time `json:"time" bson:"time"`
}

// RspMsg struct use for return error message
type RspMsg struct {
	Msg string `json:"message"`
}

// Dbsession define global var for all func
var Dbsession *mgo.Session
