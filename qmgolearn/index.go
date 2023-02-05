package qmgolearn

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	opts "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestQMCreateIndex() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn,
		ReadPreference: &qmgo.ReadPref{Mode: readpref.SecondaryPreferredMode},
	})
	if nil != err {
		fmt.Printf("qmgo conn err %v", err)
	}

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	db := client.Database("test")
	coll := db.Collection("behavior")

	var idxs = []options.IndexModel{
		{Key: []string{"device_id", "updated_at"}},
		{Key: []string{"device_id", "timestamp"}},
		{Key: []string{"uuid", "mark"}},
	}
	err = coll.CreateIndexes(ctx, idxs)
	if err != nil {
		fmt.Printf("qmgo create index err %v \n", err)
	}
	fmt.Printf("mgo create index err %v", err)

}

func TestQmgoDelIdx() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn,
		ReadPreference: &qmgo.ReadPref{Mode: readpref.PrimaryPreferredMode},
	})
	if nil != err {
		fmt.Printf("qmgo conn err %v", err)
	}

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	db := client.Database("test")
	coll := db.Collection("behavior")

	var name = "uuid_1_mark_1"
	var name1 = "uuid_1"

	var idxs = []options.IndexModel{
		{Key: []string{"uuid", "mark"}, IndexOptions: &opts.IndexOptions{Name: &name}},
		{Key: []string{"uuid"}, IndexOptions: &opts.IndexOptions{Name: &name1}},
	}
	err = coll.DropIndex(ctx, []string{"uuid"})
	if err != nil {
		fmt.Printf("qmgo remove index err %v \n", err)
	}
	fmt.Printf("mgo remove index re %+v err %v", idxs, err)

	err = coll.DropIndex(ctx, []string{"uuid", "mark"})
	if err != nil {
		fmt.Printf("qmgo remove index err %v \n", err)
	}
	fmt.Printf("mgo remove index re %+v err %v", idxs, err)
}

type IndexShow struct {
	V    int32       `bson:"v"`
	Key  primitive.M `bson:"key"`
	Name string      `bson:"name"`
	Ns   string      `bson:"ns"`
}

// https://stackoverflow.com/questions/68543962/how-to-fetch-and-inspect-mongodb-index-options-using-the-go-mongo-driver
func TestQmgoShowIdx() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoDsn,
		ReadPreference: &qmgo.ReadPref{Mode: readpref.PrimaryPreferredMode},
	})
	if nil != err {
		fmt.Printf("qmgo conn err %v", err)
	}

	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()

	// select db, choose collection
	db := client.Database("test")
	coll := db.Collection("behavior")

	if mgColl, err := coll.CloneCollection(); nil != err {

	} else {
		if cur, err := mgColl.Indexes().List(ctx, opts.ListIndexes()); nil != err {

		} else {

			var result3 = []IndexShow{}
			err = cur.All(ctx, &result3)
			fmt.Printf("qmgo show all index re %+v err %v\n", result3, err)

			var result = []bson.M{}
			err = cur.All(ctx, &result)
			fmt.Printf("qmgo show all index re %+v err %v\n", result, err)

			var result2 = []mongo.IndexView{}
			//var result options.IndexModel
			err = cur.All(ctx, &result2)
			fmt.Printf("qmgo show all index re %+v err %v\n", result2, err)

		}
	}
}
