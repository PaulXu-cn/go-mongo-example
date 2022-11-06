package mongolearn

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
    Name      string    `bson:"name"`
    Age       uint32    `bson:"age"`
    Weight    float32   `bson:"weight"`
    Studying  bool      `bson:"studying"`
    Tag       []string  `bson:"tag"`
    CreatedAt time.Time `bson:"created_at"`
}

func TestFindOne() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    var user = User{}
    var filter = bson.D{{"name", "Bob"}}
    err = coll.FindOne(ctx, filter).Decode(&user)
    if nil != err {
        if err == mongo.ErrNoDocuments {
            // Do something when no record was found
        } else {
            fmt.Printf("mongo find one err %v \n", err)
        }
    } else {
        fmt.Printf("mongo find one user %+v \n", user)
    }
}


func TestFindList() {
    var err error
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

    coll := client.Database("test").Collection("user") // 选择 db、collection

    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    var result = []User{}
    var filter = bson.M{"name": bson.M{"$in": []string{ "Rose", "Jack"}}}

    cur, err := coll.Find(ctx, filter)
    if err != nil {
        fmt.Printf("mongo find err %v \n", err)
    }
    defer cur.Close(ctx)
    for cur.Next(ctx) {
        var row = User{}
        err := cur.Decode(&row)
        if err != nil {
            fmt.Printf("mongo decode err %v \n", err)
        }
        result = append(result, row)
    }
    if err := cur.Err(); err != nil {
        fmt.Printf("mongo find list cur err %v \n", err)
    }

    for k, i := range result {
        fmt.Printf("mongo loop find list k: %d, v: %+v \n", k, i)
    }
}

