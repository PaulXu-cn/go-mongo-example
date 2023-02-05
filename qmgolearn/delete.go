package qmgolearn

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// qmgo remove one
func TestQmDelete() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn})

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	coll := client.Database("test").Collection("status")

	theId, _ := primitive.ObjectIDFromHex("6364cc67de9ae6b332b797dc")
	var filter = bson.M{"_id": theId.Hex()}
	err = coll.Remove(ctx, filter)
	fmt.Printf("qmgo remove one status, filter %+v,  err %v \n", filter, err)
}

// qmgo remove by id
func TestQmDeleteById() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn})

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	coll := client.Database("test").Collection("status")

	theId, _ := primitive.ObjectIDFromHex("6364cc67de9ae6b332b797dc")
	err = coll.RemoveId(ctx, theId)
	fmt.Printf("qmgo remove one status, filter %+v,  err %v \n", theId, err)
}

// qmgo remove many
func TestQMRemoveAll() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn})

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	coll := client.Database("test").Collection("status")

	var filter = bson.M{"name": "heihei"}
	if delRe, err := coll.RemoveAll(ctx, filter); err != nil {
		fmt.Printf("qmgo remove many err %v \n", err)
	} else {
		fmt.Printf("qmgo remove many result %+v, err %v \n", delRe, err)
	}
}
