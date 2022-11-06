package mongolearn

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
    mongoDsn = "mongodb://admin:123456@127.0.0.1:27017"
)

func TestConn() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))
    if nil != err {
        fmt.Printf("mongo connect err %v\n", err)
    } else {
        fmt.Printf("mongo connect success~\n")
    }
    defer func() {
        if err = client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    err = client.Ping(ctx, readpref.Primary())
    if nil != err {
        fmt.Printf("mongo ping err %v\n", err)
    } else {
        fmt.Printf("mongo ping success~\n")
    }
}

func TestConnUseDb() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))
    if nil != err {
        fmt.Printf("mongo connect err %v\n", err)
    } else {
        fmt.Printf("mongo connect success~\n")
        defer func() {
            if err = client.Disconnect(ctx); err != nil {
                panic(err)
            }
        }()
    }

    // 列出DB
    if dbs, err := client.ListDatabaseNames(ctx, bson.M{}, options.ListDatabases()); nil != err {
        fmt.Printf("mongo list dbs err%v", err)
    } else {
        fmt.Printf("mongo dbs %v", dbs)
    }
    
    db := client.Database("test") // 选择 DB

    // 列出 collection
    if colls, err := db.ListCollectionNames(ctx, bson.M{}, options.ListCollections()); nil != err {
        fmt.Printf("mongo list collection err%v", err)
    } else {
        fmt.Printf("mongo db's collection %v", colls)
    }

    coll := db.Collection("test") // 选择 collection
    collName := coll.Name()
    fmt.Printf("mongo collection name is %v\n", collName)
}
