package mongolearn

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/event"
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

    // test ping
    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    err = client.Ping(ctx, readpref.Primary())
    if nil != err {
        fmt.Printf("mongo ping err %v\n", err)
    } else {
        fmt.Printf("mongo ping success~\n")
    }
}

// TestConnUseDb connect, list database, list collection
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

// TestConnUseDb sql monitor
func TestConnWithMonitor() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var clientOpt = options.Client().ApplyURI(mongoDsn)
    var logMonitor = event.CommandMonitor{
        Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
            log.Printf("mongo reqId:%d start on db:%s cmd:%s sql:%+v", startedEvent.RequestID, startedEvent.DatabaseName,
                startedEvent.CommandName, startedEvent.Command)
        },
        Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
            log.Printf("mongo reqId:%d exec cmd:%s success duration %d ns", succeededEvent.RequestID,
                succeededEvent.CommandName, succeededEvent.DurationNanos)
        },
        Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {
            log.Printf("mongo reqId:%d exec cmd:%s failed duration %d ns", failedEvent.RequestID,
                failedEvent.CommandName, failedEvent.DurationNanos)
        },
    }
    // cmd monitor set
    clientOpt.SetMonitor(&logMonitor)
    client, err := mongo.Connect(ctx, clientOpt)
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

    if re, err := client.Database("test").Collection("test").UpdateOne(ctx, bson.M{"name": "cc"}, bson.M{"$set": bson.M{"age": 12}}); err != nil {
        log.Printf("%v", err)
    } else {
        log.Printf("mongo update one re %+v", re)
    }
}
