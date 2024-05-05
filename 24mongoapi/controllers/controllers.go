package controller

//first letter is always smaller of whatever we havve created whether controller or db
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/kaayush163/Golang-Learn/tree/main/24mongoapi/model"

	"github.com/gorilla/mux"
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
func deleteAllMovie() int64 {
	// filter := bson.D{{}} //seniors developers dont like in different line to write instead do in collection.DeleteMany(context,bson.D({})) directly

	// collection.DeleteMany(context.Background(), filter)
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Number of  ovies delete:,", deleteResult.DeletedCount)

	return deleteResult.DeletedCount

}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M

		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Actusl controller = file
// get all movies from databsse in golang
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) { //need to inserted into router so important to mention formal parameter w and r

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie) //like we do in ajvavsvripty res.json()

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r) //from request body we get every params easily
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()

	json.NewEncoder(w).Encode(count)
}
