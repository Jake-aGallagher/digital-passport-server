package models

import (
	"fmt"

	"example.com/digital-passport/db"
	"example.com/digital-passport/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username string
	Email    string
	Password string
	Company  string
}

func (u User) Save() (string, error) {
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return "", err
	}
	u.Password = hashed
	res, err := db.DB.Collection("users").InsertOne(db.DBctx, u)
	if err != nil {
		return "", err
	}
	fmt.Println("Inserted a single document from a struct: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
