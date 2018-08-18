package mongodbwrapper

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func NewClient() (*mongo.Client, error) {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	return client, nil
}

//查询所有
func SearchAll() {
	client, err := NewClient()
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("test").Collection("user")
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}
	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(elem)
	}
}
