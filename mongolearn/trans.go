package mongolearn

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func TestTrans() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    sess, err := client.StartSession()
    if nil != err {
        fmt.Printf("mongo start session  err %v \n", err)
    }

    // end session
    defer sess.EndSession(ctx)

    if err = mongo.WithSession(ctx, sess, func(sc mongo.SessionContext) error {
        // 开始
        if err = sess.StartTransaction(); nil != err {
            fmt.Printf("mongo strat transaction err %v \n", err)
            return err
        }

        // do curd
        var user = User{}
        var filter = bson.D{{"name", "Bob"}}
        err = coll.FindOne(sc, filter).Decode(&user)
        if nil != err {
            if err == mongo.ErrNoDocuments {
                // Do something when no record was found
            } else {
                fmt.Printf("mongo find one err in transaction %v \n", err)
                return err
            }
        } else {
            fmt.Printf("mongo find one user in transaction %+v \n", user)
        }
        // end curd

        // commit
        if err = sess.CommitTransaction(sc); err != nil {
            fmt.Printf("mongo commit transaction err %v \n", user)
            return err
        }
        return nil
    }); err != nil {
        if abortErr := sess.AbortTransaction(context.Background()); abortErr != nil {
            fmt.Printf("mongo abort transaction err %v \n", err)
        }
        fmt.Printf("mongo transaction err %v \n", err)
    }
}


func TestTrans2() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    db := client.Database("test")
    coll := db.Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    if err = client.UseSession(ctx, func(sc mongo.SessionContext) error {
        // 开始
        if err = sc.StartTransaction(); nil != err {
            fmt.Printf("mongo strat transaction err %v \n", err)
            return err
        }

        // do curd
        var user = User{}
        var filter = bson.D{{"name", "Bob"}}
        err = coll.FindOne(sc, filter).Decode(&user)
        if nil != err {
            if err == mongo.ErrNoDocuments {
                // Do something when no record was found
            } else {
                fmt.Printf("mongo find one err in transaction %v \n", err)
                return err
            }
        } else {
            fmt.Printf("mongo find one user in transaction %+v \n", user)
        }
        // end curd

        // commit
        if err = sc.CommitTransaction(sc); err != nil {
            fmt.Printf("mongo commit transaction err %v \n", user)
            return err
        }
        return nil
    }); err != nil {
        fmt.Printf("mongo transaction err %v \n", err)
    }
}
