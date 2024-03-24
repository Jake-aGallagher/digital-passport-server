package models

import (
	"errors"
	"fmt"

	"example.com/digital-passport/db"
	"example.com/digital-passport/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
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

func (u *User) ValidateCredentials() (string, string, error) {
	var result User
	err := db.DB.Collection("users").FindOne(db.DBctx, bson.M{"username": u.Username}).Decode(&result)
	if err != nil {
		return "", "", errors.New("credentials invalid")
	}

	passwordValid := utils.CheckPasswordHash(result.Password, u.Password)
	if !passwordValid {
		return "", "", errors.New("credentials invalid")
	}
	return result.Company, result.ID.Hex(), nil // needs to get user id not username
}
