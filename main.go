package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
func userChoice() char {
	fmt.Println("Enter command: ")
	var userC char
	fmt.Scanf(&userC)
	return userC
}
*/

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

/*
func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {
    collection := client.Database(dataBase).Collection(col)
    result, err := collection.InsertMany(ctx, docs)
    return result, err
}
*/

func main() {
	client, ctx, cancel, err := connect("mongodb+srv://admin:mtmE823tv12@cluster0.euonptj.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

	ping(client, ctx)

	route()

	/*

		var document interface{}
		var usrN string
		var passwd string
		var tstCrd string
		usrN = "test"
		passwd = "test"
		tstCrd = "test"

		document = bson.M {"username": usrN, "password": passwd, "cards": bson.M {"testCard": tstCrd}}

		insertOneResult, err := insertOne(client, ctx, "CardApp", "Users", document)

		if err != nil {
			panic(err)
		}

		fmt.Println("Result of InsertOne")
		fmt.Println(insertOneResult.InsertedID)
	*/
}
