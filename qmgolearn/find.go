package qmgolearn

import (
    "context"
    "fmt"
    "time"

    "github.com/qiniu/qmgo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func TestFindOne() {
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

    var filter = bson.M{"age": 10}
    var student = Student{}
    if err := cli.Find(ctx, filter).One(&student); nil != err {
        if mongo.ErrNoDocuments != err {
            fmt.Printf("qmgo find one student err %v \n", err)
        }
    } else {
        fmt.Printf("qmgo find one student result %+v, err %v \n", student, err)
    }
}

func TestFindMany() {
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

    var filter = bson.M{"age": 9}
    var students = []Student{}
    if err := cli.Find(ctx, filter).All(&students); nil != err {
        fmt.Printf("qmgo find student err %v \n", err)
    } else {
        fmt.Printf("qmgo find student result %+v, err %v \n", students, err)
    }
}
