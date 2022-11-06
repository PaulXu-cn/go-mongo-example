package mgolearn

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

func TestUpdateById() {
    var err error
    session, err := mgo.DialWithTimeout(mongoDsn, 10*time.Second)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    // select db、database
    dbConn := session.DB("test").C("people")

    var theId = bson.ObjectIdHex("6367420673abe5b71c45181a")

    var newData = bson.M{"$set": bson.M{"age": 34}}
    err = dbConn.UpdateId(theId, newData)
    if nil != err {
        fmt.Printf("mgo update one people by Id err %v\n", err)
    } else {
        fmt.Printf("mgo update one people by Id success~ \n")
    }

    var newOne = People{Age: 34}    // ❌ 这种写法会报错
    err = dbConn.UpdateId(theId, newOne)
    if nil != err {
        fmt.Printf("mgo update one people by Id err %v\n", err)
    } else {
        fmt.Printf("mgo update one people by Id success~ \n")
    }
}


func TestUpdateOne() {
    var err error
    session, err := mgo.DialWithTimeout(mongoDsn, 10*time.Second)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    // select db、database
    dbConn := session.DB("test").C("people")

    var filter = bson.M{"name": "David"}
    var newData = bson.M{"$set": bson.M{"age": 34}}
    err = dbConn.Update(filter, newData)
    if nil != err {
        fmt.Printf("mgo update one people err %v\n", err)
    } else {
        fmt.Printf("mgo update one people success~ \n")
    }
}

func TestUpdateMany() {
    var err error
    session, err := mgo.DialWithTimeout(mongoDsn, 10*time.Second)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    // select db、database
    dbConn := session.DB("test").C("people")

    var filter = bson.M{"name": "Joe"}
    var newData = bson.M{"$set": bson.M{"man": false}}
    changed, err := dbConn.UpdateAll(filter, newData)
    if nil != err {
        fmt.Printf("mgo update people err %v\n", err)
    } else {
        fmt.Printf("mgo update people result %+v \n", changed)
    }
}
