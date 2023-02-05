package qmgolearn

import (
	"context"
	"fmt"
	"time"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// qmgo Find().One()
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

// qmgo Find().All()
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

	var filter = bson.M{} // findAll()
	var students = []Student{}
	//if err := cli.Find(ctx, nil).All(&students); nil != err {	// 这里 find(nil) 代码不会报错，但是查不出任何东西，
	if err := cli.Find(ctx, filter).All(&students); nil != err {
		fmt.Printf("qmgo find student err %v \n", err)
	} else {
		fmt.Printf("qmgo find student result %+v, err %v \n", students, err)
	}
}

// qmgo Find().All()
func TestFindManyCursor() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := GetConn()
	cli := client.Database("test").Collection("test")

	var filter = bson.M{}
	var students = []Student{}

	if cur := cli.Find(ctx, filter).Cursor(); nil != cur {
		var more = true
		for more {
			var student = Student{}
			more = cur.Next(&student)
			if !more {
				if err := cur.Err(); err != nil {
					fmt.Printf("qmgo find.curor err %v\n", err)
				}
				break
			}
			students = append(students, student)
		}
	}

	fmt.Printf("qmgo find students by curosr result %+v\n", students)
}

// qmgo Find().All()
func TestFindCount() {
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
	if count, err := cli.Find(ctx, filter).Count(); nil != err {
		fmt.Printf("qmgo find student err %v \n", err)
	} else {
		fmt.Printf("qmgo count student result %d, err %v \n", count, err)
	}
}

// qmgo Find().Limit().Sort().Skip()
func TestFindLimitSortSkip() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := GetConn()
	cli := client.Database("test").Collection("test")

	var students = []Student{}
	var filter = bson.M{"age": 9}

	var sort = []string{"_id", "timestamp"}
	var limit int64 = 200
	var skip int64 = 1
	query := cli.Find(ctx, filter).Sort(sort...).Limit(limit).Skip(skip) // sort() limit() skip()

	if err := query.All(&students); nil != err {
		fmt.Printf("qmgo find student err %v \n", err)
	} else {
		fmt.Printf("qmgo find student result %+v, err %v \n", students, err)
	}
}
