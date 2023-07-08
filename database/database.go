package database

import (
	"context"
	"log"
	"time"

	"github.com/jeffwilkey/watchlist-go/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client 	 *mongo.Client
	Db		 *mongo.Database
}

var Mongo MongoInstance

func ConnectDb() {
	dbName := config.DBName
	mongoURI := config.MongoURI

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to MongoDB! 🎉")

	Mongo = MongoInstance{
		Client: client,
		Db: db,
	}
}