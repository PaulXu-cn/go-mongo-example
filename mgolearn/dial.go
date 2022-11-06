package mgolearn


import (
	"fmt"
    "gopkg.in/mgo.v2"
    "time"
)


var(
    mongoDsn =  "mongodb://admin:123456@127.0.0.1:27017"
    database = "test"
    collection = "user"
)

func TestDial() {
    var err error
    var dialInfo = mgo.DialInfo{
        Addrs: []string{"127.0.0.1:27017"},
        Timeout: 10 * time.Second,
        Database: "admin",  // 这里填认证库，并不是等下你要连接业务库！
        Username: "root",
        Password: "123456",
    }
    session, err := mgo.DialWithInfo(&dialInfo)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    err = session.Ping()
    if nil != err {
        fmt.Printf("mgo ping err %v\n", err)
    } else {
        fmt.Printf("mgo ping success~\n")
    }
}

func TestDial2() {
    var err error
    session, err := mgo.Dial(mongoDsn)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }
    session, err = mgo.DialWithTimeout(mongoDsn, 10 * time.Second)
    if nil != err {
        fmt.Printf("mgo dial err %v\n", err)
    }

    // auto close session before func finish
    defer session.Close()

    // show database
    dbs, err := session.DatabaseNames()
    if nil != err {
        fmt.Printf("mgo show database err %v\n", err)
    } else {
        fmt.Printf("mgo show database %v\n", dbs)
    }

    // select db
    dbConn := session.DB(database)

    // show collections
    colls, err := dbConn.CollectionNames()
    if nil != err {
        fmt.Printf("mgo show collection err %v\n", err)
    } else {
        fmt.Printf("mgo show collection %v\n", colls)
    }

    // select collection
    coll := dbConn.C(collection)
    fmt.Printf("mgo select db %v collection %v \n", dbConn, coll)

    //stats := mgo.GetStats()
    //fmt.Printf("mgo stats %+v", stats)
}
