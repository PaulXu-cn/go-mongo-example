package qmgolearn

import (
    "context"
    "fmt"
    "time"

    "github.com/qiniu/qmgo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func TestTrans() {
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

    callback := func(sessCtx context.Context) (interface{}, error) {

        var filter = bson.M{"age": 10}
        var student = Student{}
        if err := cli.Find(sessCtx, filter).One(&student); nil != err {
            if mongo.ErrNoDocuments != err {
                fmt.Printf("qmgo find one student err %v \n", err)
            }
        } else {
            fmt.Printf("qmgo find one student result %+v, err %v \n", student, err)
        }

        var new = Student{Name: "xm", Age: 7, Weight: 41.5, CreatedAt: time.Now()}
        if insertRe, err := cli.InsertOne(sessCtx, &new); nil != err {
            fmt.Printf("qmgo insert one student err %v \n", err)
        } else {
            fmt.Printf("qmgo insert one student result %+v, err %v \n", insertRe, err)
        }
        return nil, nil
    }

    result, err := cli.DoTransaction(ctx, callback)
    if nil != err {
        fmt.Printf("qmgo do transcation err %v \n", err)
    } else {
        fmt.Printf("qmgo do transcation result %+v, err %v \n", result, err)
    }

}

func TestTrans2() {
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

    s, err := cli.Session()
    defer s.EndSession(ctx)

    callback := func(sessCtx context.Context) (interface{}, error) {

        var filter = bson.M{"age": 10}
        var student = Student{}
        if err := cli.Find(sessCtx, filter).One(&student); nil != err {
            if mongo.ErrNoDocuments != err {
                fmt.Printf("qmgo find one student err %v \n", err)
            }
        } else {
            fmt.Printf("qmgo find one student result %+v, err %v \n", student, err)
        }

        var new = Student{Name: "xm", Age: 7, Weight: 41.5, CreatedAt: time.Now()}
        if insertRe, err := cli.InsertOne(sessCtx, &new); nil != err {
            fmt.Printf("qmgo insert one student err %v \n", err)
        } else {
            fmt.Printf("qmgo insert one student result %+v, err %v \n", insertRe, err)
        }
        return nil, nil
    }

    result, err := s.StartTransaction(ctx, callback)
    if nil != err {
        fmt.Printf("qmgo do transcation err %v \n", err)
    } else {
        fmt.Printf("qmgo do transcation result %+v, err %v \n", result, err)
    }
}
