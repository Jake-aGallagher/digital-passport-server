package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

var DBctx context.Context

type Company struct {
	Name string
}

func InitDB() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	database := os.Getenv("DATABASE")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	// Create a context to use with the connection
	DBctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Call the cancel function when we are done with the context

	// Create the URI we will use to connect to our cosmosDB
	connecturi := fmt.Sprintf(
		"mongodb://%s:%s@%s.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&replicaSet=globaldb&maxIdleTimeMS=120000&appName=@digital-passport-db@",
		username,
		password,
		database)

	// Connect to the DB
	client, err := mongo.Connect(DBctx, options.Client().ApplyURI(connecturi))

	//Check for any errors
	if err != nil {
		panic("could not connect to db")
	}

	//Ping the DB to confirm the connection
	err = client.Ping(DBctx, nil)

	//Check for any errors
	if err != nil {
		panic("could not confirm connection")
	}

	//Print confirmation of connection
	fmt.Println("Connected to MongoDB!")

	DB = client.Database(database)
}
