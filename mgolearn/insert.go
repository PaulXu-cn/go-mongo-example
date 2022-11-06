package mgolearn

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type People struct {
    Id        bson.ObjectId `bson:"_id"`
    Name      string        `bson:"name"`
    Age       int32         `bson:"age"`
    Man       bool          `bson:"man"`
    Skills    []string      `bson:"skills"`
    CreatedAt time.Time     `bson:"created_at"`
}

func TestInsert() {
    var err error
    session, err := mgo.DialWithTimeout(mongoDsn, 10*time.Second)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    // select db、database
    dbConn := session.DB("test").C("people")

    var one = People{
        Id: bson.NewObjectId(), // 注意，这里和官方库有所有不同，需要手动生成
        Name: "Lee",
        Age: 32,
        Man: true,
        Skills: []string{"cook","drive"},
        CreatedAt: time.Now()}
    err = dbConn.Insert(one)
    if nil != err {
        fmt.Printf("mgo insert people err %v\n", err)
    } else {
        fmt.Printf("mgo insert people success~ \n")
    }
}


func TestInsertMany() {
    var err error
    session, err := mgo.DialWithTimeout(mongoDsn, 10*time.Second)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    // select db、database
    dbConn := session.DB("test").C("people")

    var mans = []People{
        {
            Id:        bson.NewObjectId(), // 注意，这里和官方库有所有不同，需要手动生成
            Name:      "Joe",
            Age:       29,
            Man:       true,
            Skills:    []string{"cook", "kung fu"},
            CreatedAt: time.Now()},
        {
            Id:        bson.NewObjectId(),
            Name:      "David",
            Age:       37,
            Man:       false,
            Skills:    []string{"nursing", "drive","IT"},
            CreatedAt: time.Now()},
    }
    for _, row := range mans {
        err = dbConn.Insert(row)
        if nil != err {
            fmt.Printf("mgo batch insert people err %v\n", err)
        }
    }

    // 下面是错误示范
    err = dbConn.Insert(mans)   // ❌ 这样会报错！
    if nil != err {
        fmt.Printf("mgo batch insert people err %v\n", err)
    } else {
        fmt.Printf("mgo batch insert people success~ \n")
    }
}