package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodb-with-golang/models"
	"net/http"

	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.15")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb connected.......")

	collection = client.Database("mylistdb").Collection("mycollection")
	fmt.Println("db.......")

}


func Task(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("method right")

	var list models.List
	err := json.NewDecoder(r.Body).Decode(&list)

	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
	}
	fmt.Println("done with decoding")

	result, err := collection.InsertOne(context.TODO(), list)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done with insertion")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
	fmt.Println(result.InsertedID)

}

func Task1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	fmt.Println("method right")

	var list models.List
	err := json.NewDecoder(r.Body).Decode(&list)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	// fmt.Println(res)
	result, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer result.Close(context.TODO())
	var results []bson.M
	err = result.All(context.TODO(), &results)
	if err != nil {
		http.Error(w, "Error processing datbase results", http.StatusInternalServerError)
		log.Println("Error processing results", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Request Successful")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
	}

	fmt.Println(results)

}

func Task2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control_Allow_Methods", "DELETE")

	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	result, err := collection.DeleteOne(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount)
	// fmt.Println(result.DeletedID)

}

func Task3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH")
	if r.Method != http.MethodPatch {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Println("method right")
	var requestBody struct {
		Filter map[string]interface{} `json:"filter"`
		Update map[string]interface{} `json:"update"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Println("decoded")
	filter := requestBody.Filter
	update := map[string]interface{}{
		"$set": requestBody.Update}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated")
	response := map[string]interface{}{
		"modifiedCount": result.ModifiedCount,
	}
	fmt.Println("modified")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

