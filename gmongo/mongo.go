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

// todo 优化init
// 此处不应该用init来初始化连接，应该显示的定义一个NewClient函数，使用的时候显示调用NewClient(Client)
// 避免另外的package一引入这个包就会发生资源的初始化，资源初始化非常耗时且耗资源
/*func NewClient(MongoConn *ClientConn) Mongo {
	return &Mongo{
		MongoConn: MongoConn,
	}
}*/
func init() {
	MongoDB = &Mongo{
		MongoConn: SetConnect(),
	}
}

func SetConnect() *mongo.Client {
	//get mongo config && set client options
	config := conf.LoadConfig()
	clientOptions := options.Client().ApplyURI(config.MongoDB.Address).SetMaxPoolSize(config.MongoDB.MaxPoolSize)

	//利用context的WithTimeout做连接超时限制，达到超时，所有context取消
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//connect to mongo
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

// todo 通过接口暴露实现的函数
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

//deleteMany  同DeleteOne使用，使用DeleteMany()

//update
func UpdateOne(db, collect, queryKey, queryValue, updateKey string, updateValue interface{}) (int64, error) {
	client := MongoDB.MongoConn
	collection := client.Database(db).Collection(collect)
	query := bson.M{queryKey: queryValue}
	update := bson.M{"$set": bson.M{updateKey: updateValue}}
	result, err := collection.UpdateOne(context.TODO(), query, update)
	if result.MatchedCount == 0 {
		return -1, nil
	}
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

//一个集合中的数据个数
func CollectionCount(db, collect string) (string, int64) {
	client := MongoDB.MongoConn
	collection := client.Database(db).Collection(collect)
	collectionName := collection.Name()
	size, _ := collection.EstimatedDocumentCount(context.TODO())
	return collectionName, size
}

//查找一个集合中全部的数据  db.user.find(),并且按照age正序输出
func FindAll(db, collect string) (string, interface{}, error) {
	client := MongoDB.MongoConn
	collection := client.Database(db).Collection(collect)
	collectionName := collection.Name()
	opts := options.Find().SetSort(bson.D{{"age", 1}})
	results, err := collection.Find(context.TODO(), bson.D{}, opts)
	var data []bson.M
	if err = results.All(context.TODO(), &data); err != nil {
		return collectionName, nil, err
	}
	return collectionName, data, nil
}
