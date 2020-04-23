package gmongo

import (
	"context"
	"gin_test/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Mongo struct {
	MongoConn *mongo.Client
}

var MongoDB *Mongo

func init() {
	MongoDB = &Mongo{
		MongoConn: SetConnect(),
	}
}

func SetConnect() *mongo.Client {
	//get mongo config && set client options
	config := conf.LoadConfig()
	clientOptions := options.Client().ApplyURI(config.MongoDB.Address).SetMaxPoolSize(config.MongoDB.MaxPoolSize)

	//connect to mongo
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("mongo connect err: %v", err)
	}
	//ping the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("mongo ping err: %v", err)
	}
	log.Println("mongo connect success")
	return client
}

func FindOne(db, collect, queryKey string, queryValue interface{}) *mongo.SingleResult {
	client := MongoDB.MongoConn
	collection, _ := client.Database(db).Collection(collect).Clone() //利用已经存在的mongo连接
	filter := bson.M{queryKey: queryValue}
	result := collection.FindOne(context.TODO(), filter)
	return result
}

func InsertOne(db, collect string, insertValue interface{}) (*mongo.InsertOneResult, error) {
	client := MongoDB.MongoConn
	collection := client.Database(db).Collection(collect)
	insertResult, err := collection.InsertOne(context.TODO(), insertValue)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func DeleteOne(db, collect, deleteKey string, deleteValue interface{}) (int64, error) {
	client := MongoDB.MongoConn
	collection := client.Database(db).Collection(collect)
	filter := bson.M{deleteKey: deleteValue}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

//deleteMany
//update
//count
//exist
