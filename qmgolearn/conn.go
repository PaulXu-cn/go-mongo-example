package qmgolearn

import (
	"context"
	"fmt"
	"time"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongoDsn   = "mongodb://admin:123456@127.0.0.1:27017"
	db         = "test"
	collection = "test"
)

// qmgo newClient 方式连接库
func TestConn() *qmgo.Client {
	var timeout int64 = 10 * 1000
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 设置连接库超时时间
	defer cancel()
	var config = qmgo.Config{
		Uri: mongoDsn,
		//Database:         "admin", // 认证 库, 非后面实际使用的业务库
		ConnectTimeoutMS: &timeout,
		ReadPreference:   &qmgo.ReadPref{Mode: readpref.SecondaryPreferredMode}, // 设置主从库读取策略
	}
	client, err := qmgo.NewClient(ctx, &config)
	if nil != err {
		fmt.Printf("qmgo connect mongo err %v\n", err)
	} else {
		defer func() { //实际业务中，按照需要添加 close func
			if err = client.Close(ctx); err != nil { // 关闭连接
				panic(err)
			}
		}()
	}

	db := client.Database(db)
	coll := db.Collection(collection)
	collName := coll.GetCollectionName()
	fmt.Printf("qmgo choose collection %v\n", collName)
	return client
}

// qmgo open 方式连接库
func TestOpen() *qmgo.QmgoClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:      mongoDsn,             // mongo dsn
		Database: db, Coll: collection, // 连接时就选择 db、collection
		ReadPreference: &qmgo.ReadPref{Mode: readpref.PrimaryPreferredMode}})
	if nil != err {
		fmt.Printf("qmgo connect mongo err %v\n", err)
	} else {
		defer func() { //实际业务中，按照需要添加 close func
			if err = cli.Close(ctx); err != nil { // 关闭连接
				panic(err)
			}
		}()
	}

	err = cli.Ping(5)
	if nil != err {
		fmt.Printf("qmgo ping err %v\n", err)
	}

	ver := cli.ServerVersion()
	fmt.Printf("qmgo server version %s\n", ver)

	// 在 qmgo 中获取 collection 列表方式
	// 由于无法获取到 mongo.Client 对象，所有无法获取到 database 列表
	mongoColl, err := cli.CloneCollection()
	if nil != err {
		fmt.Printf("qmgo clone and get mongo.collection err %v\n", err)
	}
	if colls, err := mongoColl.Database().ListCollectionNames(ctx, bson.M{}); nil != err {
		fmt.Printf("qmgo list collection err %v\n", err)
	} else {
		fmt.Printf("qmgo list collection %v\n", colls)
	}
	return cli
}

// qmgo newClient 方式连接库
func GetConn() *qmgo.Client {
	var timeout int64 = 10 * 1000
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 设置连接库超时时间
	var config = qmgo.Config{
		Uri: mongoDsn,
		//Database:         "admin", // 认证 库, 非后面实际使用的业务库
		ConnectTimeoutMS: &timeout,
		ReadPreference:   &qmgo.ReadPref{Mode: readpref.SecondaryPreferredMode}, // 设置主从库读取策略
	}
	client, err := qmgo.NewClient(ctx, &config)
	if nil != err {
		fmt.Printf("qmgo connect mongo err %v\n", err)
	}

	//defer client.Close(ctx)
	return client
}
