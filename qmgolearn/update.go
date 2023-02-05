package qmgolearn

import (
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"fmt"
	"time"
)

func TestQMUpdate() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn})

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	db := client.Database("test")
	coll := db.Collection("status")

	theId, _ := primitive.ObjectIDFromHex("6364cc67de9ae6b332b797dc")
	var filter = bson.M{"_id": theId.Hex()}
	fmt.Printf("test filter %+v \n", filter)
	//filter = bson.M{"_id": ObjectIdHex("6364cc67de9ae6b332b797dc")}
	var theTime = time.Now()
	var setData = bson.M{"$set": bson.M{"owner": "heihei", "mark": 2333, "timestamp": &theTime, "updated_at": &theTime}}
	err = coll.UpdateOne(ctx, filter, setData)
	fmt.Printf("qmgo update one status, filter %+v,  err %v \n", filter, err)

	var filter2 = bson.M{"device_id": "5823df872326efb5de188038", "owner": "gopher"}
	var setData2 = bson.M{"$set": bson.M{"mark": 2333, "timestamp": &theTime, "updated_at": &theTime}}
	err = coll.UpdateOne(ctx, filter2, setData2)
	fmt.Printf("qmgo update one status 2, err %v \n", err)
}

func TestQMUpdateById() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn})

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	db := client.Database("test")
	coll := db.Collection("status")

	theId, _ := primitive.ObjectIDFromHex("6364cc67de9ae6b332b797dc")
	var theTime = time.Now()
	var setData = bson.M{"$set": bson.M{"owner": "heihei", "mark": 2333, "timestamp": &theTime, "updated_at": &theTime}}
	err = coll.UpdateId(ctx, theId, setData)
	fmt.Printf("qmgo update one by id status, id %+v,  err %v \n", theId, err)
}

func TestQMUpdateMany() {
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
	var theTime = time.Now()
	var setData = bson.M{"$set": bson.M{"owner": "heihei", "mark": 2333, "timestamp": &theTime, "updated_at": &theTime}}
	// update many
	aff, err := coll.UpdateAll(ctx, filter, setData)
	fmt.Printf("qmgo update status, affect %+v,  err %v \n", aff.UpsertedCount, err)

}
