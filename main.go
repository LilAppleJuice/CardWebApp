package main

import (
	"fmt"
	//"net/http"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	//"home/ubuntu/Documents/CS/CardWebApp/cards"

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

func ping(client *mongo.Client, ctx context.Context) error{
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
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	//})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	//http.ListenAndServe(":80", nil)

	client, ctx, cancel, err := connect("mongodb+srv://admin:mtmE823tv12@cluster0.euonptj.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)

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

	//ping(client, ctx)
	/*

	var userC char
	var isDone bool
	isDone := false
	while isDone == false {
		userChoice(&userC)
		if userC == 'h' {
			fmt.Println("\nh for help\nc to make a new card\nl to list all cards")
		} else if userC == 'c' {
			isDone := true
			cards.MakeCard("front of card", "back of card", true, false)
		} else if userC == 'l' {
			isDone := true
			cards.ListCards()
		} else {
			fmt.Println("Invalid Command Entered, Try again: ")
		}
	}
	fmt.Println("Closing Program...")

	userChoice()
	*/
}
