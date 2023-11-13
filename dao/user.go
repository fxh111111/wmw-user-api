package dao

import (
	"context"
	"errors"
	"wmw-user-api/model"
	"wmw-user-api/utility"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm/logger"
)

type user struct {
	db         string
	collection string
}

var User = &user{
	db:         "wmw",
	collection: "users",
}

func (u *user) FindByMobile(ctx context.Context, mobile string) (res *model.User, err error) {
	if u == nil {
		return nil, errors.New("dao user no init")
	}
	res = new(model.User)
	err = utility.GetMongoClient().Database(u.db).Collection(u.collection).FindOne(ctx, bson.D{{"mobile", mobile}}).Decode(res)
	if err != nil {
		logger.Default.Error(ctx, "find user by mobile failed", err)
		return nil, err
	}
	return res, nil
}

func (u *user) FindByEmail(ctx context.Context, email string) (res *model.User, err error) {
	if u == nil {
		return nil, errors.New("dao user no init")
	}
	res = new(model.User)
	err = utility.GetMongoClient().Database(u.db).Collection(u.collection).FindOne(ctx, bson.D{{"email", email}}).Decode(res)
	if err != nil {
		logger.Default.Error(ctx, "find user by email failed", err)
		return nil, err
	}
	return res, nil
}

func (u *user) FindByID(ctx context.Context, id string) (res *model.User, err error) {
	if u == nil || id == "" {
		return nil, errors.New("dao user no init")
	}
	var objID primitive.ObjectID
	objID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res = new(model.User)
	err = utility.GetMongoClient().Database(u.db).Collection(u.collection).FindOne(ctx, bson.D{{"_id", objID}}).Decode(res)
	if err != nil {
		logger.Default.Error(ctx, "find user by id failed", err)
		return nil, err
	}
	return res, nil
}

func (u *user) Add(ctx context.Context, in *model.UserRegisterReq) (id string, err error) {
	if u == nil || in == nil {
		return "", errors.New("dao user no init")
	}
	var res *mongo.InsertOneResult
	res, err = utility.GetMongoClient().Database(u.db).Collection(u.collection).InsertOne(ctx, in)
	if err != nil {
		logger.Default.Error(ctx, "insert user failed", err)
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
