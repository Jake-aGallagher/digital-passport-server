package utils

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func LogMongo(res *mongo.Cursor) {
	j, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println("results: ", string(j))
}
