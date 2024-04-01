package models

import (
	"fmt"
	"time"

	"example.com/digital-passport/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Passport struct {
	PassportId   primitive.ObjectID `bson:"_id,omitempty" json:"passportId"`
	CompanyId    string             `bson:"company_id" json:"companyId"`
	Created      time.Time          `bson:"created" json:"created"`
	PassportName string             `bson:"passport_name" json:"passportName"`
	Files        []string           `bson:"files" json:"files"`
	Locked       bool               `bson:"locked" json:"locked"`
	UseCode      string             `bson:"use_code" json:"useCode"`
	LinkedArr    []string           `bson:"linked_arr" json:"linkedArr"`
}

func GetPassportsForCompany(companyId string) ([]Passport, error) {
	var passports = []Passport{}
	filter := bson.D{{Key: "company_id", Value: companyId}}
	find := options.Find()
	find.SetSort(bson.M{"created": -1})
	results, err := db.DB.Collection("passports").Find(db.DBctx, filter, find)
	if err != nil {
		fmt.Println("err: ", err)
		return []Passport{}, err
	}
	//utils.LogMongo(results)

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

func GetPassportById(passportId string) (Passport, error) {
	var passport Passport
	objId, err := primitive.ObjectIDFromHex(passportId)
	if err != nil {
		return passport, err
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	err = db.DB.Collection("passports").FindOne(db.DBctx, filter).Decode(&passport)
	return passport, err
}

func GetLinkedMap(linkedArr []string) (map[string]string, error) {
	filter := bson.D{{Key: "use_code", Value: bson.D{{Key: "$in", Value: linkedArr}}}}
	find := options.Find()
	find.SetSort(bson.M{"created": -1})
	listMap := make(map[string]string)
	results, err := db.DB.Collection("passports").Find(db.DBctx, filter, find)
	if err != nil {
		fmt.Println("err: ", err)
		return listMap, err
	}
	defer results.Close(db.DBctx)

	for results.Next(db.DBctx) {
		var p Passport
		err := results.Decode(&p)
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		listMap[p.PassportId.Hex()] = p.PassportName
	}

	if err := results.Err(); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	return listMap, err
}

func (p Passport) Save() (string, error) {
	if p.PassportId == primitive.NilObjectID {
		p.PassportId = primitive.NewObjectID()
	}
	filter := bson.D{{Key: "_id", Value: p.PassportId}}
	update := bson.D{{Key: "$set", Value: p}}
	opts := options.Update().SetUpsert(true)
	result, err := db.DB.Collection("passports").UpdateOne(db.DBctx, filter, update, opts)
	if err != nil {
		fmt.Println("some err: ", err)
		return "", err
	}
	var res string
	if oldId, ok := result.UpsertedID.(primitive.ObjectID); ok {
		res = oldId.Hex()
	} else {
		res = p.PassportId.Hex()
	}
	return res, nil
}
