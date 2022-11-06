package mgolearn

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

func TestDeleteById() {
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

    err = dbConn.RemoveId(theId)
    if nil != err {
        fmt.Printf("mgo delete one people by Id err %v\n", err)
    } else {
        fmt.Printf("mgo delete one people by Id success~ \n")
    }
}


func TestDeleteOne() {
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
    err = dbConn.Remove(filter)
    if nil != err {
        fmt.Printf("mgo delete one people err %v\n", err)
    } else {
        fmt.Printf("mgo delete one people success~ \n")
    }
}


func TestDeleteMany() {
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
    changed, err := dbConn.RemoveAll(filter)
    if nil != err {
        fmt.Printf("mgo delete people err %v\n", err)
    } else {
        fmt.Printf("mgo delete people result %+v \n", changed)
    }
}
