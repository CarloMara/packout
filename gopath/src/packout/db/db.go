package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PackoutDB struct {
	Client *mongo.Client
	Contex context.Context
}

func Init() *PackoutDB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Print(err)
		panic(fmt.Sprintf("panic in mongo"))
	}
	ctx, _ := context.WithTimeout(context.Background(), 300*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Print(err)
		panic(fmt.Sprintf("panic in mongo connect"))
	}
	//fmt.Print("mongo inizialized without errors")
	return &PackoutDB{Client: client, Contex: ctx}
}
