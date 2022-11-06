package qmgolearn

import (
    "context"
    "fmt"
    "time"

    "github.com/qiniu/qmgo"
    "gopkg.in/mgo.v2/bson"
)

var (
    mongoDsn   = "mongodb://admin:123456@127.0.0.1:27017"
    db         = "test"
    collection = "test"
)

func TestConn() {
    var timeout int64 = 10 * 1000
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var config = qmgo.Config{
        Uri:              mongoDsn,
        Database:         "admin", // 认证 库, 非后面实际使用的业务库
        ConnectTimeoutMS: &timeout,
    }
    client, err := qmgo.NewClient(ctx, &config)
    if nil != err {
        fmt.Printf("qmgo connect mongo err %v\n", err)
    } else {
        defer func() {
            if err = client.Close(ctx); err != nil {
                panic(err)
            }
        }()
    }

    db := client.Database(db)
    coll := db.Collection(collection)
    collName := coll.GetCollectionName()
    fmt.Printf("qmgo choose collection %v\n", collName)

}

func TestOpen() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: mongoDsn, Database: db, Coll: collection})
    if nil != err {
        fmt.Printf("qmgo connect mongo err %v\n", err)
    } else {
        defer func() {
            if err = cli.Close(ctx); err != nil {
                panic(err)
            }
        }()
    }

    err = cli.Ping(5)
    if nil != err {
        fmt.Printf("qmgo ping err %v\n", err)
    }

    ver := cli.ServerVersion()
    fmt.Printf("qmgo server version %s\n", ver)

    // 在 qmgo 中获取 collection 列表方式
    // 由于无法获取到 mongo.Client 对象，所有无法获取到 database 列表
    mongoColl, err := cli.CloneCollection()
    if nil != err {
        fmt.Printf("qmgo clone and get mongo.collection err %v\n", err)
    }
    if colls, err := mongoColl.Database().ListCollectionNames(ctx, bson.M{}); nil != err {
        fmt.Printf("qmgo list collection err %v\n", err)
    } else {
        fmt.Printf("qmgo list collection %v\n", colls)
    }
}
