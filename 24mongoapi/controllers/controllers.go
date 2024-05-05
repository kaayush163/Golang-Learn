package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://kaayush163:aayush@cluster0.q5dfpmp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const colName = "watchlist"

// ctrl+shift+P
// Most Important
var collection *mongo.Collection

//Connect with mongodb

func init() {
	//databse connection

	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption) //TODO is a type of context methods

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo DB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("colelction instance is ready ")
}

// 26

// Check their documentation:

// context.Background():

// func Background() Context
// Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

// context.TODO():

// func TODO() Context
// TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).

//MONGODB Helpers

// insert 1 record
func insertOneMovie(movie model.Netflix) { //not depend on folder name but the package model coming from model folder

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("insertted one movie in db with id:", inserted.InsertedID)
}

//update 1 record

func updateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)
}

// delete one record
func deleteOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete count: ", deleteCount)

}

// delete all record s from mongodb
func deleteAllMovie() {
	// filter := bson.D{{}} //seniors developers dont like in different line to write instead do in collection.DeleteMany(context,bson.D({})) directly

	// collection.DeleteMany(context.Background(), filter)
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Number of  ovies delete:,", deleteResult.DeletedCount)

	return deleteResult.DeletedCount

}
