package qmgolearn

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type People struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Age       uint16             `bson:"age"`
	Weight    float64            `bson:"weight"`
	CreatedAt time.Time          `bson:"created_at"`
}

func TestQmgoUpsert() {
	ctx := context.Background()
	client := GetConn()

	// select db, choose collection
	db := client.Database("test")
	coll := db.Collection("status")

	theId, _ := primitive.ObjectIDFromHex("6364cc67de9ae6b332b797dc")
	var filter = bson.M{"_id": theId}
	fmt.Printf("test filter %+v \n", filter)
	var theTime = time.Now()
	var setData = bson.M{"owner": "heihei", "mark": 2333, "timestamp": &theTime, "updated_at": &theTime}
	upInfo, err := coll.Upsert(ctx, filter, setData)
	fmt.Printf("qmgo upsert status, filter %+v, updateRe %+v err %v \n", filter, upInfo, err)

	filter = bson.M{"name": "ccc"}
	var p = People{
		ID:        primitive.NilObjectID, // 这里Id需是0值
		Name:      "Sam",
		Age:       21,
		Weight:    45,
		CreatedAt: theTime,
	}
	upInfo, err = client.Database("test").Collection("people").Upsert(ctx, filter, p)
	fmt.Printf("qmgo upsert people, filter %+v, updateRe %+v err %v \n", filter, upInfo, err)
}

func TestQmgoUpsertId() {
	ctx := context.Background()
	client := GetConn()

	// select db, choose collection
	db := client.Database("test")
	coll := db.Collection("status")

	theId, _ := primitive.ObjectIDFromHex("6364cc67de9ae6b332b797dc")
	var theTime = time.Now()
	var setData = bson.M{"owner": "heihei", "mark": 2333, "timestamp": &theTime, "updated_at": &theTime}
	err := coll.UpdateId(ctx, theId, setData)
	fmt.Printf("qmgo upsert by id status, id %+v, err %v \n", theId, err)

	var p = People{
		ID:        theId, // 这里可设置为入参的Id值，或者设置为0值
		Name:      "Sam",
		Age:       21,
		Weight:    45,
		CreatedAt: theTime,
	}
	err = client.Database("test").Collection("people").UpdateId(ctx, theId, p)
	fmt.Printf("qmgo upsert by id People, id %+v, err %v \n", theId, err)
}
