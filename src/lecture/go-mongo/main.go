package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string `bson:"name"`
	Age  string `bson:"age"`
	Pnum string `bson:"pnum,omitempty"` // 변경
}

func main() {
	exam8()
}

func exam3() {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))

	//db -> collection 단계적 접근
	db := client.Database("go-ready") // database 접속
	col := db.Collection("tPerson")   // collection 접속
	filter := bson.M{"name": bson.M{"$eq": "inTest3"}}
	update := bson.M{
		"$uset": bson.M{
			"age": "",
		},
	}
	result, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document update with ID: %s\n", result.UpsertedID)

	newPerson2 := Person{Name: "inTest4", Age: "45"}
	result2, err := col.InsertOne(context.TODO(), newPerson2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with ID: %s\n", result2.InsertedID)
}
func exam8() {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))

	//db -> collection 단계적 접근
	db := client.Database("go-ready") // database 접속
	col := db.Collection("tPerson")   // collection 접속

	strName := "codz"
	filter := bson.D{{"name", strName}}

	//전체 data count 조회
	estCount, estCountErr := col.EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		panic(estCountErr)
	}

	fmt.Println("Total Document Count", estCount)
	// filter 조건에 맞는 count 조회
	count, err := col.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Filter Document Count", count)
}
func exam7() {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))

	//db -> collection 단계적 접근
	db := client.Database("go-ready") // database 접속
	col := db.Collection("tPerson")   // collection 접속
	filter := bson.D{{"name", "inTest3"}}
	result, err := col.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	var result2 Person
	err = col.FindOne(context.TODO(), filter).Decode(&result2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result2)
}
func exam6() {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))

	//db -> collection 단계적 접근
	db := client.Database("go-ready") // database 접속
	col := db.Collection("tPerson")   // collection 접속
	filter := bson.M{"name": bson.M{"$eq": "inTest3"}}
	update := bson.M{
		"$set": bson.M{
			"age": "50",
		},
	}
	result, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document update with ID: %s\n", result.UpsertedID)
	var result2 Person
	err = col.FindOne(context.TODO(), filter).Decode(&result2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result2)
}
func exam5() {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))

	//db -> collection 단계적 접근
	db := client.Database("go-ready") // database 접속
	col := db.Collection("tPerson")   // collection 접속

	filter := bson.D{}

	cur, err := col.Find(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		return
	} else if err != nil {
		panic(err)
	}
	var results []Person
	for cur.Next(context.TODO()) {
		var elem Person
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", results)
}
