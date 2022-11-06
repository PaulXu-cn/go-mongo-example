package mongolearn

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func TestUpdateOne() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    objId, err := primitive.ObjectIDFromHex("63666c194e6076de7d167a55")
    if nil != err {
        fmt.Printf("mongo gen objID  err %v \n", err)
    }

    var upOpt = options.Update()
    //var newData = bson.M{"$set": User{
    //    Age:    33,
    //    Weight: float32(78.3),
    //}}    // ❌
    //var newData = bson.D{{"$set", bson.D{{"age", 33}, {"weight", float32(78.3)}}}}    // ❌
    var newData = bson.M{"$set": bson.M{"age": 33, "weight": float32(78.3)}}
    upRe, err := coll.UpdateByID(ctx, objId, newData, upOpt)
    if nil != err {
        fmt.Printf("mongo update user byID err %v \n", err)
    } else {
        fmt.Printf("mongo update user byID return %+v err %v\n", upRe, err)
    }

    var filter = bson.D{{"name", "Bob"}}
    upRe, err = coll.UpdateOne(ctx, filter, newData)
    if nil != err {
        fmt.Printf("mongo update one user err %v \n", err)
    } else {
        fmt.Printf("mongo update one user return %+v err %v\n", upRe, err)
    }
}

func TestUpdateMany() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    var filter = bson.D{{"name", "Bob"}}
    var newData = bson.M{"$set": User{
        Age:    33,
        Weight: float32(78.3),
    }}
    upRe, err := coll.UpdateMany(ctx, filter, newData)
    if nil != err {
        fmt.Printf("mongo update maney user err %v \n", err)
    } else {
        fmt.Printf("mongo update maney user return %+v err %v\n", upRe, err)
    }
}
