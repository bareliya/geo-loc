package db

import (
	"context"
	"log"

	//"time"
	"fmt"

	"github.com/citymall/geo-loc/util"

	//"github.com/citymall/geo-loc/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbConnectionType struct {
	MongoClient *mongo.Client
	isconnected bool
}

var MongoDbConnection = &mongodbConnectionType{}

func (db *mongodbConnectionType) ConnectMongodb() {

	if !db.isconnected {
		mongodb := util.GetConfig().Mongodb
		host := mongodb["url"].(string)
		fmt.Println(host)

		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(host))
		if err != nil {
			log.Fatal(err)
		}

		// defer func() {
		// 	if err := client.Disconnect(context.TODO()); err != nil {
		// 		log.Fatal(err)
		// 	}
		// }()

		db.MongoClient = client
		db.isconnected = true
		log.Println("DB Connected!")
	} else {
		// Do Nothing
	}
}

func (db *mongodbConnectionType) CloseConnection() {
	log.Fatal("cleaning up db connection")
	db.MongoClient.Disconnect(context.TODO())
	db.isconnected = false
}

func (db *mongodbConnectionType) InserManyDoc(docs []interface{}) error {
	if !db.isconnected {
		db.ConnectMongodb()
	}
	mongodb := util.GetConfig().Mongodb
	database := mongodb["database"].(string)
	collection := mongodb["collecton"].(string)
	client := db.MongoClient
	coll := client.Database(database).Collection(collection)
	log.Println("Data is being Added in mongoDB .....")
	_, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		//log.Fatal(err)
		return err

	}

	log.Println("successfully added !")
	return err

}
