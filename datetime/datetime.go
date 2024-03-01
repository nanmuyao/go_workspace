package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// doc: https://zhuanlan.zhihu.com/p/144308830
func mongoDemos() {
	var (
		client     *mongo.Client
		err        error
		db         *mongo.Database
		collection *mongo.Collection
	)
	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 my_db
	db = client.Database("my_db")

	//3.选择表 my_collection
	collection = db.Collection("my_collection")
	collection = collection
}

func timeDemos() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().AddDate(0, 0, -1))
}

func GetTimeNdaysAgo(days int) time.Time {
	// 获取当前时间
	now := time.Now()

	// 计算七天前的时间
	nDaysAgo := now.AddDate(0, 0, days)

	// 设置时间为0点
	sevenDaysAgoMidnight := time.Date(
		nDaysAgo.Year(), nDaysAgo.Month(), nDaysAgo.Day(),
		0, 0, 0, 0, nDaysAgo.Location())

	return sevenDaysAgoMidnight
}

func main() {
	fmt.Println("hello")
	//fmt.Println(time.Now())
	//timeDemos()
	fmt.Println(GetTimeNdaysAgo(1))
}
