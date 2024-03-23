package models

import (
	"fmt"

	"example.com/digital-passport/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Name string
}

func (c Company) Save() (string, error) {
	res, err := db.DB.Collection("companies").InsertOne(db.DBctx, c)
	if err != nil {
		fmt.Println("err: ", err)
		return "", err
	}
	fmt.Println("Inserted a single document from a struct: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
