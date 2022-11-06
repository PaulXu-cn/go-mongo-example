package qmgolearn

import (
    "context"
    "fmt"
    "time"

    "github.com/qiniu/qmgo"
    "github.com/qiniu/qmgo/field"
)

type Student struct {
    field.DefaultField `bson:",inline"` // 在结构体注入 field.DefaultField, Qmgo 会自动更新 createAt、updateAt and _id 值.

    Name      string    `bson:"name"`
    Age       uint16    `bson:"age"`
    Weight    float64   `bson:"weight"`
    CreatedAt time.Time `bson:"created_at"`
}

func TestInsertOne() {
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

    var new = Student{
        Name:      "xm",
        Age:       7,
        Weight:    41.5,
        CreatedAt: time.Now(),
    }
    if insertRe, err := cli.InsertOne(ctx, &new); nil != err {
        fmt.Printf("qmgo insert one student err %v \n", err)
    } else {
        fmt.Printf("qmgo insert one student result %+v, err %v \n", insertRe, err)
    }
}

func TestInsertMany() {
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

    var students = []Student{
        {
            Name:      "xh",
            Age:       6,
            Weight:    37.8,
            CreatedAt: time.Now(),
        },
        {
            Name:      "hll",
            Age:       9,
            Weight:    44.2,
            CreatedAt: time.Now(),
        },
    }
    if insertRe, err := cli.InsertMany(ctx, students); nil != err {
        fmt.Printf("qmgo batch insert student err %v \n", err)
    } else {
        fmt.Printf("qmgo batch insert student result %+v, err %v \n", insertRe, err)
    }
}
