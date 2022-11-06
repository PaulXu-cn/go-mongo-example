package mgolearn

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

func TestFindOne() {
    var err error
    session, err := mgo.DialWithTimeout(mongoDsn, 10*time.Second)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    // select db、database
    dbConn := session.DB("test").C("people")

    var one = People{}
    var theId = bson.ObjectIdHex("636731bb73abe5b5a790b9d9")
    err = dbConn.FindId(theId).One(&one)
    if nil != err {
        if mgo.ErrNotFound != err {
            fmt.Printf("mgo find people by Id err %v\n", err)
        }
    } else {
        fmt.Printf("mgo find people by Id re %+v \n", one)
    }

    var filter = bson.M{"name": "Lee"}
    err = dbConn.Find(filter).One(&one)
    if nil != err {
        if mgo.ErrNotFound != err {
            fmt.Printf("mgo find one people err %v\n", err)
        }
    } else {
        fmt.Printf("mgo find one people re %+v \n", one)
    }
}

func TestFind() {
    var err error
    session, err := mgo.DialWithTimeout(mongoDsn, 10*time.Second)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    // select db、database
    dbConn := session.DB("test").C("people")

    var men = []People{}
    var filter = bson.M{"name": bson.M{"$regex": "^l", "$options": "im"}}
    err = dbConn.Find(filter).All(&men)
    if nil != err {
        if mgo.ErrNotFound != err {
            fmt.Printf("mgo find people err %v\n", err)
        }
    } else {
        fmt.Printf("mgo find people re %+v \n", men)
    }
}
