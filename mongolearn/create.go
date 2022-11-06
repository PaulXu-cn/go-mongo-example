package mongolearn

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreate() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    var newData = bson.D{{"name", "Tom"}, {"age", 18}, {"weight", 65.3},
        {"studying", true}, {"tag", []string{"student", "man", "outgoing"}},
        {"created_at", time.Now()}}
    res, err := coll.InsertOne(ctx, newData)
    id := res.InsertedID
    fmt.Printf("mongo insert new data %+v insert ID %v, err %v \n", newData, id, err)

}

func TestCreateMulti() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    //var users = []User{
    //    {
    //        "Bob",
    //        19,
    //        71.5,
    //        false,
    //        []string{"man"},
    //        time.Now(),
    //    },
    //    {
    //        "Rose",
    //        17,
    //        53.1,
    //        true,
    //        []string{"woman"},
    //        time.Now(),
    //    },
    //}
    var users = []interface{}{
        bson.D{{"name", "Bob2"}, {"age", 24}, {"weight", float32(63.7)},
            {"studying", false}, {"tag", []string{"freeman", "man", "music"}},
            {"created_at", time.Now()}},
        bson.D{{"name", "Tom2"}, {"age", 28}, {"weight", float32(58.1)},
            {"studying", false}, {"tag", []string{"women", "flower"}},
            {"created_at", time.Now()}},
    }
    insertRe, err := coll.InsertMany(ctx, users)
    if nil != err {
        fmt.Printf("mongo batch insert user err %v \n", err)
    } else {
        fmt.Printf("mongo batch insert user result %+v, err %v \n", insertRe, err)
    }
}
