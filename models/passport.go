package models

import (
	"fmt"

	"example.com/digital-passport/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Passport struct {
	CompanyId    string
	PassportId   string
	PassportName string
	Files        []string
	Locked       bool
}

func GetPassportsForCompany(companyId string) ([]Passport, error) {
	var passports = []Passport{}
	filter := bson.D{{Key: "companyId", Value: companyId}}
	results, err := db.DB.Collection("passports").Find(db.DBctx, filter)
	if err != nil {
		fmt.Println("err: ", err)
		return []Passport{}, err
	}

	for results.Next(db.DBctx) {
		var p Passport
		err := results.Decode(&p)
		if err != nil {
			fmt.Println("err: ", err)
		}
		passports = append(passports, p)
	}

	return passports, nil
}

func (p Passport) Save() (string, error) {
	res, err := db.DB.Collection("passports").InsertOne(db.DBctx, p)
	if err != nil {
		return "", err
	}
	fmt.Println("Inserted a single document from a struct: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
