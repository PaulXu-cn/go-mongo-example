package mongolearn

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeleteOne() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    var filter = bson.D{{"name", "Bob2"}}
    delRe, err := coll.DeleteOne(ctx, filter)
    if nil != err {
            fmt.Printf("mongo delete one err %v \n", err)
    } else {
        fmt.Printf("mongo delete one user deleteRe %+v, err %v \n", delRe, err)
    }
}


func TestDeleteMany() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    var filter = bson.D{{"name", "Bob2"}}
    delRe, err := coll.DeleteMany(ctx, filter)
    if nil != err {
        fmt.Printf("mongo batch delete err %v \n", err)
    } else {
        fmt.Printf("mongo batch delete user deleteRe %+v, err %v \n", delRe, err)
    }
}
