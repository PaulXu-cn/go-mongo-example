package qmgolearn

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBulk1() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn})

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	coll := client.Database("test").Collection("status")

	blk := coll.Bulk()
	blk.SetOrdered(true) // 批量执行是否排序

	blk.InsertOne(bson.M{}) // 批量操作，所以就没有提供 isnertMany

	blk.UpdateOne(nil, bson.M{})
	blk.UpdateId(primitive.NewObjectID(), bson.M{})
	blk.UpdateAll(nil, bson.M{})

	blk.Upsert(nil, bson.M{})
	blk.UpsertId(primitive.NewObjectID(), bson.M{})
	blk.UpsertOne(nil, bson.M{})

	blk.Remove(bson.M{"_id": primitive.NewObjectID()})
	blk.RemoveId(primitive.NewObjectID())
	blk.RemoveAll(bson.M{"_id": primitive.NewObjectID()})

	if blkRe, err := blk.Run(ctx); err != nil {
		fmt.Printf("qmgo bulk run err %v \n", err)
	} else {
		fmt.Printf("qmgo bulk run result %+v, err %v \n", blkRe, err)
	}
}
