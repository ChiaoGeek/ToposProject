package orm

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MonogoClient struct {
	Addr string
	Database string
	Port string
	ExceedTime int
	MClient *mongo.Client
}

type QueryResult struct {
	ID         primitive.ObjectID	`json:"id" bson:"_id"`
	LSTMODDATE	string	`json:"LSTMODDATE" bson:"LSTMODDATE"`
	CNSTRCT_YR	float64	`json:"CNSTRCT_YR" bson:"CNSTRCT_YR"`
	LSTSTATYPE	string	`json:"LSTSTATYPE" bson:"LSTSTATYPE"`
	FEAT_CODE	string	`json:"FEAT_CODE" bson:"FEAT_CODE"`
	GROUNDELEV	float64	`json:"GROUNDELEV" bson:"GROUNDELEV"`
	SHAPE_AREA	float64	`json:"SHAPE_AREA" bson:"SHAPE_AREA"`
	MPLUTO_BBL	string	`json:"MPLUTO_BBL" bson:"name"`
	GEOMSOURCE	string	`json:"GEOMSOURCE" bson:"GEOMSOURCE"`
	GEOLATLON	string	`json:"GEOLATLON" bson:"GEOLATLON"`
	HEIGHTROOF	float64	`json:"HEIGHTROOF" bson:"HEIGHTROOF"`
	BIN	string	`json:"BIN" bson:"BIN"`
	DOITT_ID	string	`json:"DOITT_ID" bson:"DOITT_ID"`
	SHAPE_LEN	float64	`json:"SHAPE_LEN" bson:"SHAPE_LEN"`
	BASE_BBL	string	`json:"BASE_BBL" bson:"BASE_BBL"`
	GEOTYPE	string	`json:"GEOTYPE" bson:"GEOTYPE"`
}

func (m *MonogoClient) SetUp() {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}
	m.MClient = client
}


func (m *MonogoClient) Query(collection string, input bson.M) []QueryResult{
	res := make([]QueryResult, 0)
	if len(input) == 0 {
		return res
	}
	c := m.MClient.Database(m.Database).Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	findOptions := options.Find()
	findOptions.SetLimit(100)
	//findOptions.SetSort(bson.D{{"SHAPE_LEN", -1}})
	curr, err := c.Find(ctx, input, findOptions)
	defer curr.Close(ctx)
	if err != nil {
		log.Fatal("query error: ", err)
	}

	for curr.Next(ctx) {
		var result QueryResult

		err := curr.Decode(&result)
		if err != nil { log.Fatal(err) }
		// do something with result....
		res = append(res, result)
	}
	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}
	//r,_ := json.Marshal(res)
	//return r
	return res
}

func (m *MonogoClient) GetAveHeightByBaseBbl(collection string, baseBbl string) float64 {
	c := m.MClient.Database(m.Database).Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	curr, err := c.Find(ctx, bson.M{"BASE_BBL" : baseBbl})
	defer curr.Close(ctx)
	if err != nil {
		fmt.Println("1", err)
		return -1
	}
	var res float64
	var count int
	for curr.Next(ctx) {
		var result QueryResult
		err := curr.Decode(&result)
		if err != nil {
			fmt.Println("2", err)
		}
		res += result.HEIGHTROOF
		count++
	}
	if err := curr.Err(); err != nil {
		fmt.Println("3", err)
		return -1
	}
	if count == 0 {
		fmt.Println("4")
		return -1
	}
	return res / float64(count);
}
