package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string             `bson:"firstName" json:"firstName" form:"firstName" binding:"required"`
	LastName  string             `bson:"lastName" json:"lastName" form:"lastName" binding:"required"`
	Mobile    string             `bson:"mobile" json:"mobile" form:"mobile" binding:"required"`
	Email     string             `bson:"email" json:"email" form:"email" binding:"required"`
	Password  string             `bson:"password" json:"-" form:"password" binding:"required"`
}

type UserRegisterReq struct {
	FirstName string `bson:"firstName" json:"firstName" form:"firstName" binding:"required"`
	LastName  string `bson:"lastName" json:"lastName" form:"lastName" binding:"required"`
	Mobile    string `bson:"mobile" json:"mobile" form:"mobile" binding:"required"`
	Email     string `bson:"email" json:"email" form:"email" binding:"required"`
	Password  string `bson:"password" json:"password" form:"password" binding:"required"`
}

type UserLoginReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
